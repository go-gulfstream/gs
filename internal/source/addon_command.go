package source

import (
	"fmt"

	"github.com/dave/dst"
	dstlib "github.com/dave/dst"
)

func commandsEventsEncodingAddon(dst *dstlib.File, src *dstlib.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	// from template
	srcInitFn, err := findFuncDeclByName(src, initDeclSelector)
	if err != nil {
		return err
	}
	srcMarshalFn, err := findFuncDeclByName(src, marshalFuncSelector)
	if err != nil {
		// no payload
		return nil
	}
	srcUnmarshalFn, err := findFuncDeclByName(src, unmarshalFuncSelector)
	if err != nil {
		// no payload
		return nil
	}

	srcMarshalFn.Decorations().Before = dstlib.EmptyLine
	srcUnmarshalFn.Decorations().Before = dstlib.EmptyLine

	if len(srcInitFn.Body.List) > 0 {
		dstInitFn, err := findFuncDeclByName(dst, initDeclSelector)
		if err != nil {
			dst.Decls = append(dst.Decls, srcInitFn)
		} else {
			dstInitFn.Body.List = append(dstInitFn.Body.List, srcInitFn.Body.List[0])
		}
	}

	dst.Decls = append(dst.Decls, srcMarshalFn)
	dst.Decls = append(dst.Decls, srcUnmarshalFn)

	return nil
}

func commandsEventsAddon(dst *dstlib.File, src *dstlib.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	// from template
	srcConstDecl, err := findGenDeclByTok(src, constDeclSelector)
	if err != nil {
		return err
	}
	dstConstDecl, err := findGenDeclByTok(dst, constDeclSelector)
	if err != nil {
		dst.Decls = append(dst.Decls, srcConstDecl)
	} else {
		dstConstDecl.Specs = append(dstConstDecl.Specs, srcConstDecl.Specs[0])
	}

	srcType, err := findGenDeclByTok(src, typeDeclSelector)
	if err == nil {
		dst.Decls = append(dst.Decls, srcType)
	}

	return nil
}

func commandStateAddon(dst *dstlib.File, src *dstlib.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	applyFunc := src.Decls[2]
	applyFunc.Decorations().Before = dstlib.EmptyLine
	dst.Decls = append(dst.Decls, applyFunc)

	// from template
	srcMutateFunc, err := findFuncDeclByName(src, mutateFuncSelector)
	if err != nil {
		return err
	}
	srcSwitch, srcCase, err := findSwitchStmt(srcMutateFunc)
	if err != nil {
		return fmt.Errorf("%v from template", err)
	}
	srcSwitch.Decorations().Before = dstlib.None
	srcSwitch.Decorations().After = dstlib.None

	dstMutateFunc, err := findFuncDeclByName(dst, mutateFuncSelector)
	if err != nil {
		return err
	}
	dstSwitch, _, err := findSwitchStmt(dstMutateFunc)
	if err != nil {
		dstMutateFunc.Body.List = append(dstMutateFunc.Body.List, srcSwitch)
	} else {
		dstSwitch.Body.List = append(dstSwitch.Body.List, srcCase)
	}

	return nil
}

func commandsAddon(dst *dstlib.File, src *dstlib.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	// from template
	srcConstDecl, err := findGenDeclByTok(src, constDeclSelector)
	if err != nil {
		return err
	}
	dstConstDecl, err := findGenDeclByTok(dst, constDeclSelector)
	if err != nil {
		dst.Decls = append(dst.Decls, srcConstDecl)
	} else {
		dstConstDecl.Specs = append(dstConstDecl.Specs, srcConstDecl.Specs[0])
	}

	srcType, err := findGenDeclByTok(src, typeDeclSelector)
	if err == nil {
		dst.Decls = append(dst.Decls, srcType)
	}

	fn, err := findFuncDeclByName(src, "")
	if err != nil {
		return err
	}

	fn.Decorations().Before = dstlib.EmptyLine
	dst.Decls = append(dst.Decls, fn)

	return nil
}

func commandsEncodingAddon(dst *dstlib.File, src *dstlib.File) error {
	if len(src.Imports) > 0 {
		dst.Imports = append(dst.Imports, src.Imports...)
	}

	// from template
	srcInitFn, err := findFuncDeclByName(src, initDeclSelector)
	if err != nil {
		return err
	}
	srcMarshalFn, err := findFuncDeclByName(src, marshalFuncSelector)
	if err != nil {
		// no payload
		return nil
	}
	srcUnmarshalFn, err := findFuncDeclByName(src, unmarshalFuncSelector)
	if err != nil {
		// no payload
		return nil
	}

	srcMarshalFn.Decorations().Before = dstlib.EmptyLine
	srcUnmarshalFn.Decorations().Before = dstlib.EmptyLine

	if len(srcInitFn.Body.List) > 0 {
		dstInitFn, err := findFuncDeclByName(dst, initDeclSelector)
		if err != nil {
			dst.Decls = append(dst.Decls, srcInitFn)
		} else {
			dstInitFn.Body.List = append(dstInitFn.Body.List, srcInitFn.Body.List[0])
		}
	}

	dst.Decls = append(dst.Decls, srcMarshalFn)
	dst.Decls = append(dst.Decls, srcUnmarshalFn)

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

func commandMutationTestAddon(dst *dst.File, src *dst.File) error {
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
	dst.Decls = append(dst.Decls, method)

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
	index := findRecvByName(dst.Decls, commandMutationImplSelector)

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

func commandMutationProjectionAddon(dst *dst.File, src *dst.File) error {
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

func commandMutationImplProjectionAddon(dst *dst.File, src *dst.File) error {
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
	index := findRecvByName(dst.Decls, projectionImplSelector)

	if index == 0 {
		dst.Decls = append(dst.Decls, method)
	} else {
		dst.Decls = insertFuncDecl(dst.Decls, method, index)
	}

	return nil
}

func commandControllerProjectionAddon(dst *dst.File, src *dst.File) error {
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

	method.Decorations().Before = dstlib.EmptyLine

	dst.Decls = append(dst.Decls, method)
	funcDecl, err := findFuncDeclByName(dst, newControllerSelector)
	if err != nil {
		return err
	}
	if len(funcDecl.Body.List) > 0 {
		prevIdx := len(funcDecl.Body.List) - 1
		funcDecl.Body.List = insertExprStmt(funcDecl.Body.List, exprStmt, prevIdx)
	} else {
		funcDecl.Body.List = append(funcDecl.Body.List, exprStmt)
	}
	return nil
}

func projectionTestAddon(dst *dst.File, src *dst.File) error {
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
	dst.Decls = append(dst.Decls, method)

	return nil
}
