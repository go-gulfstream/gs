package schema

import (
	"bytes"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Project       project       `yaml:"project"`
	Contributor   contributor   `yaml:"contributor"`
	Mutations     mutations     `yaml:"mutations"`
	StreamStorage streamStorage `yaml:"storage"`
	Publisher     publisher     `yaml:"publisher"`
}

func (m *Manifest) MarshalBinary() ([]byte, error) {
	return yaml.Marshal(&m)
}

func (m *Manifest) UnmarshalBinary(data []byte) error {
	return yaml.Unmarshal(data, &m)
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
	AdapterName string           `yaml:"adapter_name"`
	AdapterID   publisherAdapter `yaml:"adapter_id"`
}

type mutations struct {
	Commands []CommandMutation `yaml:"from_commands"`
	Events   []EventMutation   `yaml:"from_events"`
}

type project struct {
	Name      string    `yaml:"name"`
	CreatedAt time.Time `yaml:"created_at"`
	GoModules string    `yaml:"go_modules"`
}

type contributor struct {
	Author      string `yaml:"author"`
	Email       string `yaml:"email"`
	Description string `yaml:"description"`
}

type streamStorage struct {
	AdapterID     storageAdapter `yaml:"adapter_id"`
	AdapterName   string         `yaml:"adapter_name"`
	EnableJournal bool           `yaml:"enable_journal"`
}

type CommandMutation struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Event   string `yaml:"event"`
}

type EventMutation struct {
	Name  string `yaml:"name"`
	Event string `yaml:"event"`
}
