package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/source"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

func applyCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "apply [PATH]",
		Short: "Apply manifest file to project",
		Long:  "Apply a manifest to the current project. The manifest file must be created into project directory \n$ gs manifest [PATH]\n$ gs init [PATH]\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateApplyCommandArgs(args); err != nil {
				return err
			}
			drawBanner()
			return runApplyCommand(args[0])
		},
	}
	return command
}

func validateApplyCommandArgs(args []string) error {
	lenArgs := len(args)
	if lenArgs != 1 {
		return fmt.Errorf("invalid number of arguments. got %d, expected 1\n\nfor example:\n$ gs apply ~/myproject\n", lenArgs)
	}
	if _, err := os.Stat(args[0]); err != nil {
		return err
	}
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		return err
	}
	files = filterDotFiles(files)
	if len(files) == 1 && files[0].Name() == manifestFilename {
		initCmd := boldStyle("$ gs init ", args[0])
		return fmt.Errorf("first, execute init command: [ %s ] and after run apply", initCmd)
	}
	manifest := filepath.Join(args[0], manifestFilename)
	_, err = os.Stat(manifest)
	if os.IsNotExist(err) {
		return fmt.Errorf("the manifest file %s/gulfstream.yml does not exist", args[0])
	}
	return nil
}

func runApplyCommand(projectPath string) error {
	manifest, err := loadManifestFromFile(projectPath)
	if err != nil {
		return err
	}

	if err := schema.Validate(projectPath, manifest); err != nil {
		return err
	}
	if err := source.Validate(projectPath, manifest); err != nil {
		return err
	}

	info, err := source.Stats(projectPath, manifest)
	if err != nil {
		return err
	}

	statusOk := greenColor("[ADD]")
	statusSkip := yellowColor("[SKIP]")

	var successCounter, skipCounter int

	fmt.Printf("=> Apply command mutations...\n")
	if err := schema.WalkCommandMutationAddons(projectPath, manifest,
		func(m schema.CommandMutation, file schema.File) error {
			status := statusOk
			var skip bool
			if info.CommandMutationExists(m.Mutation) {
				status = statusSkip
				skip = true
			}
			fmt.Printf("%s %s, %s => %s\n", status,
				m.Mutation, file.Addon, file.Path)
			if skip {
				skipCounter++
				return nil
			}
			successCounter++
			dst, err := source.FromFile(file.Path)
			if err != nil {
				return err
			}
			return source.Modify(dst, file.Addon, file.TemplateData)
		}); err != nil {
		return err
	}

	fmt.Printf("=> Apply event mutations...\n")
	if err := schema.WalkEventMutationAddons(projectPath, manifest,
		func(m schema.EventMutation, file schema.File) error {
			status := statusOk
			var skip bool
			if info.EventMutationExists(m.Mutation) {
				status = statusSkip
				skip = true
			}
			fmt.Printf("%s %s, %s => %s\n", status,
				m.Mutation, file.Addon, file.Path)
			if skip {
				skipCounter++
				return nil
			}
			successCounter++
			dst, err := source.FromFile(file.Path)
			if err != nil {
				return err
			}
			return source.Modify(dst, file.Addon, file.TemplateData)
		}); err != nil {
		return err
	}

	if err := source.FlushToDisk(); err != nil {
		return err
	}

	fmt.Println("===============================================")
	if successCounter > 0 {
		fmt.Printf("Added: %d\n", successCounter)
	}
	if skipCounter > 0 {
		fmt.Printf("Skipped: %d\n", skipCounter)
	}

	runGoTools(projectPath)

	return nil
}
