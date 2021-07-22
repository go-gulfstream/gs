package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/format"

	"github.com/fatih/color"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

const manifestFilename = "gulfstream.yml"

func initCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "init [PATH]",
		Short: "Create a new Gulfstream project",
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
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		return err
	}
	errAlreadyExists := fmt.Errorf("project already exists")
	manifest := filepath.Join(args[0], manifestFilename)
	fi, err := os.Stat(manifest)
	if err != nil {
		if os.IsExist(err) && len(files) > 1 {
			return errAlreadyExists
		} else if os.IsNotExist(err) && len(files) > 0 {
			return fmt.Errorf("directory not empty %s", args[0])
		}
	} else if fi.Size() > 0 && len(files) > 1 {
		return errAlreadyExists
	}
	return nil
}

func runInitCommand(path string) error {
	var (
		manifest *schema.Manifest
		nof      bool
	)

	// from manifest file
	manifestFile := filepath.Join(path, manifestFilename)
	if _, err := os.Stat(manifestFile); err == nil {
		data, err := ioutil.ReadFile(manifestFile)
		if err != nil {
			return err
		}
		manifest = new(schema.Manifest)
		if err := manifest.UnmarshalBinary(data); err != nil {
			return err
		}
		if err := manifest.Validate(); err != nil {
			return err
		}
		manifest.Sanitize()
	} else {
		// from setup wizard
		wizard := schema.NewSetupWizard()
		if err := wizard.Run(); err != nil {
			return err
		}
		manifest = wizard.Manifest()
		nof = true
	}

	yellowColor := color.New(color.FgYellow).SprintFunc()
	redColor := color.New(color.FgRed).SprintFunc()
	greenColor := color.New(color.FgGreen).SprintfFunc()

	if err := schema.Walk(path, manifest,
		func(file schema.File) (err error) {
			if file.IsDir {
				// make dir
				err = os.Mkdir(file.Path, 0755)
			} else {
				// fix imports and code format.
				if file.IsGo() {
					source, err := format.Source(file.TemplateData)
					if err != nil {
						return err
					}
					file.TemplateData = source
				}
				// make file
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

	if nof {
		data, err := manifest.MarshalBinary()
		if err != nil {
			return err
		}
		manifest := filepath.Join(path, manifestFilename)
		return ioutil.WriteFile(manifest, data, 0777)
	}

	return nil
}
