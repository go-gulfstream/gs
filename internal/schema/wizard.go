package schema

import (
	"fmt"

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

func (w *Wizard) setupName() error {
	prompt := promptui.Prompt{
		Label:     "ProjectName [for example: My Project] :",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.Name = res

	return nil
}

func (w *Wizard) setupPackageName() error {
	prompt := promptui.Prompt{
		Label:     "GoPackageName [for example: myproject] :",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.PackageName = sanitizePackageName(res)

	return nil
}

func (w *Wizard) setupStreamName() error {
	prompt := promptui.Prompt{
		Label:     "GoStreamName [for example: Myproject] :",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.StreamName = sanitizeStreamName(res)

	return nil
}

func (w *Wizard) setupGoModules() error {
	prompt := promptui.Prompt{
		Label:     "GoModules [for example: github.com/go-gulfstream/myproject] :",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.GoModules = sanitizeName(res)

	return nil
}

func (w *Wizard) setupGoEventsPkg() error {
	prompt := promptui.Prompt{
		Label:     "GoEventsPkg [for example: myprojectevents] :",
		Default:   w.manifest.PackageName + "events",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.EventsPkgName = sanitizePackageName(res)

	return nil
}

func (w *Wizard) setupGoCommandsPkg() error {
	prompt := promptui.Prompt{
		Label:     "GoCommandsPkg [for example: myprojectcommands] :",
		Default:   w.manifest.PackageName + "commands",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.CommandsPkgName = sanitizePackageName(res)

	return nil
}

func (w *Wizard) setupGoStreamPkg() error {
	prompt := promptui.Prompt{
		Label:     "GoStreamPkg [for example: myprojectstream] :",
		Default:   w.manifest.PackageName + "stream",
		Templates: inputTpl(),
		Validate:  validateInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.StreamPkgName = sanitizePackageName(res)

	return nil
}

func (w *Wizard) setupDescription() error {
	prompt := promptui.Prompt{
		Label:     "Description :",
		Templates: inputTpl(),
	}

	res, err := prompt.Run()
	if err != nil {
		return err
	}

	w.manifest.Description = res

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
	adapterID++

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
	adapterID++

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

func (w *Wizard) Manifest() *Manifest {
	return w.manifest
}

func (w *Wizard) Run() error {
	for _, wizardFunc := range []func() error{
		w.setupName,
		w.setupPackageName,
		w.setupStreamName,
		w.setupGoModules,
		w.setupGoCommandsPkg,
		w.setupGoStreamPkg,
		w.setupGoEventsPkg,
		w.setupDescription,
		w.setupStreamStorage,
		w.setupStreamPublisher,
	} {
		if err := wizardFunc(); err != nil {
			return err
		}
	}
	return nil
}

func inputTpl() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
}

func validateInput(s string) error {
	if len(s) < 3 {
		return fmt.Errorf("too short value")
	}
	return nil
}
