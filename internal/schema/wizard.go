package schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

const (
	redisStreamStorageAdapter storageAdapter = iota
	postgresStreamStorageAdapter
)

const (
	kafkaStreamPublisherAdapter publisherAdapter = iota
	connectorStreamPublisherAdapter
)

type (
	storageAdapter   int
	publisherAdapter int
)

func (a storageAdapter) IsRedis() bool {
	return redisStreamStorageAdapter == a
}

func (a storageAdapter) IsPostgreSQL() bool {
	return postgresStreamStorageAdapter == a
}

func (a storageAdapter) String() string {
	switch a {
	case postgresStreamStorageAdapter:
		return fmt.Sprintf("%d. PostgreSQL - Stream storage adapter", a)
	case redisStreamStorageAdapter:
		return fmt.Sprintf("%d. Redis - Stream storage adapter", a)
	}
	return "Unknown"
}

func (a publisherAdapter) String() string {
	switch a {
	case kafkaStreamPublisherAdapter:
		return fmt.Sprintf("%d. Kafka publisher - Stream publisher adapter", a)
	case connectorStreamPublisherAdapter:
		return fmt.Sprintf("%d. WAL Connector - Stream publisher adapter", a)
	}
	return "Unknown"
}

func (a publisherAdapter) IsKafka() bool {
	return kafkaStreamPublisherAdapter == a
}

func (a publisherAdapter) IsConnector() bool {
	return connectorStreamPublisherAdapter == a
}

func Adapters() []string {
	return []string{
		redisStreamStorageAdapter.String(),
		postgresStreamStorageAdapter.String(),
		kafkaStreamPublisherAdapter.String(),
		connectorStreamPublisherAdapter.String(),
	}
}

type Wizard struct {
	manifest *Manifest
}

func NewSetupWizard() *Wizard {
	wiz := &Wizard{
		manifest: new(Manifest),
	}
	return wiz
}

func (w *Wizard) setupContributor() error {
	prompt := promptui.Prompt{
		Label: "Author",
	}
	author, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{
		Label: "Email",
	}
	email, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{
		Label: "Description",
	}
	desc, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.Contributor.Author = author
	w.manifest.Contributor.Email = email
	w.manifest.Contributor.Description = desc

	return nil
}

func (w *Wizard) setupStreamPublisher() error {
	var adapters []string
	if w.manifest.StreamStorage.AdapterID.IsPostgreSQL() {
		adapters = []string{
			kafkaStreamPublisherAdapter.String(),
			connectorStreamPublisherAdapter.String(),
		}
	}
	if w.manifest.StreamStorage.AdapterID.IsRedis() {
		adapters = []string{
			kafkaStreamPublisherAdapter.String(),
		}
	}
	prompt := promptui.Select{
		Label: "Select stream publisher adapter",
		Items: adapters,
	}
	adapterID, adapterName, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.Publisher.AdapterName = adapterName
	w.manifest.Publisher.AdapterID = publisherAdapter(adapterID)

	return nil
}

func (w *Wizard) setupStreamStorage() error {
	prompt := promptui.Select{
		Label: "Select stream storage adapter",
		Items: []string{
			redisStreamStorageAdapter.String(),
			postgresStreamStorageAdapter.String(),
		},
	}
	adapterID, adapterName, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.StreamStorage.AdapterID = storageAdapter(adapterID)
	w.manifest.StreamStorage.AdapterName = adapterName

	if !w.manifest.StreamStorage.AdapterID.IsPostgreSQL() {
		return nil
	}

	// enable journal if needed
	promptJournal := promptui.Prompt{
		Label:     fmt.Sprintf("%s enable event journal?", adapterName),
		IsConfirm: true,
	}
	_, err = promptJournal.Run()
	if err == nil {
		w.manifest.StreamStorage.EnableJournal = true
	}

	return nil
}

func (w *Wizard) setupProjectInfo() error {
	prompt := promptui.Prompt{
		Label: "Project name",
		Validate: func(s string) error {
			if len(s) > 3 {
				return nil
			}
			return fmt.Errorf("project name to short")
		},
	}
	projectName, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{
		Label:   "Go module (go.mod)",
		Default: projectName,
		Validate: func(s string) error {
			if len(s) > 3 {
				return nil
			}
			return fmt.Errorf("go.mod module to short")
		},
	}
	goMod, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.Project.Name = strings.ToLower(projectName)
	w.manifest.Project.CreatedAt = time.Now().UTC()
	w.manifest.Project.GoModules = strings.ToLower(goMod)

	return nil
}

func (w *Wizard) Manifest() *Manifest {
	return w.manifest
}

func (w *Wizard) Run() error {
	for _, wizardFunc := range []func() error{
		w.setupProjectInfo,
		w.setupStreamStorage,
		w.setupStreamPublisher,
		w.setupContributor,
	} {
		if err := wizardFunc(); err != nil {
			return err
		}
	}
	return nil
}