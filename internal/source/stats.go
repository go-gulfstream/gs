package source

import (
	"io/ioutil"

	dstlib "github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/go-gulfstream/gs/internal/schema"
)

type Info struct {
	commandMutations map[string]struct{}
	eventMutations   map[string]struct{}
}

func (i *Info) CommandMutationExists(name string) bool {
	_, found := i.commandMutations[name]
	return found
}

func (i *Info) EventMutationExists(name string) bool {
	_, found := i.eventMutations[name]
	return found
}

var statsFunc = map[string]func(path string, src *dstlib.File, info *Info) error{
	schema.CommandMutationAddon: commandMutationInfo,
	schema.EventMutationAddon:   eventMutationInfo,
}

func Stats(path string, m *schema.Manifest) (*Info, error) {
	info := &Info{
		commandMutations: make(map[string]struct{}),
		eventMutations:   make(map[string]struct{}),
	}
	files := schema.AddonFiles(path, m, true)
	for _, file := range files {
		fn, ok := statsFunc[file.Addon]
		if !ok {
			continue
		}
		data, err := ioutil.ReadFile(file.Path)
		if err != nil {
			return nil, err
		}
		src, err := decorator.Parse(data)
		if err != nil {
			return nil, err
		}
		if err := fn(file.Path, src, info); err != nil {
			return nil, err
		}
	}
	return info, nil
}

func commandMutationInfo(_ string, src *dstlib.File, info *Info) error {
	spec, err := findTypeSpecByName(src, commandMutationSelector)
	if err != nil {
		return err
	}
	methods, err := findInterfaceMethods(spec)
	if err != nil {
		return nil
	}
	for _, method := range methods {
		info.commandMutations[method] = struct{}{}
	}
	return nil
}

func eventMutationInfo(_ string, src *dstlib.File, info *Info) error {
	spec, err := findTypeSpecByName(src, eventMutationSelector)
	if err != nil {
		return err
	}
	methods, err := findInterfaceMethods(spec)
	if err != nil {
		return nil
	}
	for _, method := range methods {
		info.eventMutations[method] = struct{}{}
	}
	return nil
}
