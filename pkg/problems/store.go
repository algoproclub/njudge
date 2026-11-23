package problems

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/afero"
)

var ErrorProblemNotFound = errors.New("problem not found")

type NotFoundError struct {
	Name string
}

func (err NotFoundError) Error() string {
	return "problem not found: " + err.Name
}

func (err NotFoundError) Is(target error) bool {
	return target == ErrorProblemNotFound
}

var ErrorProblemParse = errors.New("can't parse problems")

type ProblemParseError struct {
	Errors   []error
	Problems []string
}

func (err ProblemParseError) Error() string {
	return fmt.Sprintf("can't parse problems: %v", err.Errors)
}

func (ProblemParseError) Is(target error) bool {
	return target == ErrorProblemParse
}

// Store is an interface which is used to access a bunch of problems for example from the filesystem
type Store interface {
	ListProblems() ([]string, error)
	HasProblem(string) (bool, error)
	GetProblem(string) (Problem, error)
	MustGetProblem(string) Problem
	UpdateProblems() error
	UpdateProblem(path string, id string) error
}

var (
	FsStoreIgnoreFile = ".njudge_ignore"
	FsStorePrefixFile = ".njudge_prefix"
)

type FsStore struct {
	cs           ConfigStore
	fs           afero.Fs
	root         string
	ignorePrefix bool

	ByPath     map[string]string
	signatures map[string]string

	problems     map[string]Problem
	problemsList []string

	sync.RWMutex
}

type FsStoreOptions func(*FsStore)

func FsStoreUseConfigStore(cs ConfigStore) FsStoreOptions {
	return func(fs *FsStore) {
		fs.cs = cs
	}
}

func FsStoreUseFs(afs afero.Fs) FsStoreOptions {
	return func(fs *FsStore) {
		fs.fs = afs
	}
}

func FsStoreIgnorePrefix() FsStoreOptions {
	return func(fs *FsStore) {
		fs.ignorePrefix = true
	}
}

func NewFsStore(root string, options ...FsStoreOptions) *FsStore {
	fsStore := &FsStore{
		cs:           globalConfigStore,
		fs:           afero.NewOsFs(),
		root:         root,
		problems:     make(map[string]Problem),
		problemsList: make([]string, 0),
		ByPath:       make(map[string]string),
		signatures:   make(map[string]string),
	}
	for _, opt := range options {
		opt(fsStore)
	}

	return fsStore
}

func (s *FsStore) ListProblems() ([]string, error) {
	s.RLock()
	defer s.RUnlock()

	lst := make([]string, len(s.problemsList))
	copy(lst, s.problemsList)

	return lst, nil
}

func (s *FsStore) HasProblem(p string) (bool, error) {
	if _, err := s.GetProblem(p); err != nil && errors.Is(err, ErrorProblemNotFound) {
		return false, nil
	}

	return true, nil
}

func (s *FsStore) GetProblem(p string) (Problem, error) {
	s.RLock()
	defer s.RUnlock()

	if res, ok := s.problems[p]; ok {
		return res, nil
	}

	for _, prob := range s.problems {
		if prob.Name() == p {
			return prob, nil
		}
	}

	return nil, NotFoundError{p}
}

func (s *FsStore) MustGetProblem(p string) Problem {
	res, err := s.GetProblem(p)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *FsStore) UpdateProblems() error {
	errs := ProblemParseError{Errors: make([]error, 0), Problems: make([]string, 0)}
	lst := make([]string, 0)

	prefix, atPath := "", ""
	getId := func(path string) string {
		rel, err := filepath.Rel(atPath, path)
		if err != nil || prefix == "" {
			return filepath.Base(path)
		}

		if len(rel) == 1 || rel[0] != '.' {
			return prefix + "_" + filepath.Base(path)
		}

		return filepath.Base(path)
	}

	if err := afero.Walk(s.fs, s.root, func(path string, info fs.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return err
		}

		if strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if s.root == path {
			return nil
		}

		//skip directories recursively with the .njudge_ignore file
		if _, err := s.fs.Stat(filepath.Join(path, FsStoreIgnoreFile)); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		} else {
			return filepath.SkipDir
		}

		if !s.ignorePrefix {
			prefixFile := filepath.Join(path, FsStorePrefixFile)
			if _, err := s.fs.Stat(prefixFile); err != nil {
				if !errors.Is(err, os.ErrNotExist) {
					return err
				}
			} else {
				prefixBytes, err := afero.ReadFile(s.fs, prefixFile)
				if err != nil {
					return err
				}

				prefix = strings.TrimSpace(string(prefixBytes))
				atPath = path
			}
		}

		// unchanged problems aren't parsed again
		s.RLock()
		name, known := s.ByPath[path]
		oldSignature := s.signatures[path]
		s.RUnlock()
		if known && name == getId(path) && oldSignature != "" {
			if sig, err := s.signature(path); err == nil && sig == oldSignature {
				lst = append(lst, getId(path))
				return filepath.SkipDir
			}
		}

		if err := s.UpdateProblem(path, getId(path)); err != nil {
			if !errors.Is(err, ErrorNoMatch) {
				errs.Errors = append(errs.Errors, fmt.Errorf("%s: %w", info.Name(), err))
				errs.Problems = append(errs.Problems, info.Name())

				if errors.Is(err, ErrorNoMatch) {
					return nil
				} else {
					return filepath.SkipDir
				}
			} else {
				return nil
			}
		} else {
			lst = append(lst, getId(path))
			return filepath.SkipDir
		}
	}); err != nil {
		return err
	}

	s.Lock()
	s.problemsList = make([]string, len(lst))
	copy(s.problemsList, lst)
	s.Unlock()

	if len(errs.Errors) == 0 {
		return nil
	}
	return errs
}

type ProblemWrapper struct {
	Problem

	NameOverride string
}

func (pw ProblemWrapper) Name() string {
	return pw.NameOverride
}

// signature returns a hash of the names, sizes, modification times and modes of the files in a problem's directory.
func (s *FsStore) signature(path string) (string, error) {
	h := sha1.New()
	if err := afero.Walk(s.fs, path, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if p != path && info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		rel, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}

		fmt.Fprintf(h, "%s|%d|%d|%d\n", rel, info.Size(), info.ModTime().UnixNano(), info.Mode())
		return nil
	}); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func (s *FsStore) UpdateProblem(path string, name string) error {
	// The signature of a known problem is calculated before parsing, so changes made during the parse aren't missed.
	// New paths might not be problems at all (for example categories), so theirs is calculated only after a successful parse.
	s.RLock()
	_, known := s.ByPath[path]
	s.RUnlock()

	signature := ""
	if known {
		signature, _ = s.signature(path)
	}

	prob, err := s.cs.Parse(s.fs, path)
	if err != nil {
		s.Lock()
		delete(s.signatures, path)
		s.Unlock()
		return err
	}
	if !known {
		signature, _ = s.signature(path)
	}

	s.Lock()
	defer s.Unlock()

	s.ByPath[path] = name
	s.problems[name] = ProblemWrapper{prob, name}
	if signature != "" {
		s.signatures[path] = signature
	} else {
		delete(s.signatures, path)
	}
	return nil
}
