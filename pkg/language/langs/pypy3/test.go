package pypy3

import (
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	PYTHON3_aplusb = `inp = input().split(' ')
a,b = int(inp[0]), int(inp[1]) 
print(a+b)`
	PYTHON3_ce    = `inp = input(()`
	PYTHON3_print = `print("Hello world")`
	PYTHON3_tl    = `x = 0
while True:
	x = x+1`
	PYTHON3_re = `x = [1,2,3]
print(x[4])`
	PYTHON3_rediv0 = `print(1/0)`
)

func (p pypy3) Test(t *testing.T, s sandbox.Sandbox) error {
	for _, test := range []language.Test{
		{"pypy3_aplusb", p, PYTHON3_aplusb, sandbox.VerdictOK, "1 2", "3\n", 1 * time.Second, 128 * memory.MiB},
		{"pypy3_re", p, PYTHON3_ce, sandbox.VerdictRE, "", "", 1 * time.Second, 128 * memory.MiB},
		{"pypy3_print", p, PYTHON3_print, sandbox.VerdictOK, "", "Hello world\n", 1 * time.Second, 128 * memory.MiB},
		{"pypy3_tl", p, PYTHON3_tl, sandbox.VerdictTL, "", "", 100 * time.Millisecond, 128 * memory.MiB},
		{"pypy3_re", p, PYTHON3_re, sandbox.VerdictRE, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
		{"pypy3_rediv0", p, PYTHON3_rediv0, sandbox.VerdictRE, "", "", 1000 * time.Millisecond, 128 * memory.MiB},
	} {
		t.Run(test.Name, func(t *testing.T) {
			if err := test.Run(s); err != nil {
				t.Error(err)
			}
		})
	}

	return nil
}
