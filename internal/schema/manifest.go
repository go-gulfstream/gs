package schema

import (
	"time"

	"github.com/go-gulfstream/gs/internal/strutil"
)

const (
	YesOp = "yes"
	NoOp  = "no"
)

type Manifest struct {
	Name            string        `yaml:"name"`
	PackageName     string        `yaml:"go_package_name"`
	StreamName      string        `yaml:"go_stream_name"`
	GoModules       string        `yaml:"go_modules"`
	GoVersion       string        `yaml:"go_version"`
	GoGetPackages   []string      `yaml:"go_get_packages"`
	EventsPkgName   string        `yaml:"go_events_pkg_name"`
	CommandsPkgName string        `yaml:"go_commands_pkg_name"`
	StreamPkgName   string        `yaml:"go_stream_pkg_name"`
	Description     string        `yaml:"description"`
	Mutations       Mutations     `yaml:"mutations"`
	ImportEvents    []string      `yaml:"import_events"`
	StreamStorage   streamStorage `yaml:"storage_adapter"`
	StreamPublisher publisher     `yaml:"publisher_adapter"`
	CommandBus      commandBus    `yaml:"commandbus_adapter"`
	Contributors    []Contributor `yaml:"contributors"`
	CreatedAt       time.Time     `yaml:"created_at"`
	UpdatedAt       time.Time     `yaml:"updated_at"`
}

func New() *Manifest {
	return new(Manifest)
}

func (m *Manifest) SetPublisherFromString(name string) {
	switch name {
	case DefaultName:
		m.StreamPublisher = publisher{
			Name:      DefaultStreamPublisher.String(),
			AdapterID: DefaultStreamPublisher,
		}
	case KafkaStreamPublisherName:
		m.StreamPublisher = publisher{
			Name:      KafkaStreamPublisher.String(),
			AdapterID: KafkaStreamPublisher,
		}
	case ConnectorStreamPublisherName:
		m.StreamPublisher = publisher{
			Name:      ConnectorStreamPublisher.String(),
			AdapterID: ConnectorStreamPublisher,
		}
	}
}

func (m *Manifest) SetStreamStorageFromString(name string, journal bool) {
	switch name {
	case DefaultName:
		m.StreamStorage = streamStorage{
			Name:          DefaultStreamStorage.String(),
			AdapterID:     DefaultStreamStorage,
			EnableJournal: journal,
		}
	case RedisStreamStorageName:
		m.StreamStorage = streamStorage{
			Name:          RedisStreamStorage.String(),
			AdapterID:     RedisStreamStorage,
			EnableJournal: journal,
		}
	case PostgresStreamStorageName:
		m.StreamStorage = streamStorage{
			Name:          PostgresStreamStorage.String(),
			AdapterID:     PostgresStreamStorage,
			EnableJournal: journal,
		}
	}
}

func (m *Manifest) SetCommandBusFromString(name string) {
	switch name {
	case HTTPCommandBusName:
		m.CommandBus = commandBus{
			Name:      HTTPCommandBus.String(),
			AdapterID: HTTPCommandBus,
		}
	case GRPCCommandBusName:
		m.CommandBus = commandBus{
			Name:      GRPCCommandBus.String(),
			AdapterID: GRPCCommandBus,
		}
	case NATSCommandBusName:
		m.CommandBus = commandBus{
			Name:      NATSCommandBus.String(),
			AdapterID: NATSCommandBus,
		}
	}
}

type publisher struct {
	Name      string           `yaml:"name,omitempty"`
	AdapterID publisherAdapter `yaml:"id"`
}

type Mutations struct {
	Commands []CommandMutation `yaml:"from_commands"`
	Events   []EventMutation   `yaml:"from_events"`
}

func (m Mutations) HasCommands() bool {
	return len(m.Commands) > 0
}

func (m Mutations) HasEvents() bool {
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

func (s streamStorage) IsDefault() bool {
	return s.AdapterID == DefaultStreamStorage
}

func (s streamStorage) IsRedis() bool {
	return s.AdapterID == RedisStreamStorage
}

func (s streamStorage) IsPostgres() bool {
	return s.AdapterID == PostgresStreamStorage
}

type commandBus struct {
	Name      string            `yaml:"name,omitempty"`
	AdapterID commandBusAdapter `yaml:"id"`
}

func (c commandBus) IsNATS() bool {
	return c.AdapterID == NATSCommandBus
}

func (c commandBus) IsHTTP() bool {
	return c.AdapterID == HTTPCommandBus
}

func (c commandBus) IsGRPC() bool {
	return c.AdapterID == GRPCCommandBus
}

type CommandMutation struct {
	Mutation string  `yaml:"mutation"`
	Command  Command `yaml:"in_command"`
	Event    Event   `yaml:"out_event"`
	Create   string  `yaml:"allow_create_stream,omitempty"`
	Delete   string  `yaml:"allow_delete_stream,omitempty"`
}

func (c CommandMutation) ControllerName() string {
	return strutil.LcFirst(c.Mutation) + "CommandController"
}

type EventMutation struct {
	Mutation string `yaml:"mutation"`
	InEvent  Event  `yaml:"in_event"`
	OutEvent Event  `yaml:"out_event"`
	Create   string `yaml:"allow_create_stream,omitempty"`
	Delete   string `yaml:"allow_delete_stream,omitempty"`
}

func (c EventMutation) ControllerName() string {
	return strutil.LcFirst(c.Mutation) + "EventController"
}

type Command struct {
	Name    string `yaml:"name"`
	Payload string `yaml:"payload"`
}

type Event struct {
	Name    string `yaml:"name"`
	Payload string `yaml:"payload"`
}

func (e Event) LcFirstName() string {
	return strutil.LcFirst(e.Name)
}
