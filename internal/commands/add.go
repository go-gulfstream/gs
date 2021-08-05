package commands

import (
	"fmt"

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
		Short: "Mutation manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateApplyCommandArgs(args); err != nil {
				return err
			}
			drawBanner()
			return runAddCommand(args[0], flags)
		},
	}
	command.Flags().BoolVarP(&flags.Apply, "apply", "a", false, "add and apply changes to the project")
	return command
}

func runAddCommand(projectPath string, f addFlags) error {
	wiz := uiwizard.NewMutation()
	if err := wiz.Run(); err != nil {
		return err
	}
	if !wiz.HasChanges() {
		fmt.Printf("no data to add\n")
		return nil
	}
	manifest, err := loadManifestFromFile(projectPath)
	if err != nil {
		return err
	}
	wiz.Apply(manifest)

	if err := writeManifestFile(projectPath, manifest, true); err != nil {
		return err
	}
	if f.Apply {
		err = runApplyCommand(projectPath)
	} else {
		fmt.Printf("%s!\n", yellowColor("Adding without applying"))
	}
	return err
}
