package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/schema"

	"github.com/spf13/cobra"
)

func manifestCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "manifest [PATH]",
		Short: "Create empty manifest file for new project",
		RunE: func(cmd *cobra.Command, args []string) error {
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
			if len(files) > 0 {
				return fmt.Errorf("directory not empty %s", args[0])
			}
			data, err := schema.MarshalBlankManifest()
			if err != nil {
				return err
			}
			manifestFile := filepath.Join(args[0], manifestFilename)
			if _, err := os.Stat(manifestFile); err == nil {
				return fmt.Errorf("manifest file already exists")
			}
			return ioutil.WriteFile(manifestFile, data, 0755)
		},
	}
	return command
}
