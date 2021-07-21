package schema

import (
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
