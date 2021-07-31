package format

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"

	"golang.org/x/tools/imports"
)

func Source(filepath string, data []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", data, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	fixImport(file)

	formatted, err := gofmtFile(file, fset)
	if err != nil {
		return nil, err
	}
	return imports.Process(filepath, formatted, &imports.Options{
		Comments: true,
	})
}

func File(filepath string, file *ast.File) ([]byte, error) {
	fset := token.NewFileSet()

	fixImport(file)

	formatted, err := gofmtFile(file, fset)
	if err != nil {
		return nil, err
	}
	return imports.Process(filepath, formatted, nil)
}

func fixImport(file *ast.File) {
	stats := make(map[string]int)
	for _, importSpec := range file.Imports {
		stats[importSpec.Path.Value]++
	}
	astutil.Apply(file, func(c *astutil.Cursor) bool {
		spec, ok := c.Node().(*ast.ImportSpec)
		if !ok {
			return true
		}
		n, ok := stats[spec.Path.Value]
		if !ok {
			return true
		}
		if n > 1 && spec.Path.Kind == token.STRING {
			deleteImport(file, spec, c)
			stats[spec.Path.Value]--
		}
		return true
	}, nil)
}

func deleteImport(file *ast.File, spec *ast.ImportSpec, c *astutil.Cursor) {
	imports := file.Imports[:0]
	for _, imp := range file.Imports {
		if imp != spec {
			imports = append(imports, imp)
			continue
		}
	}
	c.Delete()
	tmp := file.Imports[len(imports):]
	for i := range tmp {
		tmp[i] = nil
	}
	file.Imports = imports
}

func gofmtFile(f *ast.File, fset *token.FileSet) ([]byte, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func fileToByte(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var output []byte
	buffer := bytes.NewBuffer(output)
	if err := printer.Fprint(buffer, fset, file); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
