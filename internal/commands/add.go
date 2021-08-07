package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/go-gulfstream/gs/internal/uiwizard"
	"github.com/spf13/cobra"
)

type addFlags struct {
	Apply bool
}

func addCommand() *cobra.Command {
	var flags addFlags
	command := &cobra.Command{
		Use:   "add [PATH]",
		Short: "Manage adding",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateAddCommandArgs(args); err != nil {
				return err
			}
			drawBanner()
			return runAddCommand(args[0], flags)
		},
	}
	command.Flags().BoolVarP(&flags.Apply, "apply", "a", false, "add and apply changes to the project")
	return command
}

func validateAddCommandArgs(args []string) error {
	lenArgs := len(args)
	if lenArgs != 1 {
		return fmt.Errorf("invalid number of arguments. got %d, expected 1\n\nfor example:\n$ gs apply ~/myproject\n", lenArgs)
	}
	manifest := filepath.Join(args[0], manifestFilename)
	_, err := os.Stat(manifest)
	if os.IsNotExist(err) {
		return fmt.Errorf("the manifest file %s/gulfstream.yml does not exist\n$ gs manifest -i %s\n", args[0], args[0])
	}
	return nil
}

func runAddCommand(projectPath string, f addFlags) error {
	manifest, err := loadManifestFromFile(projectPath)
	if err != nil {
		return err
	}

	schema.Index(manifest)

	var projectInit bool
	if countProjectFiles(projectPath) <= 1 {
		projectInit = true
		fmt.Printf("\n%s\n\n", yellowColor("Attention! project not created"))
	}

	wiz := uiwizard.NewMutation()
	if err := wiz.Run(); err != nil {
		return err
	}
	if !wiz.HasChanges() {
		fmt.Printf("no data to add\n")
		return nil
	}

	wiz.Apply(manifest)

	if err := writeManifestToFile(projectPath, manifest, true); err != nil {
		return err
	}
	if f.Apply {
		if projectInit {
			err = runInitCommand(projectPath)
		} else {
			err = runApplyCommand(projectPath)
		}
	} else {
		fmt.Printf("%s!\n", yellowColor("Adding without applying"))
	}
	return err
}
