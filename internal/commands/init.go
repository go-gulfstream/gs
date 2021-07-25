package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/goutil"

	storeagepostgres "github.com/go-gulfstream/gulfstream/pkg/storage/postgres"

	"github.com/fatih/color"
	"github.com/go-gulfstream/gs/internal/format"
	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

const manifestFilename = "gulfstream.yml"

var (
	yellowColor = color.New(color.FgYellow).SprintFunc()
	redColor    = color.New(color.FgRed).SprintFunc()
	greenColor  = color.New(color.FgGreen).SprintfFunc()
)

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

func runInitCommand(path string) (err error) {
	var manifest *schema.Manifest

	// from manifest file
	manifestFile := filepath.Join(path, manifestFilename)
	if _, err := os.Stat(manifestFile); err == nil {
		data, err := ioutil.ReadFile(manifestFile)
		if err != nil {
			return err
		}
		manifest, err = schema.DecodeManifest(data)
		if err != nil {
			return err
		}
		schema.SanitizeManifest(manifest)
		if err := schema.ValidateManifest(manifest); err != nil {
			return err
		}
	} else {
		// from setup wizard
		wizard := schema.NewSetupWizard()
		if err := wizard.Run(); err != nil {
			return err
		}
		manifest = wizard.Manifest()
	}

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

	if err := writeManifestFile(path, manifest, true); err != nil {
		return err
	}

	if err := writeSchema(path, manifest); err != nil {
		return err
	}

	runGoTools(path)

	return nil
}

func writeSchema(path string, m *schema.Manifest) error {
	if m.StreamStorage.AdapterID.IsPostgreSQL() {
		filename := filepath.Join(path, "gulfstream-schema.sql")
		return ioutil.WriteFile(filename, []byte(storeagepostgres.Schema), 0755)
	}
	return nil
}

func runGoTools(path string) {
	if !goutil.GoInstall() {
		return
	}

	fmt.Printf("======== go tools ========\n")
	fmt.Printf("go mod download:\n")
	out, err := goutil.RunGoMod(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))

	fmt.Printf("go mod tidy:\n")
	out, err = goutil.RunGoModTidy(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))

	fmt.Printf("go test ./...:\n")
	out, err = goutil.RunGoTest(path)
	if err != nil {
		fmt.Printf("%s - %s\n", redColor("[ERR]"), err)
		return
	}
	fmt.Printf("%s - %s\n", greenColor("[OK]"), string(out))
}
