package schema

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Project       project       `yaml:"project"`
	Contributor   contributor   `yaml:"contributor"`
	Mutations     mutations     `yaml:"mutations"`
	StreamStorage streamStorage `yaml:"storage"`
	Publisher     publisher     `yaml:"publisher"`
}

func (m *Manifest) Validate() error {
	return validator.New().Struct(m)
}

func (m *Manifest) MarshalBinary() ([]byte, error) {
	return yaml.Marshal(&m)
}

func (m *Manifest) UnmarshalBinary(data []byte) error {
	return yaml.Unmarshal(data, &m)
}

func (m *Manifest) Sanitize() {
	dupcommand := make(map[string]struct{})
	dupevent := make(map[string]struct{})
	commands := make([]CommandMutation, 0)
	for _, mc := range m.Mutations.Commands {
		if len(mc.Mutation) == 0 || !mc.Command.Validate() {
			continue
		}
		_, foundCommand := dupcommand[mc.Mutation]
		if foundCommand {
			continue
		}
		dupcommand[mc.Mutation] = struct{}{}
		mc.Event.Name = strings.Title(mc.Event.Name)
		mc.Event.Payload = strings.Title(mc.Event.Payload)
		mc.Command.Name = strings.Title(mc.Command.Name)
		mc.Command.Payload = strings.Title(mc.Command.Payload)
		mc.Mutation = strings.Title(mc.Mutation)
		commands = append(commands, mc)
	}
	m.Mutations.Commands = commands

	events := make([]EventMutation, 0)
	for _, ec := range m.Mutations.Events {
		if len(ec.Mutation) == 0 || ec.Event.Validate() {
			continue
		}
		_, foundEvent := dupcommand[ec.Mutation]
		if foundEvent {
			continue
		}
		dupevent[ec.Mutation] = struct{}{}
		ec.Mutation = strings.Title(ec.Mutation)
		ec.Event.Name = strings.Title(ec.Event.Name)
		ec.Event.Payload = strings.Title(ec.Event.Payload)
		ec.Package = strings.ToUpper(ec.Package)
		events = append(events, ec)
	}
	m.Mutations.Events = events
}

func MarshalBlankManifest() ([]byte, error) {
	emptyManifest := new(Manifest)
	emptyManifest.Project.Name = "myproject"
	emptyManifest.Project.CreatedAt = time.Now()
	emptyManifest.Project.GoModules = "github.com/go-gulfstream/myproject"
	emptyManifest.Mutations.Commands = []CommandMutation{{}}
	emptyManifest.Mutations.Events = []EventMutation{{}}
	data, err := emptyManifest.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	buf.WriteString("\n# available storage adapters:\n")
	for id, adapter := range storageAdapters {
		buf.WriteString(fmt.Sprintf("# id:%d, name: %s\n", id, adapter))
	}
	buf.WriteString("\n# available publisher adapters:\n")
	for id, adapter := range publisherAdapters {
		buf.WriteString(fmt.Sprintf("# id:%d, name: %s\n", id, adapter))
	}
	return buf.Bytes(), nil
}

type publisher struct {
	AdapterID publisherAdapter `yaml:"id"`
}

type mutations struct {
	Commands []CommandMutation `yaml:"from_commands"`
	Events   []EventMutation   `yaml:"from_events"`
}

func (m mutations) HasCommand() bool {
	return len(m.Commands) > 0
}

func (m mutations) HasEvents() bool {
	return len(m.Events) > 0
}

type project struct {
	Name      string    `yaml:"name" validate:"gte=3"`
	CreatedAt time.Time `yaml:"created_at"`
	GoModules string    `yaml:"go_modules" validate:"gte=3"`
}

type contributor struct {
	Author      string `yaml:"author"`
	Email       string `yaml:"email"`
	Description string `yaml:"description"`
}

type streamStorage struct {
	AdapterID     storageAdapter `yaml:"id"`
	EnableJournal bool           `yaml:"enable_journal"`
}

type CommandMutation struct {
	Mutation string  `yaml:"mutation"`
	Command  Command `yml:"command"`
	Event    Event   `yml:"event"`
	Create   bool    `yml:"create"`
}

type Command struct {
	Name    string `yml:"name"`
	Payload string `yml:"payload"`
}

func (c Command) Validate() bool {
	return len(c.Name) > 0
}

type Event struct {
	Name    string `yml:"name"`
	Payload string `yml:"payload"`
}

func (e Event) Validate() bool {
	return len(e.Name) > 0
}

type EventMutation struct {
	Mutation string `yaml:"mutation"`
	Package  string `yaml:"pkg"`
	Event    Event  `yaml:"event"`
	Create   bool   `yml:"create"`
}
