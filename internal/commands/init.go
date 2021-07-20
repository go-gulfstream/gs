package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

func initCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "init [PATH]",
		Short: "Init a new gulfstream project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateArgsInitCommand(args); err != nil {
				return err
			}
			drawBanner()
			return runInitCommand()
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
	manifest := filepath.Join(args[0], ".gulfstream.yml")
	if _, err := os.Stat(manifest); os.IsExist(err) {
		return err
	}
	return nil
}

func runInitCommand() error {

	// project name
	prompt := promptui.Prompt{
		Label: "Project name",
		Validate: func(s string) error {
			if len(s) > 3 {
				return nil
			}
			return fmt.Errorf("project name to short")
		},
	}
	projectName, err := prompt.Run()
	if err != nil {
		return err
	}

	// stream name
	prompt = promptui.Prompt{
		Label:   "Stream name",
		Default: projectName,
		Validate: func(s string) error {
			if len(s) > 3 {
				return nil
			}
			return fmt.Errorf("stream name to short")
		},
	}
	streamName, err := prompt.Run()
	if err != nil {
		return err
	}

	// go mod
	prompt = promptui.Prompt{
		Label:   "Go module (go.mod)",
		Default: projectName,
		Validate: func(s string) error {
			if len(s) > 3 {
				return nil
			}
			return fmt.Errorf("go.mod module to short")
		},
	}
	goMod, err := prompt.Run()
	if err != nil {
		return err
	}

	// author
	prompt = promptui.Prompt{
		Label: "Author",
	}
	author, err := prompt.Run()
	if err != nil {
		return err
	}

	// email
	prompt = promptui.Prompt{
		Label: "Email",
	}
	email, err := prompt.Run()
	if err != nil {
		return err
	}

	// description
	prompt = promptui.Prompt{
		Label: "Description",
	}
	desc, err := prompt.Run()
	if err != nil {
		return err
	}

	_ = projectName
	_ = streamName
	_ = goMod
	_ = author
	_ = email
	_ = desc

	return nil
}
