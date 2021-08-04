package commands

import (
	"github.com/spf13/cobra"
)

func New() (*cobra.Command, error) {
	root := &cobra.Command{
		Use:   "gs",
		Short: "Standard Tooling for Go-Gulfstream Development",
	}

	root.AddCommand(initCommand())
	root.AddCommand(manifestCommand())
	root.AddCommand(applyCommand())
	root.AddCommand(addCommand())

	return root, nil
}
