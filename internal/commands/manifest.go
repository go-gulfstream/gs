package commands

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
			files = filterDotFiles(files)
			if len(files) > 0 {
				return fmt.Errorf("directory %s not empty. found files: %d",
					args[0], len(files))
			}
			data, err := blankManifest()
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

func filterDotFiles(files []fs.FileInfo) []fs.FileInfo {
	filtered := make([]fs.FileInfo, 0)
	for _, fi := range files {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		filtered = append(filtered, fi)
	}
	return filtered
}

func blankManifest() ([]byte, error) {
	emptyManifest := new(schema.Manifest)
	emptyManifest.Name = "My project name"
	emptyManifest.PackageName = "myproject"
	emptyManifest.StreamName = "Myproject"
	emptyManifest.GoModules = "github.com/go-gulfstream/myproject"
	emptyManifest.Description = "My project short description"
	emptyManifest.Contributors = []schema.Contributor{
		{
			Author: "author",
			Email:  "author@gmail.com",
		},
	}
	emptyManifest.EventsPkgName = "myprojectevents"
	emptyManifest.CommandsPkgName = "myprojectcommands"
	emptyManifest.StreamPkgName = "myprojectstream"
	emptyManifest.Mutations.Commands = []schema.CommandMutation{
		{
			Mutation: "CreateTreeMutation",
			Command: schema.Command{
				Name:    "CreateTree",
				Payload: "CreateTreePayload",
			},
			Event: schema.Event{
				Name:    "TreeCreated",
				Payload: "TreeCreatedPayload",
			},
			Operations: schema.Operations{
				Create: true,
			},
		},
		{
			Mutation: "UpdateTreeMutation",
			Command: schema.Command{
				Name:    "UpdateTree",
				Payload: "UpdateTreePayload",
			},
			Event: schema.Event{
				Name:    "TreeUpdated",
				Payload: "TreeUpdatedPayload",
			},
		},
		{
			Mutation: "RemoveTreeMutation",
			Command: schema.Command{
				Name:    "RemoveTree",
				Payload: "RemoveTreePayload",
			},
			Event: schema.Event{
				Name:    "TreeRemoved",
				Payload: "TreeRemovedPayload",
			},
			Operations: schema.Operations{
				Delete: true,
			},
		},
	}
	emptyManifest.Mutations.Events = []schema.EventMutation{{
		Mutation: "UpdateCounterMutation",
		InEvent: schema.Event{
			Name:    "counterevents.CounterUpdated",
			Payload: "counterevents.CounterUpdatedPayload",
		},
		OutEvent: schema.Event{
			Name:    "TreeCounterUpdated",
			Payload: "TreeCounterUpdatedPayload",
		},
	}}
	emptyManifest.ImportEvents = []string{
		"github.com/go-gulfstream/counterevents",
	}
	data, err := schema.EncodeManifest(emptyManifest)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	buf.WriteString("\n# available storage adapters:\n")
	for id, adapter := range schema.StorageAdapters {
		buf.WriteString(fmt.Sprintf("# id:%d, name: %s\n", id, adapter))
	}
	buf.WriteString("\n# available publisher adapters:\n")
	for id, adapter := range schema.PublisherAdapters {
		buf.WriteString(fmt.Sprintf("# id:%d, name: %s\n", id, adapter))
	}
	return buf.Bytes(), nil
}
