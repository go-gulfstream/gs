package schema

const (
	YesOp = "yes"
	NoOp  = "no"
)

type Manifest struct {
	Name            string        `yaml:"name"`
	PackageName     string        `yaml:"go_package_name"`
	StreamName      string        `yaml:"go_stream_name"`
	GoModules       string        `yaml:"go_modules"`
	EventsPkgName   string        `yaml:"go_events_pkg_name"`
	CommandsPkgName string        `yaml:"go_commands_pkg_name"`
	StreamPkgName   string        `yaml:"go_stream_pkg_name"`
	Description     string        `yaml:"description"`
	Mutations       mutations     `yaml:"mutations"`
	ImportEvents    []string      `yaml:"import_events"`
	StreamStorage   streamStorage `yaml:"storage_adapter"`
	StreamPublisher publisher     `yaml:"publisher_adapter"`
	Contributors    []Contributor `yaml:"contributors"`
}

type publisher struct {
	Name      string           `yaml:"name,omitempty"`
	AdapterID publisherAdapter `yaml:"id"`
}

type mutations struct {
	Commands []CommandMutation `yaml:"from_commands"`
	Events   []EventMutation   `yaml:"from_events"`
}

func (m mutations) HasCommands() bool {
	return len(m.Commands) > 0
}

func (m mutations) HasEvents() bool {
	return len(m.Events) > 0
}

type Contributor struct {
	Author string `yaml:"author"`
	Email  string `yaml:"email"`
}

type streamStorage struct {
	Name          string         `yaml:"name,omitempty"`
	AdapterID     storageAdapter `yaml:"id"`
	EnableJournal bool           `yaml:"enable_journal"`
}

type CommandMutation struct {
	Mutation string  `yaml:"name"`
	Command  Command `yaml:"in_command"`
	Event    Event   `yaml:"out_event"`
	Create   string  `yaml:"allow_create_stream,omitempty"`
	Delete   string  `yaml:"allow_delete_stream,omitempty"`
}

type EventMutation struct {
	Mutation string `yaml:"name"`
	InEvent  Event  `yaml:"in_event"`
	OutEvent Event  `yaml:"out_event"`
	Create   string `yaml:"allow_create_stream,omitempty"`
	Delete   string `yaml:"allow_delete_stream,omitempty"`
}

type Command struct {
	Name    string `yaml:"name"`
	Payload string `yaml:"payload"`
}

type Event struct {
	Name    string `yaml:"name"`
	Payload string `yaml:"payload"`
}
