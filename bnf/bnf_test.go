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
