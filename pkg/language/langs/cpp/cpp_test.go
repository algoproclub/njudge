package cpp

import (
	"github.com/mraron/njudge/pkg/language/sandbox"
	"testing"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

const (
	aplusb = `#include<iostream>
using namespace std;
int main() {
	int a,b;
	cin>>a>>b;
	cout<<a+b<<"\n";
}`
	compilerError = `#include<lol>
lmao;
int main(a,b,c);
`
	print = `#include<iostream>
using namespace std;
int main() {
	cout<<"Hello world";
	return 0;
}`
	timelimitExceeded = `#include<iostream>
using namespace std;
int main() {
	int n=0;
	while(1) n++; 
}`
	runtimeError = `#include<iostream>
using namespace std;
void dfs(int x){
	dfs(x+1);
	if(x==int(1e9)) cerr<<"lel\n";
}
int main() {
	dfs(-10000);
}`
	runtimeErrorDiv0 = `#include<iostream>
using namespace std;
int main() {
	cerr<<(1/0);
}`
	longSleep = `#include<unistd.h>
int main() {
	sleep(20);
}
`
	shortSleep = `#include<unistd.h>
int main() {
	sleep(1);
}
`
)

func TestCompileAndRun(t *testing.T) {
	for _, test := range []language.LanguageTest{
		{sandbox.NewDummy(), latest, aplusb, language.VERDICT_OK, "1 2", "3\n", 1 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, compilerError, language.VERDICT_CE, "", "", 1 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, print, language.VERDICT_OK, "", "Hello world", 1 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, timelimitExceeded, language.VERDICT_TL, "", "", 100 * time.Millisecond, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, runtimeError, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, runtimeErrorDiv0, language.VERDICT_RE, "", "", 1000 * time.Millisecond, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, longSleep, language.VERDICT_TL, "", "", 2 * time.Second, 128 * 1024 * 1024},
		{sandbox.NewDummy(), latest, shortSleep, language.VERDICT_OK, "", "", 2 * time.Second, 128 * 1024 * 1024},
	} {
		test.Run(t)
	}
}
