package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/source"

	storeagepostgres "github.com/go-gulfstream/gulfstream/pkg/storage/postgres"

	"github.com/fatih/color"
	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

const manifestFilename = "gulfstream.yml"

var (
	yellowColor = color.New(color.FgYellow).SprintFunc()
	redColor    = color.New(color.FgRed).SprintFunc()
	greenColor  = color.New(color.FgGreen).SprintfFunc()
	boldStyle   = color.New(color.Bold).SprintFunc()
)

func initCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "init [PATH]",
		Short: "Create a new Gulfstream project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateInitCommandArgs(args); err != nil {
				return err
			}
			drawBanner()
			return runInitCommand(args[0])
		},
	}
	return command
}

func validateInitCommandArgs(args []string) error {
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

func runInitCommand(projectPath string) (err error) {
	manifest, err := loadManifestFromFile(projectPath)
	if err != nil {
		return err
	}

	if err := schema.Walk(projectPath, manifest,
		func(file schema.File) (err error) {
			if file.IsDir {
				// make dir
				err = os.Mkdir(file.Path, 0755)
			} else {
				// fix imports and code format.
				if file.IsGo() {
					src, err := source.Format(file.Path, file.TemplateData)
					if err != nil {
						return err
					}
					file.TemplateData = src
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

	if err := writeManifestToFile(projectPath, manifest, true); err != nil {
		return fmt.Errorf("writeManifestToFile %v", err)
	}

	if err := writeSchema(projectPath, manifest); err != nil {
		return fmt.Errorf("writeSchema %s", err)
	}

	runGoTools(projectPath, manifest.GoGetPackages)

	return nil
}

func writeSchema(path string, m *schema.Manifest) error {
	if m.StreamStorage.AdapterID.IsPostgreSQL() {
		filename := filepath.Join(path, "/deployments/gulfstream-schema.sql")
		return ioutil.WriteFile(filename, []byte(storeagepostgres.Schema), 0755)
	}
	return nil
}
