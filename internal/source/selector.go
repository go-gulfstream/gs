package source

import (
	"fmt"

	dstlib "github.com/dave/dst"
)

const (
	streamPackageSelector         = "stream"
	commandMutationSelector       = "CommandMutation"
	eventMutationSelector         = "EventMutation"
	commandMutationImplSelector   = "commandMutation"
	projectionImplSelector        = "projection"
	eventMutationImplSelector     = "eventMutation"
	makeCommandControllerSelector = "MakeCommandControllers"
	makeEventControllerSelector   = "MakeEventControllers"
	newControllerSelector         = "NewController"
	stateSelector                 = "root"
	mutateFuncSelector            = "Mutate"
	constDeclSelector             = "const"
	typeDeclSelector              = "type"
	initDeclSelector              = "init"
	marshalFuncSelector           = "MarshalBinary"
	unmarshalFuncSelector         = "UnmarshalBinary"
)

func insertFuncDecl(a []dstlib.Decl, method *dstlib.FuncDecl, index int) []dstlib.Decl {
	return append(a[:index], append([]dstlib.Decl{method}, a[index:]...)...)
}

func insertExprStmt(a []dstlib.Stmt, expr *dstlib.ExprStmt, index int) []dstlib.Stmt {
	return append(a[:index], append([]dstlib.Stmt{expr}, a[index:]...)...)
}

func findRecvByName(decls []dstlib.Decl, recvName string) (index int) {
	for i, decl := range decls {
		fn, ok := decl.(*dstlib.FuncDecl)
		if !ok || fn.Recv == nil {
			continue
		}
		if len(fn.Recv.List) == 0 {
			continue
		}
		name := fn.Recv.List[0].Type.(*dstlib.StarExpr).X.(*dstlib.Ident).Name
		if name == recvName {
			index = i
		}
	}
	return
}

func findTypeSpecByName(file *dstlib.File, specName string) (spec *dstlib.TypeSpec, err error) {
	dstlib.Inspect(file, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.TypeSpec:
			if typ.Name.Name == specName {
				spec = typ
				return false
			}
		}
		return true
	})
	if spec == nil {
		err = fmt.Errorf("source: not found %s", specName)
	}
	return
}

func findFuncDeclByName(file *dstlib.File, name string) (res *dstlib.FuncDecl, err error) {
	dstlib.Inspect(file, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.FuncDecl:
			if typ.Name.Name == name || name == "" {
				res = typ
				return false
			}
		}
		return true
	})
	if res == nil {
		err = fmt.Errorf("source: can't find %s func declaration by name", name)
	}
	return
}

func findInterfaceMethods(typeSpec *dstlib.TypeSpec) ([]string, error) {
	spec, ok := typeSpec.Type.(*dstlib.InterfaceType)
	if !ok {
		return nil, fmt.Errorf("source: can't find interface methods")
	}
	result := make([]string, 0, len(spec.Methods.List))
	for _, field := range spec.Methods.List {
		result = append(result, field.Names[0].Name)
	}
	return result, nil
}

func findSwitchStmt(typeSpec *dstlib.FuncDecl) (switchStmt *dstlib.SwitchStmt, caseClause *dstlib.CaseClause, err error) {
	dstlib.Inspect(typeSpec, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.SwitchStmt:
			switchStmt = typ
		case *dstlib.CaseClause:
			caseClause = typ
			return false
		}
		return true
	})
	if switchStmt == nil {
		err = fmt.Errorf("source: switchStmt not found")
	}
	return
}

func findGenDeclByTok(file *dstlib.File, tok string) (decl *dstlib.GenDecl, err error) {
	dstlib.Inspect(file, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.GenDecl:
			if typ.Tok.String() == tok {
				decl = typ
				return false
			}
		}
		return true
	})
	if decl == nil {
		err = fmt.Errorf("source: genDecl not found")
	}
	return
}
