package source

import (
	"fmt"

	"github.com/dave/dst/decorator"

	dstlib "github.com/dave/dst"

	"github.com/go-gulfstream/gs/internal/schema"
)

var addonsFunc = map[string]func(dst, src *dstlib.File) error{
	// Events
	schema.EventsEventsAddon:         eventsEventsAddon,
	schema.EventsEventsEncodingAddon: eventsEventsEncodingAddon,
	schema.EventStateAddon:           eventStateAddon,
	schema.EventControllerAddon:      eventControllerAddon,
	schema.EventMutationAddon:        eventMutationAddon,
	schema.EventMutationImplAddon:    eventMutationImplAddon,
	schema.EventMutationTestAddon:    eventMutationTestAddon,

	// Commands
	schema.CommandsAddon:               commandsAddon,
	schema.CommandsEncodingAddon:       commandsEncodingAddon,
	schema.CommandStateAddon:           commandStateAddon,
	schema.CommandControllerAddon:      commandControllerAddon,
	schema.CommandMutationAddon:        commandMutationAddon,
	schema.CommandMutationImplAddon:    commandMutationImplAddon,
	schema.CommandMutationTestAddon:    commandMutationTestAddon,
	schema.CommandsEventsAddon:         commandsEventsAddon,
	schema.CommandsEventsEncodingAddon: commandsEventsEncodingAddon,
}

func ApplyAddon(dst *dstlib.File, addon string, addonSource []byte) error {
	if len(addonSource) == 0 {
		return nil
	}
	src, err := decorator.Parse(addonSource)
	if err != nil {
		return err
	}

	fn, found := addonsFunc[addon]
	if !found {
		return fmt.Errorf("source: ApplyAddon(%sAddon) => modificator not specified", addon)
	}
	return fn(dst, src)
}
