package cpp

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"strings"
	"time"

	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"

	"github.com/mraron/njudge/pkg/language"
)

type Cpp struct {
	id   string
	name string

	compileArgs []string

	pchCacheDir  string
	pchCacheSize int
}

type Option func(*Cpp)

func WithCompileArgs(args []string) Option {
	return func(cpp *Cpp) {
		cpp.compileArgs = make([]string, len(args))
		copy(cpp.compileArgs, args)
	}
}

func WithPCHCache(dir string, cacheSize int) Option {
	return func(cpp *Cpp) {
		cpp.pchCacheDir = dir
		cpp.pchCacheSize = cacheSize
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

func (c Cpp) runCompiler(ctx context.Context, s sandbox.Sandbox, output io.Writer, extraArgs []string) (*sandbox.Status, error) {
	rc := sandbox.RunConfig{
		MaxProcesses:     200,
		InheritEnv:       true,
		TimeLimit:        10 * time.Second,
		MemoryLimit:      512 * memory.MiB,
		Stdout:           output,
		Stderr:           output,
		WorkingDirectory: s.Pwd(),
	}

	return s.Run(
		ctx,
		rc,
		"/usr/bin/g++",
		append(c.compileArgs, extraArgs...)...,
	)
}

func (c Cpp) buildPCH(ctx context.Context, s sandbox.Sandbox) *os.File {
	if c.pchCacheDir == "" {
		return nil
	}

	// Pre-compiled headers can only be used with the same compiler and flags they were built with.
	cacheKey := sha1.New()

	compilerInfo, err := os.Stat("/usr/bin/g++")
	if err != nil {
		return nil
	}
	io.WriteString(cacheKey, compilerInfo.ModTime().String())
	io.WriteString(cacheKey, strings.Join(c.compileArgs, " "))
	pchFileName := hex.EncodeToString(cacheKey.Sum(nil)) + ".gch"

	if cachedFile, err := os.Open(c.pchCacheDir + "/" + pchFileName); err == nil {
		return cachedFile
	}

	if err := sandbox.CreateFile(s, sandbox.File{
		Name:   "pch.h",
		Source: io.NopCloser(strings.NewReader("#include <bits/stdc++.h>")),
	}); err != nil {
		return nil
	}

	if _, err := c.runCompiler(ctx, s, io.Discard, []string{"pch.h", "-o", pchFileName}); err != nil {
		return nil
	}

	file, err := sandbox.ExtractFile(s, pchFileName)
	if err != nil {
		return nil
	}
	defer file.Source.Close()

	outFile, err := os.Create(c.pchCacheDir + "/" + pchFileName)
	if err != nil {
		return nil
	}

	if _, err := io.Copy(outFile, file.Source); err != nil {
		return nil
	}
	outFile.Sync()
	outFile.Seek(0, io.SeekStart)

	// TODO: Remove old files if pchCacheSize is reached

	return outFile
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

	if pch := c.buildPCH(ctx, s); pch != nil {
		s.MkdirAll("/pch/bits", 0755)
		err = sandbox.CreateFile(s, sandbox.File{
			Name:   "/pch/bits/stdc++.h.gch",
			Source: pch,
		})

		if err == nil {
			extraParams = append(extraParams, "-isystem", "pch")
		}
	}

	if _, err := c.runCompiler(ctx, s, stderr, extraParams); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "a.out")
}

func (Cpp) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	return sandbox.RunBinary(ctx, s, binary, stdin, stdout, tl, ml)
}

var DefaultCompileArgs = []string{"-O2", "-static", "-DONLINE_JUDGE"}

var Std11 = New("cpp11", "C++ 11", WithCompileArgs(append(DefaultCompileArgs, "-std=c++11")))
var Std14 = New("cpp14", "C++ 14", WithCompileArgs(append(DefaultCompileArgs, "-std=c++14")))
var Std17 = New("cpp17", "C++ 17", WithCompileArgs(append(DefaultCompileArgs, "-std=c++17")))

func init() {
	language.DefaultStore.Register("cpp11", Std11)
	language.DefaultStore.Register("cpp14", Std14)
	language.DefaultStore.Register("cpp17", Std17)
}
