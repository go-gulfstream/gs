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

	return root, nil
}
