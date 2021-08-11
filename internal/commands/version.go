package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCommand(version string) *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Show the gs version information",
		Run: func(cmd *cobra.Command, args []string) {
			drawBanner()
			fmt.Printf("%s %s\n", greenColor("[INFO]"), version)
		},
	}
	return command
}
