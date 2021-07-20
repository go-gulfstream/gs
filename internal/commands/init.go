package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"

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
	fi, err := os.Stat(manifest)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("project already exists")
		}
		return nil
	}
	if fi.Size() > 0 {
		return fmt.Errorf("project already exists")
	}
	return nil
}

func runInitCommand(path string) error {
	wizard := schema.NewSetupWizard()
	if err := wizard.Run(); err != nil {
		return err
	}

	yellowColor := color.New(color.FgYellow).SprintFunc()
	redColor := color.New(color.FgRed).SprintFunc()
	greenColor := color.New(color.FgGreen).SprintfFunc()

	if err := schema.Walk(path, wizard.Manifest(),
		func(file schema.File) (err error) {
			if file.IsDir {
				err = os.Mkdir(file.Path, 0755)
			} else {
				err = ioutil.WriteFile(file.Path, file.TemplateData, 0755)
			}
			if os.IsExist(err) {
				fmt.Printf("%s - %s\n", yellowColor("[SKIP]"), file.Path)
				return nil
			}

			if err == nil {
				fmt.Printf("%s - %s\n", greenColor("[OK]"), file.Path)
			} else {
				fmt.Printf("%s - %s\n", redColor("[ERR]"), file.Path)
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
