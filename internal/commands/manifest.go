package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type manifestFlags struct {
	interactive  bool
	showManifest bool
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

	command.Flags().BoolVarP(&flags.showManifest, "show", "s", false, "show content of manifest file after creation")
	command.Flags().BoolVarP(&flags.interactive, "interactive", "i", false, "with enable editor")
	return command
}

func runManifestCommand(projectPath string, f manifestFlags) error {
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
