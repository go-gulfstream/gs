package source

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"

	dstlib "github.com/dave/dst"

	"github.com/go-gulfstream/gs/internal/schema"
)

var addonsFunc = map[string]func(dst, src *dstlib.File) error{
	schema.EventsAddon:              eventsAddon,
	schema.StateAddon:               stateAddon,
	schema.EventControllerAddon:     eventControllerAddon,
	schema.EventMutationAddon:       eventMutationAddon,
	schema.CommandsAddon:            commandsAddon,
	schema.CommandControllerAddon:   commandControllerAddon,
	schema.CommandMutationAddon:     commandMutationAddon,
	schema.CommandMutationImplAddon: commandMutationImplAddon,
}

const (
	commandMutationSelector       = "commandMutation"
	makeCommandControllerSelector = "MakeCommandControllers"
)

func ModifyFromAddon(dst *dstlib.File, addon string, addonSource []byte) error {
	if len(addonSource) == 0 {
		return nil
	}
	if schema.CommandControllerAddon == addon {
		fmt.Println(string(addonSource))
	}
	src, err := decorator.Parse(addonSource)
	if err != nil {
		return err
	}
	fn, found := addonsFunc[addon]
	if !found {
		return fmt.Errorf("source: ModifyFromAddon(%sAddon) => modificator not specified", addon)
	}
	return fn(dst, src)
}

func eventsAddon(dst *dstlib.File, src *dstlib.File) error {
	return nil
}

func stateAddon(dst *dstlib.File, src *dstlib.File) error {
	return nil
}

func eventControllerAddon(dst *dstlib.File, src *dstlib.File) error {
	return nil
}

func eventMutationAddon(dst *dstlib.File, src *dstlib.File) error {
	return nil
}

func commandsAddon(dst *dstlib.File, src *dstlib.File) error {
	return nil
}

func commandControllerAddon(dst *dst.File, src *dst.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	var exprStmt *dstlib.ExprStmt
	var method *dstlib.FuncDecl
	dstlib.Inspect(src, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.FuncDecl:
			if typ.Name.Name == "render" {
				exprStmt = typ.Body.List[0].(*dstlib.ExprStmt)
			} else {
				method = typ
			}
			return false
		}
		return true
	})

	method.Decorations().After = dstlib.EmptyLine
	dst.Decls = append(dst.Decls, method)
	funcDecl, err := findFuncDeclByName(dst, makeCommandControllerSelector)
	if err != nil {
		return err
	}
	funcDecl.Body.List = append(funcDecl.Body.List, exprStmt)

	return nil
}

func commandMutationImplAddon(dst *dst.File, src *dst.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	var method *dstlib.FuncDecl
	dstlib.Inspect(src, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.FuncDecl:
			method = typ
			return false
		}
		return true
	})

	method.Decorations().After = dstlib.EmptyLine
	index := findRecvByName(dst.Decls, commandMutationSelector)

	if index == 0 {
		dst.Decls = append(dst.Decls, method)
	} else {
		dst.Decls = insertFuncDecl(dst.Decls, method, index)
	}

	return nil
}

func commandMutationAddon(dst *dst.File, src *dst.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	var method *dstlib.Field
	dstlib.Inspect(src, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.InterfaceType:
			method = typ.Methods.List[0]
			return false
		}
		return true
	})

	dstlib.Inspect(dst, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.InterfaceType:
			typ.Methods.List = append(typ.Methods.List, method)
			return false
		}
		return true
	})

	return nil
}

func insertFuncDecl(a []dstlib.Decl, method *dstlib.FuncDecl, index int) []dstlib.Decl {
	return append(a[:index], append([]dstlib.Decl{method}, a[index:]...)...)
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

func findFuncDeclByName(file *dstlib.File, name string) (res *dstlib.FuncDecl, err error) {
	dstlib.Inspect(file, func(node dstlib.Node) bool {
		switch typ := node.(type) {
		case *dstlib.FuncDecl:
			if typ.Name.Name == name {
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
