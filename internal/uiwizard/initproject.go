package uiwizard

import (
	"fmt"

	"github.com/go-gulfstream/gs/internal/schema"
)

type Project struct {
	name        string
	streamName  string
	packageName string
	gomodules   string
	commandPkg  string
	eventPkg    string
	streamPkg   string
	description string
	publisher   selectItem
	storage     selectItem
	journal     bool
}

func NewProject() *Project {
	return new(Project)
}

func (p *Project) Apply(m *schema.Manifest) {
	m.Name = p.name
	m.PackageName = p.packageName
	m.StreamName = p.streamName
	m.GoModules = p.gomodules
	m.EventsPkgName = p.eventPkg
	m.CommandsPkgName = p.commandPkg
	m.StreamPkgName = p.streamPkg
	m.Description = p.description
	m.SetStreamStorageFromString(p.storage.ID, p.journal)
	m.SetPublisherFromString(p.publisher.ID)
}

func (p *Project) Confirm() (bool, error) {
	return confirmControl("Proceed?")
}

func (p *Project) Run() (err error) {
	sectionControl("Project info")
	p.name, err = inputControl("Name", "My Project", true)
	if err != nil {
		return err
	}
	p.streamName, err = inputControl("StreamName", "MyProject", true)
	if err != nil {
		return err
	}
	p.gomodules, err = inputControl("GoModules", "github.com/myproject", true)
	if err != nil {
		return err
	}
	p.packageName, err = inputControl("GoPackage", "myproject", true)
	if err != nil {
		return err
	}
	p.commandPkg, err = inputControl("GoCommandPkg", fmt.Sprintf("%scommand", p.packageName), false)
	if err != nil {
		return err
	}
	p.eventPkg, err = inputControl("GoEventPkg", fmt.Sprintf("%sevents", p.packageName), false)
	if err != nil {
		return err
	}
	p.streamPkg, err = inputControl("GoStreamPkg", fmt.Sprintf("%sstream", p.packageName), false)
	if err != nil {
		return err
	}
	p.description, err = inputControl("Description", "", false)
	if err != nil {
		return err
	}
	storage, journal, err := p.selectStreamStorage()
	if err != nil {
		return err
	}
	p.storage = storage
	p.journal = journal
	p.publisher, err = p.selectStreamPublisher()
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) selectStreamPublisher() (selectItem, error) {
	adapters := []selectItem{
		{
			ID:   schema.DefaultName,
			Name: "Default",
			Help: "In memory adapter",
		},
		{
			ID:   schema.KafkaStreamPublisherAdapterName,
			Name: schema.KafkaStreamPublisherAdapter.String(),
			Help: "Kafka adapter",
		},
	}
	if p.storage.ID == schema.PostgresStreamStorageAdapterName {
		adapters = append(adapters, selectItem{
			ID:   schema.ConnectorStreamPublisherAdapterName,
			Name: schema.ConnectorStreamPublisherAdapter.String(),
			Help: "CDC adapter from postgres to kafka",
		})
	}
	return selectControl("Select stream publisher adapter", adapters)
}

func (p *Project) selectStreamStorage() (selectItem, bool, error) {
	item, err := selectControl("Select stream storage adapter", []selectItem{
		{
			ID:   schema.DefaultName,
			Name: "Default",
			Help: "In memory adapter",
		},
		{
			ID:   schema.RedisStreamStorageAdapterName,
			Name: schema.RedisStreamStorageAdapter.String(),
			Help: "Redis adapter",
		},
		{
			ID:   schema.PostgresStreamStorageAdapterName,
			Name: schema.PostgresStreamStorageAdapter.String(),
			Help: "CDC adapter from postgres to kafka",
		},
	})
	if err != nil {
		return selectItem{}, false, err
	}

	if item.ID == schema.PostgresStreamStorageAdapterName {
		ok, err := confirmControl(fmt.Sprintf("%s enable event journal?", item.Name))
		if err != nil {
			return selectItem{}, false, err
		}
		return item, ok, nil
	}

	return item, false, nil
}
