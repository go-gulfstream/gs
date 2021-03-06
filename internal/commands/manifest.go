package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-gulfstream/gs/internal/goutil"

	"github.com/go-gulfstream/gs/internal/uiwizard"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

type manifestFlags struct {
	interactive  bool
	showManifest bool
	testData     bool
}

func manifestCommand() *cobra.Command {
	var flags manifestFlags
	command := &cobra.Command{
		Use:   "manifest [PATH]",
		Short: "Create manifest file for new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateManifestArgs(args); err != nil {
				return err
			}
			drawBanner()
			return runManifestCommand(args[0], flags)
		},
	}
	command.Flags().BoolVarP(&flags.showManifest, "print", "p", false, "show content of manifest file before creation")
	command.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "with enable editor")
	command.Flags().BoolVarP(&flags.testData, "testdata", "d", false, "test data")
	return command
}

func runManifestCommand(projectPath string, f manifestFlags) error {
	manifest := schema.New()
	manifest.GoVersion = goutil.Version()
	manifest.CreatedAt = time.Now().UTC()
	manifest.UpdatedAt = time.Now().UTC()

	schema.SetDefaultGoModules(manifest)

	var isInteractiveMode bool
	if f.interactive && !f.testData {
		isInteractiveMode = true
		projwiz := uiwizard.NewProject()
		if err := projwiz.Run(); err != nil {
			return err
		}
		projwiz.Apply(manifest)

		addwiz := uiwizard.NewMutation()
		if err := addwiz.Run(); err != nil {
			return err
		}
		addwiz.Apply(manifest)
		if f.showManifest {
			printManifest(manifest)
		}
		next, err := projwiz.Confirm()
		if err != nil {
			return err
		}
		if !next {
			return nil
		}
	}

	if f.interactive && !f.testData {
		schema.SanitizeManifest(manifest)
		if err := schema.ValidateManifest(manifest); err != nil {
			return err
		}
	}

	if f.testData {
		genTestData(manifest)
	}

	if f.showManifest && !isInteractiveMode {
		printManifest(manifest)
	}

	if err := writeManifestToFile(projectPath, manifest, false); err != nil {
		return err
	}

	fmt.Println(greenColor("Success!"))
	return nil
}

func validateManifestArgs(args []string) error {
	lenArgs := len(args)
	if lenArgs != 1 {
		return fmt.Errorf("invalid number of arguments. got %d, expected 1", lenArgs)
	}
	if _, err := os.Stat(args[0]); err != nil {
		return err
	}
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		return err
	}
	files = filterDotFiles(files)
	if len(files) > 0 {
		return fmt.Errorf("directory %s not empty. found files: %d",
			args[0], len(files))
	}
	return nil
}

func genTestData(m *schema.Manifest) {
	m.Name = "My project"
	m.StreamName = "Session"
	m.PackageName = "session"
	m.GoModules = "github.com/myproject/session"
	m.EventsPkgName = "sessionevents"
	m.CommandsPkgName = "sessioncommands"
	m.StreamPkgName = "sessionstream"
	m.Description = "Some description"

	schema.SetGoModules(m, []string{
		"github.com/go-gulfstream/tmpevents/pkg/tmpevents",
	})

	m.ImportEvents = []string{
		"github.com/go-gulfstream/tmpevents/pkg/tmpevents",
	}
	m.Mutations.Commands = append(m.Mutations.Commands, schema.CommandMutation{
		Mutation: "UpdateCounter",
		Command: schema.Command{
			Name:    "CounterInfo",
			Payload: "CounterInfoPayload",
		},
		Event: schema.Event{
			Name:    "CounterUpdated",
			Payload: "CounterUpdatedPayload",
		},
	})
	m.Mutations.Events = append(m.Mutations.Events, schema.EventMutation{
		Mutation: "RegisterSession",
		InEvent: schema.Event{
			Name:    "tmpevents.SessionRegistered",
			Payload: "tmpevents.SessionRegisteredPayload",
		},
		OutEvent: schema.Event{
			Name:    "Confirmed",
			Payload: "ConfirmedPayload",
		},
	})
}
