package source

import (
	"fmt"
	"io/ioutil"

	"github.com/dave/dst/decorator"

	dstlib "github.com/dave/dst"
	"github.com/go-gulfstream/gs/internal/schema"
)

var validationFunc = map[string]func(path string, src *dstlib.File) error{
	schema.CommandMutationAddon:     commandMutationValidation,
	schema.CommandMutationImplAddon: commandMutationImplValidation,
	schema.CommandControllerAddon:   commandMutationControllerValidation,

	schema.EventMutationAddon:     eventMutationValidation,
	schema.EventMutationImplAddon: eventMutationImplValidation,
	schema.EventControllerAddon:   eventMutationControllerValidation,

	schema.EventStateAddon: stateValidation,
}

func Validate(path string, m *schema.Manifest) error {
	files := schema.AddonFiles(path, m, true)
	for _, file := range files {
		fn, ok := validationFunc[file.Addon]
		if !ok {
			continue
		}
		data, err := ioutil.ReadFile(file.Path)
		if err != nil {
			return err
		}
		src, err := decorator.Parse(data)
		if err != nil {
			return err
		}
		if err := fn(file.Path, src); err != nil {
			return err
		}
	}
	return nil
}

func commandMutationValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findTypeSpecByName(src, commandMutationSelector)
	if err != nil {
		return fmt.Errorf("%v interface - filename: %s", err, path)
	}
	return nil
}

func commandMutationImplValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findTypeSpecByName(src, commandMutationImplSelector)
	if err != nil {
		return fmt.Errorf("%v type - filename: %s", err, path)
	}
	return nil
}

func commandMutationControllerValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findFuncDeclByName(src, makeCommandControllerSelector)
	if err != nil {
		return fmt.Errorf("%v filename: %s", err, path)
	}
	return nil
}

func eventMutationValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findTypeSpecByName(src, eventMutationSelector)
	if err != nil {
		return fmt.Errorf("%v interface - filename: %s", err, path)
	}
	return nil
}

func eventMutationImplValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findTypeSpecByName(src, eventMutationImplSelector)
	if err != nil {
		return fmt.Errorf("%v type - filename: %s", err, path)
	}
	return nil
}

func eventMutationControllerValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findFuncDeclByName(src, makeEventControllerSelector)
	if err != nil {
		return fmt.Errorf("%v filename: %s", err, path)
	}
	return nil
}

func stateValidation(path string, src *dstlib.File) error {
	if src.Name.Name != streamPackageSelector {
		return fmt.Errorf("source: invalid package name got %s, expected %s %s",
			src.Name.Name, streamPackageSelector, path)
	}
	_, err := findTypeSpecByName(src, stateSelector)
	if err != nil {
		return fmt.Errorf("%v type - filename: %s", err, path)
	}
	return nil
}
