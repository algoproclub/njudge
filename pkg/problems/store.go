package problems

import (
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

type problemNotFoundError struct {
	name string
}

func (perr problemNotFoundError) Error() string {
	return "problem not found: " + perr.name
}

func (perr problemNotFoundError) Is(target error) bool {
	return target == ErrorProblemNotFound
}

var ErrorProblemParse = errors.New("can't parse problems")

type ProblemParseError struct {
	Errors   []error
	Problems []string
}

func (perr ProblemParseError) Error() string {
	return fmt.Sprintf("can't parse problems: %v", perr.Problems)
}

func (perr ProblemParseError) Is(target error) bool {
	return target == ErrorProblemParse
}

//Store is an interface which is used to access a bunch of problems for example from the filesystem
type Store interface {
	List() ([]string, error)
	Has(string) (bool, error)
	Get(string) (Problem, error)
	MustGet(string) Problem
	Update() error
	UpdateProblem(string, string) error
}

type FsStore struct {
	cs  ConfigStore
	fs  afero.Fs
	dir string

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

func NewFsStore(dir string, options ...FsStoreOptions) *FsStore {
	fsStore := &FsStore{cs: globalConfigStore, fs: afero.NewOsFs(), dir: dir, problems: make(map[string]Problem), problemsList: make([]string, 0)}
	for _, opt := range options {
		opt(fsStore)
	}

	return fsStore
}

func (s *FsStore) List() ([]string, error) {
	s.RLock()
	defer s.RUnlock()

	lst := make([]string, len(s.problemsList))
	copy(lst, s.problemsList)

	return lst, nil
}

func (s *FsStore) Has(p string) (bool, error) {
	s.RLock()
	defer s.RUnlock()

	for _, key := range s.problemsList {
		if key == p {
			return true, nil
		}
	}

	return false, nil
}

func (s *FsStore) Get(p string) (Problem, error) {
	s.RLock()
	defer s.RUnlock()

	if res, ok := s.problems[p]; ok {
		return res, nil
	}
	return nil, problemNotFoundError{p}
}

func (s *FsStore) MustGet(p string) Problem {
	res, err := s.Get(p)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *FsStore) Update() error {
	errs := ProblemParseError{Errors: make([]error, 0), Problems: make([]string, 0)}
	lst := make([]string, 0)

	if err := afero.Walk(s.fs, s.dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return err
		}

		if strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if s.dir == path {
			return nil
		}

		//skip directories recursively with the .njudge_ignore file
		if _, err := os.Stat(filepath.Join(path, ".njudge_ignore")); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		} else {
			return nil
		}

		if err := s.UpdateProblem(info.Name(), path); err != nil {
			if !errors.Is(err, ErrorNoMatch) {
				errs.Errors = append(errs.Errors, fmt.Errorf("%s: %w", info.Name(), err))
				errs.Problems = append(errs.Problems, info.Name())

				if err == ErrorNoMatch {
					return nil
				} else {
					return filepath.SkipDir
				}
			} else {
				return nil
			}
		} else {
			lst = append(lst, info.Name())
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

func (s *FsStore) UpdateProblem(p string, path string) error {
	s.Lock()
	defer s.Unlock()

	prob, err := s.cs.Parse(path)
	if err != nil {
		return err
	}

	s.problems[p] = prob
	return nil
}
