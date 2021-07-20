package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

const manifestFilename = ".gulfstream.yml"

func initCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "init [PATH]",
		Short: "Init a new gulfstream project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateArgsInitCommand(args); err != nil {
				return err
			}
			drawBanner()
			return runInitCommand(args[0])
		},
	}
	return command
}

func validateArgsInitCommand(args []string) error {
	lenArgs := len(args)
	if lenArgs != 1 {
		return fmt.Errorf("invalid number of arguments. got %d, expected 1", lenArgs)
	}
	if _, err := os.Stat(args[0]); err != nil {
		return err
	}
	manifest := filepath.Join(args[0], manifestFilename)
	if _, err := os.Stat(manifest); os.IsExist(err) {
		return err
	}
	return nil
}

func runInitCommand(path string) error {
	wizard := schema.NewSetupWizard()
	if err := wizard.Run(); err != nil {
		return err
	}

	if err := schema.Walk(path, wizard.Manifest(),
		func(file schema.File) (err error) {
			if file.IsDir {
				err = os.Mkdir(file.Path, 0755)
			} else {
				err = ioutil.WriteFile(file.Path, file.TemplateData, 0777)
			}
			if os.IsExist(err) {
				return nil
			}
			return
		}); err != nil {
		return err
	}

	data, err := wizard.Manifest().MarshalBinary()
	if err != nil {
		return err
	}
	manifest := filepath.Join(path, manifestFilename)
	return ioutil.WriteFile(manifest, data, 0777)
}
