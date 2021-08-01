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
		Short: "Update a Gulfstream project from manifest file",
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
	manifest := filepath.Join(args[0], manifestFilename)
	_, err = os.Stat(manifest)
	if os.IsNotExist(err) {
		return fmt.Errorf("the manifest file %s/gulfstream.yml does not exist", args[0])
	}
	return nil
}

func runApplyCommand(path string) error {
	manifestFile := filepath.Join(path, manifestFilename)
	data, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return err
	}
	manifest, err := schema.DecodeManifest(data)
	if err != nil {
		return err
	}

	schema.SanitizeManifest(manifest)
	if err := schema.ValidateManifest(manifest); err != nil {
		return err
	}

	if err := schema.Validate(path, manifest); err != nil {
		return err
	}
	if err := source.Validate(path, manifest); err != nil {
		return err
	}

	info, err := source.Stats(path, manifest)
	if err != nil {
		return err
	}

	statusOk := greenColor("[ADD]")
	statusSkip := yellowColor("[EXISTS]")

	if err := schema.WalkCommandMutationAddons(path, manifest,
		func(m schema.CommandMutation, file schema.File) error {
			status := statusOk
			var skip bool
			if info.CommandMutationExists(m.Mutation) {
				status = statusSkip
				skip = true
			}
			fmt.Printf("%s - CommmandMutation.%s{InCommand: %s, OutEvent: %s}, %s => %s\n", status,
				m.Mutation, m.Command.Name, m.Event.Name, file.Addon, file.Path)
			if skip {
				return nil
			}
			dst, err := source.FromFile(file.Path)
			if err != nil {
				return err
			}
			return source.Modify(dst, file.Addon, file.TemplateData)
		}); err != nil {
		return err
	}

	if err := schema.WalkEventMutationAddons(path, manifest,
		func(m schema.EventMutation, file schema.File) error {
			status := statusOk
			var skip bool
			if info.EventMutationExists(m.Mutation) {
				status = statusSkip
				skip = true
			}
			fmt.Printf("%s - EventMutation.%s{InEvent: %s, OutEvent: %s} => %s\n", status,
				m.Mutation, m.InEvent.Name, m.OutEvent.Name, file.Path)
			if skip {
				return nil
			}
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

	return nil
}
