package cpp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mraron/njudge/pkg/language/sandbox"

	"github.com/spf13/afero"
)

// SkipAutoCompile makes AutoCompile a no-op. It is meant for processes (like the web server) that
// only need problem metadata and may not have a C++ compiler available.
var SkipAutoCompile bool

// AutoCompile compiles src into dst unless dst already is a non-empty executable.
//
// The binary is written to a temporary file next to dst and renamed into place only after a successful
// compile, so concurrent readers (or other processes compiling the same file) never see a partial,
// empty or non-executable dst, and a failed compile leaves any existing dst untouched.
func AutoCompile(ctx context.Context, fs afero.Fs, s sandbox.Sandbox, workDir, src, dst string) error {
	if src == "" || SkipAutoCompile {
		return nil
	}

	if st, err := fs.Stat(dst); err == nil && st.Size() != 0 && st.Mode()&0100 == 0100 {
		return nil
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	file, err := fs.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := s.Init(ctx); err != nil {
		return err
	}
	defer func() {
		_ = s.Cleanup(ctx)
	}()

	contents, err := afero.ReadFile(fs, src)
	if err != nil {
		return err
	}

	var headers []sandbox.File
	for _, header := range ExtractHeaderNames(fs, workDir, contents) {
		headerContents, err := afero.ReadFile(fs, filepath.Join(workDir, header))
		if err != nil {
			return err
		}

		headers = append(headers, sandbox.File{
			Name:   header,
			Source: io.NopCloser(bytes.NewReader(headerContents)),
		})
	}

	var errorStream bytes.Buffer
	resBinary, err := Std17.Compile(context.TODO(), s, sandbox.File{
		Name:   filepath.Base(src),
		Source: file,
	}, &errorStream, headers)
	if err != nil {
		return errors.Join(err, fmt.Errorf("compile error: %v", errorStream.String()))
	}

	tmp, err := afero.TempFile(fs, filepath.Dir(dst), filepath.Base(dst)+".tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()

	if _, err = io.Copy(tmp, resBinary.Source); err != nil {
		return errors.Join(err, tmp.Close(), fs.Remove(tmpName))
	}
	if err = tmp.Close(); err != nil {
		return errors.Join(err, fs.Remove(tmpName))
	}
	if err = fs.Chmod(tmpName, 0755); err != nil {
		return errors.Join(err, fs.Remove(tmpName))
	}
	if err = fs.Rename(tmpName, dst); err != nil {
		return errors.Join(err, fs.Remove(tmpName))
	}

	return nil
}
