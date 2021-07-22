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
	commands := make([]CommandMutation, 0)
	for _, mc := range m.Mutations.Commands {
		if len(mc.Name) < 2 || len(mc.Command) < 2 {
			continue
		}
		mc.Event = strings.Title(mc.Event)
		mc.Command = strings.Title(mc.Command)
		mc.Name = strings.Title(mc.Name)
		commands = append(commands, mc)
	}
	m.Mutations.Commands = commands

	events := make([]EventMutation, 0)
	for _, ec := range m.Mutations.Events {
		if len(ec.Name) < 2 || len(ec.Event) < 2 {
			continue
		}
		ec.Name = strings.Title(ec.Name)
		ec.Event = strings.Title(ec.Event)
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
	return false
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
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Event   string `yaml:"event"`
}

type EventMutation struct {
	Package string `yaml:"pkg"`
	Name    string `yaml:"name"`
	Event   string `yaml:"event"`
}
