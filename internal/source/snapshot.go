package source

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	path string
	file *ast.File
	pkg  *ast.Package
}

func (fi FileInfo) Inspect(fn func(node ast.Node) bool) {
	ast.Inspect(fi.file, fn)
}

func (fi FileInfo) Path() string {
	return fi.path
}

func (fi FileInfo) File() *ast.File {
	return fi.file
}

func (fi FileInfo) Package() *ast.Package {
	return fi.pkg
}

type Snapshot struct {
	rootPath string
	byNames  map[string]*ast.File
	byPos    map[token.Pos]*ast.File
	packages map[*ast.File]*ast.Package
	fileSet  *token.FileSet
}

func NewSnapshot(path string) (*Snapshot, error) {
	idx := &Snapshot{
		rootPath: path,
		fileSet:  token.NewFileSet(),
		byNames:  make(map[string]*ast.File),
		byPos:    make(map[token.Pos]*ast.File),
		packages: make(map[*ast.File]*ast.Package),
	}
	if err := idx.scan(); err != nil {
		return nil, err
	}
	return idx, nil
}

func (i *Snapshot) TotalFiles() int {
	return len(i.byNames)
}

func (i *Snapshot) Walk(walkFn func(info FileInfo) error) error {
	for path, file := range i.byNames {
		pkg, found := i.packages[file]
		if !found {
			continue
		}
		fi := FileInfo{
			path: path,
			file: file,
			pkg:  pkg,
		}
		if err := walkFn(fi); err != nil {
			return err
		}
	}
	return nil
}

func (i *Snapshot) scan() error {
	indexPath := func(path string) error {
		packages, err := parser.ParseDir(i.fileSet, path, func(info fs.FileInfo) bool {
			return true
		}, parser.ParseComments)
		if err != nil {
			return err
		}
		for _, pkg := range packages {
			for path, file := range pkg.Files {
				i.byPos[file.Pos()] = file
				i.byNames[path] = file
				i.packages[file] = pkg
			}
		}
		return nil
	}
	if err := indexPath(i.rootPath); err != nil {
		return err
	}
	return filepath.Walk(i.rootPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() || dirStartsWithDot(path) {
			return nil
		}
		return indexPath(path)
	})
}

func dirStartsWithDot(dir string) bool {
	if strings.HasPrefix(dir, ".") {
		return true
	}
	paths := strings.Split(dir, "/")
	for _, path := range paths {
		if strings.HasPrefix(path, ".") {
			return true
		}
	}
	return false
}
