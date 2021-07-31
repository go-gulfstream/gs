package source

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/go-gulfstream/gs/internal/format"
)

var files = map[string]*ast.File{}

func Load(filename string) (*ast.File, error) {
	file, found := files[filename]
	if found {
		return file, nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	file, err = parseSource(data)
	if err != nil {
		return nil, err
	}
	files[filename] = file
	return file, nil
}

func Save() error {
	for filename, astFile := range files {
		data, err := format.File(filename, astFile)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(filename, data, 0755); err != nil {
			return err
		}
	}
	return nil
}

func parseSource(src []byte) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", src, parser.ParseComments)
}
