package format

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func Source(data []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", data, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return gofmtFile(file, fset)
}

func gofmtFile(f *ast.File, fset *token.FileSet) ([]byte, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
