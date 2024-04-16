package csharp

import (
	"context"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"io"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

type CSharp struct{}

func (CSharp) ID() string {
	return "csharp"
}

func (CSharp) DisplayName() string {
	return "C# (mono)"
}

func (CSharp) DefaultFilename() string {
	return "main.cs"
}

func (CSharp) Compile(ctx context.Context, s sandbox.Sandbox, f sandbox.File, stderr io.Writer, extras []sandbox.File) (*sandbox.File, error) {
	err := sandbox.CreateFile(s, f)
	if err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		InheritEnv:    true,
		DirectoryMaps: []sandbox.DirectoryMap{{"/etc", "/etc", nil}},
		MaxProcesses:  -1,
		TimeLimit:     10 * time.Second,
		MemoryLimit:   1 * memory.GiB,
		Stdout:        stderr,
		Stderr:        stderr,
	}
	if _, err := s.Run(ctx, rc, "/usr/bin/mcs", sandbox.SplitArgs("-out:main.exe -optimize+ "+f.Name)...); err != nil {
		return nil, err
	}

	return sandbox.ExtractFile(s, "main.exe")
}

func (CSharp) Run(ctx context.Context, s sandbox.Sandbox, binary sandbox.File, stdin io.Reader, stdout io.Writer, tl time.Duration, ml memory.Amount) (*sandbox.Status, error) {
	stat := sandbox.Status{}
	stat.Verdict = sandbox.VerdictXX

	if err := sandbox.CreateFile(s, binary); err != nil {
		return nil, err
	}

	rc := sandbox.RunConfig{
		MaxProcesses:     -1,
		InheritEnv:       true,
		Stdin:            stdin,
		Stdout:           stdout,
		TimeLimit:        tl,
		MemoryLimit:      ml,
		WorkingDirectory: s.Pwd(),
	}
	return s.Run(ctx, rc, "/usr/bin/mono", "main.exe")
}

func init() {
	language.DefaultStore.Register("csharp", CSharp{})
}
