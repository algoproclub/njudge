package cpp

import (
	"context"
	"io"
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
}

type Option func(*Cpp)

func WithCompileArgs(args []string) Option {
	return func(cpp *Cpp) {
		cpp.compileArgs = make([]string, len(args))
		copy(cpp.compileArgs, args)
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
