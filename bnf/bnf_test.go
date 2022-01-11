package bnf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/suntong/gocc-grammars/bnf/lexer"
	"github.com/suntong/gocc-grammars/bnf/parser"
)

var testData = []string{
	// "// line comment",
	// "/* block comment */",
	"floatnumber ::=  pointfloat | exponentfloat ;",
	`<expr> ::= <term> "+" <expr>
        |  <term>
        ;`,
}

func TestPass(t *testing.T) {
	for _, ts := range testData {
		_, err := test([]byte(ts))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFail(t *testing.T) {
	_, err := test([]byte("a = b"))
	if err == nil {
		t.Fatal("unexpected parse error")
	} else {
		fmt.Printf("  Parsing failed as expected: %v\n", err)
	}
}

func TestFiles(t *testing.T) {
	files := []string{}
	err := filepath.Walk("test",
		func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				// fmt.Printf("  Testing file: %s (%d bytes)\n", path, f.Size())
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	// fmt.Printf("  Testing files: %v\n", files)

	p := parser.NewParser()
	for _, tf := range files {
		fmt.Printf("  Testing file: %s\n", tf)
		ts, err := ioutil.ReadFile(tf)
		if err != nil {
			panic(err)
		}
		s := lexer.NewLexer(ts)
		_, err = p.Parse(s)
		if err != nil {
			t.Error(err)
		}
	}
}

func test(src []byte) (astree interface{}, err error) {
	fmt.Printf("  Testing: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a
	}
	return
}

/*

$ go test -v .
=== RUN   TestPass
  Testing: floatnumber ::=  pointfloat | exponentfloat ;
  Testing: <expr> ::= <term> "+" <expr>
        |  <term>
        ;
--- PASS: TestPass (0.00s)
=== RUN   TestFail
  Testing: a = b ;
  Parsing failed as expected: 1:3: error: expected assign; got: unknown/invalid token "="
--- PASS: TestFail (0.00s)
=== RUN   TestFiles
  Testing file: test/calc.bnf
  Testing file: test/float.bnf
  Testing file: test/regexp-plg.bnf
  Testing file: test/shell.bnf
--- PASS: TestFiles (0.00s)
PASS
ok      github.com/suntong/gocc-grammars/bnf    0.001s

*/
