package format

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"golang.org/x/tools/imports"
)

func Source(filepath string, data []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", data, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	formatted, err := gofmtFile(file, fset)
	if err != nil {
		return nil, err
	}
	return imports.Process(filepath, formatted, &imports.Options{})
}

func gofmtFile(f *ast.File, fset *token.FileSet) ([]byte, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
