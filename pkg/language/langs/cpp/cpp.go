package cpp

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/flock"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"

	"github.com/mraron/njudge/pkg/language"
)

type Cpp struct {
	id   string
	name string

	compileArgs []string

	pch             *pchCache
	maxArtifactSize int64
}

type Option func(*Cpp)

func WithCompileArgs(args []string) Option {
	return func(cpp *Cpp) {
		cpp.compileArgs = make([]string, len(args))
		copy(cpp.compileArgs, args)
	}
}

func WithPCHCache(dir string, cacheSize int64) Option {
	return func(cpp *Cpp) {
		if dir == "" {
			cpp.pch = nil
			return
		}
		cpp.pch = &pchCache{
			dir:         dir,
			maxSize:     cacheSize,
			building:    make(map[string]bool),
			failedUntil: make(map[string]time.Time),
		}
	}
}

// WithMaxArtifactSize limits the size of the compiled a.out artifact extracted
// from the sandbox. A non-positive value preserves unlimited extraction.
func WithMaxArtifactSize(maxBytes int64) Option {
	return func(cpp *Cpp) {
		cpp.maxArtifactSize = maxBytes
	}
}

func New(ID, name string, opts ...Option) *Cpp {
	res := &Cpp{
		id:   ID,
		name: name,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func (c Cpp) ID() string {
	return c.id
}

func (c Cpp) DisplayName() string {
	return c.name
}

func (c Cpp) DefaultFilename() string {
	return "main.cpp"
}

func (c Cpp) runCompiler(ctx context.Context, s sandbox.Sandbox, output io.Writer, extraArgs []string, dirMaps []sandbox.DirectoryMap) (*sandbox.Status, error) {
	rc := sandbox.RunConfig{
		MaxProcesses:     200,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      768 * memory.MiB,
		Stdout:           output,
		Stderr:           output,
		WorkingDirectory: s.Pwd(),
		DirectoryMaps:    dirMaps,
	}

	return s.Run(
		ctx,
		rc,
		"/usr/bin/g++",
		append(c.compileArgs, extraArgs...)...,
	)
}

// pchCache is the shared state of a precompiled header cache, which is laid out as <dir>/<key>/bits/stdc++.h.gch.
type pchCache struct {
	dir     string
	maxSize int64

	sync.Mutex
	building    map[string]bool
	failedUntil map[string]time.Time
}

const pchFileName = "stdc++.h.gch"

// pchRetryAfter is the time to wait before building a PCH again after a failed build.
const pchRetryAfter = 10 * time.Minute

func (p *pchCache) startBuild(key string) bool {
	p.Lock()
	defer p.Unlock()

	if p.building[key] || time.Now().Before(p.failedUntil[key]) {
		return false
	}
	p.building[key] = true
	return true
}

func (p *pchCache) finishBuild(key string, err error) {
	p.Lock()
	defer p.Unlock()

	delete(p.building, key)
	if err != nil {
		p.failedUntil[key] = time.Now().Add(pchRetryAfter)
	}
}

func (p *pchCache) file(key string) string {
	return filepath.Join(p.dir, key, "bits", pchFileName)
}

// cleanUp removes the least recently used PCHs (except for keep) until the cache fits into maxSize.
func (p *pchCache) cleanUp(keep string) {
	type Entry struct {
		Name string
		Time time.Time
		Size int64
	}
	entries := []Entry{}
	var totalSize int64 = 0

	files, err := os.ReadDir(p.dir)
	if err != nil {
		return
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		if !file.IsDir() {
			if strings.HasPrefix(file.Name(), ".build-") && time.Since(info.ModTime()) > time.Hour {
				_ = os.Remove(filepath.Join(p.dir, file.Name()))
			}
			continue
		}

		gch, err := os.Stat(p.file(file.Name()))
		if err != nil {
			continue
		}
		entries = append(entries, Entry{
			Name: file.Name(),
			Time: gch.ModTime(),
			Size: gch.Size(),
		})
		totalSize += gch.Size()
	}

	slices.SortFunc(entries, func(i, j Entry) int {
		return i.Time.Compare(j.Time)
	})

	for _, entry := range entries {
		if totalSize <= p.maxSize {
			break
		}

		if entry.Name == keep {
			continue
		}

		if err := os.RemoveAll(filepath.Join(p.dir, entry.Name)); err != nil {
			continue
		}
		totalSize -= entry.Size
	}
}

func (c Cpp) buildPCH(ctx context.Context, s sandbox.Sandbox, key string) error {
	if err := sandbox.CreateFile(s, sandbox.File{
		Name:   "pch.h",
		Source: io.NopCloser(strings.NewReader("#include <bits/stdc++.h>")),
	}); err != nil {
		return err
	}

	status, err := c.runCompiler(ctx, s, io.Discard, []string{"pch.h", "-o", pchFileName}, nil)
	if err != nil {
		return err
	}
	if status.Verdict != sandbox.VerdictOK {
		return fmt.Errorf("pch build failed: %v", status.Verdict)
	}

	file, err := s.Open(pchFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := os.MkdirAll(c.pch.dir, 0755); err != nil {
		return err
	}

	tmp, err := os.CreateTemp(c.pch.dir, ".build-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	_, err = io.Copy(tmp, file)
	if err == nil {
		err = tmp.Chmod(0644)
	}
	if err = errors.Join(err, tmp.Close()); err != nil {
		return err
	}

	lock := flock.New(filepath.Join(c.pch.dir, "cache.lock"))
	if err := lock.Lock(); err != nil {
		return err
	}
	defer lock.Unlock()

	if err := os.MkdirAll(filepath.Dir(c.pch.file(key)), 0755); err != nil {
		return err
	}
	if err := os.Rename(tmp.Name(), c.pch.file(key)); err != nil {
		return err
	}

	c.pch.cleanUp(key)
	return nil
}

// pchKey identifies the PCH of the current compiler and flags.
func (c Cpp) pchKey() (string, error) {
	// Pre-compiled headers can only be used with the same compiler and flags they were built with.
	cacheKey := sha1.New()

	compilerInfo, err := os.Stat("/usr/bin/g++")
	if err != nil {
		return "", err
	}
	io.WriteString(cacheKey, compilerInfo.ModTime().String())
	io.WriteString(cacheKey, strings.Join(c.compileArgs, " "))
	return hex.EncodeToString(cacheKey.Sum(nil)), nil
}

// usePCH makes the precompiled header available for the compilation inside s, building it if necessary.
// It returns the directory maps and the compiler arguments needed to use it, both are empty if no PCH can be used.
func (c Cpp) usePCH(ctx context.Context, s sandbox.Sandbox) ([]sandbox.DirectoryMap, []string) {
	if c.pch == nil {
		return nil, nil
	}

	key, err := c.pchKey()
	if err != nil {
		return nil, nil
	}

	if _, err := os.Stat(c.pch.file(key)); err == nil {
		now := time.Now() // least recently used PCHs are evicted first
		_ = os.Chtimes(c.pch.file(key), now, now)
	} else {
		if !c.pch.startBuild(key) {
			return nil, nil
		}
		err := c.buildPCH(ctx, s, key)
		if ctx.Err() != nil {
			c.pch.finishBuild(key, nil) // a cancelled build isn't a failed one
		} else {
			c.pch.finishBuild(key, err)
		}
		if err != nil {
			return nil, nil
		}
	}

	if mapper, ok := s.(sandbox.DirectoryMapper); ok && mapper.SupportsDirectoryMaps() {
		return []sandbox.DirectoryMap{{Inside: "/pch", Outside: filepath.Join(c.pch.dir, key)}}, []string{"-isystem", "/pch"}
	}

	if err := s.MkdirAll("pch/bits", 0755); err != nil {
		return nil, nil
	}
	if err := os.Link(c.pch.file(key), filepath.Join(s.Pwd(), "pch", "bits", pchFileName)); err != nil {
		file, err := os.Open(c.pch.file(key))
		if err != nil {
			return nil, nil
		}
		if err := sandbox.CreateFile(s, sandbox.File{
			Name:   filepath.Join("pch", "bits", pchFileName),
			Source: file,
		}); err != nil {
			return nil, nil
		}
	}
	return nil, []string{"-isystem", "pch"}
}

func (c Cpp) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	err := sandbox.CreateFile(s, f)
	if err != nil {
		return nil, err
	}

	extraParams := []string{f.Name}
	for _, extra := range extras {
		err := sandbox.CreateFile(s, extra)
		if err != nil {
			return nil, err
		}

		if !strings.HasSuffix(extra.Name, ".h") {
			extraParams = append(extraParams, extra.Name)
		}
	}

	dirMaps, pchParams := c.usePCH(ctx, s)
	extraParams = append(extraParams, pchParams...)

	if _, err := c.runCompiler(ctx, s, stderr, extraParams, dirMaps); err != nil {
		return nil, err
	}

	if c.maxArtifactSize > 0 {
		return sandbox.ExtractFileWithLimit(s, "a.out", c.maxArtifactSize)
	}

	return sandbox.ExtractFile(s, "a.out")
}

func (Cpp) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return sandbox.RunBinary(ctx, s, binary, stdin, stdout, tl, ml)
}

var DefaultCompileArgs = []string{"-O2", "-static", "-DONLINE_JUDGE"}

var DefaultPCHCacheDir = filepath.Join(os.TempDir(), "njudge", "pch")

const DefaultPCHCacheSize = 1 << 30

var Std11 = New("cpp11", "C++ 11", WithCompileArgs(append(DefaultCompileArgs, "-std=c++11")), WithPCHCache(DefaultPCHCacheDir, DefaultPCHCacheSize))
var Std14 = New("cpp14", "C++ 14", WithCompileArgs(append(DefaultCompileArgs, "-std=c++14")), WithPCHCache(DefaultPCHCacheDir, DefaultPCHCacheSize))
var Std17 = New("cpp17", "C++ 17", WithCompileArgs(append(DefaultCompileArgs, "-std=c++17")), WithPCHCache(DefaultPCHCacheDir, DefaultPCHCacheSize))

// ConfigureDefaultPCHCache changes the PCH cache of Std11, Std14 and Std17, an empty dir disables it.
func ConfigureDefaultPCHCache(dir string, cacheSize int64) {
	for _, c := range []*Cpp{Std11, Std14, Std17} {
		WithPCHCache(dir, cacheSize)(c)
	}
}

func init() {
	language.DefaultStore.Register("cpp11", Std11)
	language.DefaultStore.Register("cpp14", Std14)
	language.DefaultStore.Register("cpp17", Std17)
}
