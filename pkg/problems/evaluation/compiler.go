package evaluation

import (
	"bytes"
	"fmt"
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/mraron/njudge/pkg/problems"
	"golang.org/x/net/context"
	"io"
	"testing/iotest"
)

type BytesSolution struct {
	lang language.Language
	src  []byte
}

func NewByteSolution(lang language.Language, src []byte) *BytesSolution {
	return &BytesSolution{
		lang: lang,
		src:  src,
	}
}

func (b *BytesSolution) GetLanguage() language.Language {
	return b.lang
}

func (b *BytesSolution) GetFile(ctx context.Context) (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewBuffer(b.src)), nil
}

type CompileCopyFile struct {
}

func (c CompileCopyFile) Compile(ctx context.Context, problem problems.Judgeable, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	f, err := solution.GetFile(ctx)
	return &problems.CompilationResult{
		CompiledFile:       f,
		CompilationMessage: "",
	}, err
}

type Compile struct{}

func (c Compile) Compile(ctx context.Context, problem problems.Judgeable, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	lang := solution.GetLanguage()

	f, err := solution.GetFile(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	// TODO add ctx to language
	stderr, bin := &bytes.Buffer{}, &bytes.Buffer{}
	stderrTruncater := iotest.TruncateWriter(stderr, 1<<16)
	if err := lang.Compile(sandbox, language.File{
		Name:   lang.DefaultFileName(),
		Source: f,
	}, bin, stderrTruncater, nil); err != nil {
		return &problems.CompilationResult{
			CompiledFile:       nil,
			CompilationMessage: stderr.String(),
		}, err
	}

	return &problems.CompilationResult{
		CompiledFile:       io.NopCloser(bin),
		CompilationMessage: stderr.String(),
	}, nil
}

type CompileCheckSupported struct{}

func (c CompileCheckSupported) Compile(ctx context.Context, problem problems.Judgeable, solution problems.Solution, sandbox sandbox.Sandbox) (*problems.CompilationResult, error) {
	lst, found := problem.Languages(), false

	for _, l := range lst {
		if l.Id() == solution.GetLanguage().Id() {
			found = true
		}
	}

	if !found {
		return &problems.CompilationResult{
			CompiledFile:       nil,
			CompilationMessage: "",
		}, fmt.Errorf("language is not supported: %s", solution.GetLanguage().Id())
	}

	return Compile{}.Compile(ctx, problem, solution, sandbox)
}
