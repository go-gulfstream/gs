package source

import (
	"fmt"
	"go/ast"

	"github.com/go-gulfstream/gs/internal/schema"
)

var addons = map[string]func(dst, src *ast.File) error{
	schema.EventsAddon:            appendEventsAddon,
	schema.StateAddon:             appendStateAddon,
	schema.EventControllerAddon:   appendEventControllerAddon,
	schema.EventMutationAddon:     appendEventMutationAddon,
	schema.CommandsAddon:          appendCommandsAddon,
	schema.CommandControllerAddon: appendCommandControllerAddon,
	schema.CommandMutationAddon:   appendCommandMutationAddon,
}

func Append(dst *ast.File, addon string, source []byte) (err error) {
	src, err := parseSource(source)
	if err != nil {
		return err
	}
	appendFn, ok := addons[addon]
	if !ok {
		return fmt.Errorf("source: Append(%s) => addon not found", addon)
	}
	return appendFn(dst, src)
}

func appendEventsAddon(dst *ast.File, src *ast.File) error {
	return nil
}

func appendStateAddon(dst *ast.File, src *ast.File) error {
	return nil
}

func appendEventControllerAddon(dst *ast.File, src *ast.File) error {
	return nil
}

func appendEventMutationAddon(dst *ast.File, src *ast.File) error {
	return nil
}

func appendCommandsAddon(dst *ast.File, src *ast.File) error {
	return nil
}

func appendCommandControllerAddon(dst *ast.File, src *ast.File) error {
	return nil
}

func appendCommandMutationAddon(dst *ast.File, src *ast.File) error {
	return nil
}
