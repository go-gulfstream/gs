package schema

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

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
	//prompt := promptui.Prompt{
	//	Label: "Author",
	//}
	//author, err := prompt.Run()
	//if err != nil {
	//	return err
	//}
	//
	//prompt = promptui.Prompt{
	//	Label: "Email",
	//}
	//email, err := prompt.Run()
	//if err != nil {
	//	return err
	//}
	//
	//prompt = promptui.Prompt{
	//	Label: "Description",
	//}
	//desc, err := prompt.Run()
	//if err != nil {
	//	return err
	//}

	//w.manifest.Contributor.Author = author
	//w.manifest.Contributor.Email = email
	//w.manifest.Contributor.Description = desc

	return nil
}

func (w *Wizard) setupStreamPublisher() error {
	var adapters []string
	if w.manifest.StreamStorage.AdapterID.IsPostgreSQL() {
		adapters = []string{
			KafkaStreamPublisherAdapter.String(),
			ConnectorStreamPublisherAdapter.String(),
		}
	}
	if w.manifest.StreamStorage.AdapterID.IsRedis() {
		adapters = []string{
			KafkaStreamPublisherAdapter.String(),
		}
	}
	prompt := promptui.Select{
		Label: "Select stream publisher adapter",
		Items: adapters,
	}
	adapterID, _, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.StreamPublisher.AdapterID = publisherAdapter(adapterID)

	return nil
}

func (w *Wizard) setupStreamStorage() error {
	prompt := promptui.Select{
		Label: "Select stream storage adapter",
		Items: []string{
			RedisStreamStorageAdapter.String(),
			PostgresStreamStorageAdapter.String(),
		},
	}
	adapterID, adapterName, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.StreamStorage.AdapterID = storageAdapter(adapterID)

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

	w.manifest.Name = strings.ToLower(projectName)
	//w.manifest.Project.CreatedAt = time.Now().UTC()
	w.manifest.GoModules = strings.ToLower(goMod)

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
