package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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
			emptyManifest := new(schema.Manifest)
			emptyManifest.Project.Name = "someproject"
			emptyManifest.Project.CreatedAt = time.Now()
			emptyManifest.Project.GoModules = "github.com/go-gulfstream/myproject"
			emptyManifest.Mutations.Commands = []schema.CommandMutation{{
				Name:    "AssignPaymentMutation",
				Command: "AssignPayment",
				Event:   "PaymentAssigned",
			}}
			emptyManifest.Mutations.Events = []schema.EventMutation{{
				Name:  "CheckBalance",
				Event: "PaymentCreated",
			}}
			data, err := emptyManifest.MarshalBinary()
			if err != nil {
				return err
			}
			buf := bytes.NewBuffer(data)
			buf.WriteString("\n# available adapters:\n")
			for _, adapter := range schema.Adapters() {
				buf.WriteString("# - id:" + adapter + "\n")
			}
			manifestFile := filepath.Join(args[0], manifestFilename)
			if _, err := os.Stat(manifestFile); err == nil {
				return fmt.Errorf("manifest file already exists")
			}
			return ioutil.WriteFile(manifestFile, buf.Bytes(), 0755)
		},
	}
	return command
}
