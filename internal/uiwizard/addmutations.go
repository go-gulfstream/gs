package uiwizard

import (
	"fmt"

	"github.com/go-gulfstream/gs/internal/schema"
)

const (
	commandMutation = "commandMutation"
	eventMutation   = "eventMutation"
)

type Mutation struct {
	importEvents     map[string]struct{}
	commandMutations []schema.CommandMutation
	eventMutations   []schema.EventMutation
}

type group struct {
	Name    string
	Payload string
	Create  string
	Delete  string
}

func NewMutation() *Mutation {
	return &Mutation{
		importEvents:     make(map[string]struct{}),
		commandMutations: make([]schema.CommandMutation, 0),
		eventMutations:   make([]schema.EventMutation, 0),
	}
}

func (a *Mutation) HasChanges() bool {
	return len(a.eventMutations) > 0 || len(a.commandMutations) > 0
}

func (a *Mutation) Apply(m *schema.Manifest) {
	importEvents := make([]string, 0, len(a.importEvents))
	for ie := range a.importEvents {
		importEvents = append(importEvents, ie)
	}

	m.Mutations.Commands = append(m.Mutations.Commands, a.commandMutations...)
	m.Mutations.Events = append(m.Mutations.Events, a.eventMutations...)
	m.ImportEvents = append(m.ImportEvents, importEvents...)

	schema.SanitizeManifest(m)
}

func (a *Mutation) handleCommandMutation(prefix string) error {
	name, err := inputControl("mutation", "", true)
	if err != nil {
		return err
	}
	sectionControl(prefix + " incommand")
	commandInfo, err := a.inputMutationGroup(true)
	if err != nil {
		return err
	}
	sectionControl(prefix + " outevent")
	eventInfo, err := a.inputMutationGroup(false)
	if err != nil {
		return err
	}
	a.commandMutations = append(a.commandMutations, schema.CommandMutation{
		Mutation: name,
		Command: schema.Command{
			Name:    commandInfo.Name,
			Payload: commandInfo.Payload,
		},
		Event: schema.Event{
			Name:    eventInfo.Name,
			Payload: eventInfo.Payload,
		},
	})
	return nil
}

func (a *Mutation) handleEventMutation(prefix string) error {
	name, err := inputControl("mutation", "", true)
	if err != nil {
		return err
	}
	sectionControl(prefix + " inevent")
	inEventInfo, err := a.inputMutationGroup(true)
	if err != nil {
		return err
	}
	sectionControl(prefix + " outevent")
	outEventInfo, err := a.inputMutationGroup(false)
	if err != nil {
		return err
	}
	pkg, err := inputControl("package", "", false)
	if err != nil {
		return err
	}
	a.importEvents[pkg] = struct{}{}
	a.eventMutations = append(a.eventMutations, schema.EventMutation{
		Mutation: name,
		InEvent: schema.Event{
			Name:    inEventInfo.Name,
			Payload: inEventInfo.Payload,
		},
		OutEvent: schema.Event{
			Name:    outEventInfo.Name,
			Payload: outEventInfo.Payload,
		},
	})
	return nil
}

func (a *Mutation) inputMutationGroup(withOpts bool) (group, error) {
	name, err := inputControl("name", "", true)
	if err != nil {
		return group{}, err
	}
	payload, err := inputControl("payload", "", false)
	if err != nil {
		return group{}, err
	}
	g := group{
		Name:    name,
		Payload: payload,
	}
	if withOpts {
		opts, err := selectControl("options", []selectItem{
			{
				ID:   "skip",
				Name: "Default",
				Help: "Default mode",
			},
			{
				ID:   "create",
				Name: "Allow create",
				Help: "Allows to create a stream if it does not exist",
			},
			{
				ID:   "delete",
				Name: "Allow delete",
				Help: "Allows to hard delete the stream from storage",
			},
		})
		if err != nil {
			return group{}, err
		}
		switch opts.ID {
		default:
			g.Create = "no"
			g.Delete = "no"
		case "crate":
			g.Create = "yes"
		case "delete":
			g.Delete = "yes"
		}
	}
	return g, nil
}

func (a *Mutation) selectMutationType() (selectItem, error) {
	return selectControl(
		"Select the type of stream mutation",
		[]selectItem{
			{
				ID:   commandMutation,
				Name: "Command mutation",
				Help: "Mutation from a command",
			},
			{
				ID:   eventMutation,
				Name: "Event mutation",
				Help: "Mutation from an event",
			},
		},
	)
}

func (a *Mutation) Run() error {
	var mutationNum int
	for {
		if mutationNum > 0 {
			lineControl()
		}
		item, err := a.selectMutationType()
		if err != nil {
			return err
		}
		section := fmt.Sprintf("%d.%s", mutationNum, item.Name)
		sectionControl(section)
		switch item.ID {
		case commandMutation:
			err = a.handleCommandMutation(section)
		case eventMutation:
			err = a.handleEventMutation(section)
		}
		if err != nil {
			return err
		}
		next, err := confirmControl("add more?")
		if err != nil {
			return err
		}
		if !next {
			break
		}
		mutationNum++
	}
	return nil
}
