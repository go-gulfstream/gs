package commands

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-gulfstream/gs/internal/goutil"

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
			dataFlag := cmd.Flag("data")
			withData := dataFlag.Value.String() == "true"
			manifest := blankManifest(withData)
			return writeManifestFile(args[0], manifest, false)
		},
	}
	command.Flags().BoolP("data", "d", false, "generate with data example")
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

func blankManifest(withData bool) *schema.Manifest {
	goVersion := goutil.Version()

	if !withData {
		manifest := new(schema.Manifest)
		manifest.GoVersion = goVersion
		manifest.Contributors = []schema.Contributor{{}}
		manifest.Mutations.Commands = []schema.CommandMutation{{}}
		manifest.Mutations.Events = []schema.EventMutation{{}}
		return manifest
	}

	manifest := new(schema.Manifest)
	manifest.GoVersion = goVersion
	manifest.Name = "My project"
	manifest.PackageName = "myproject"
	manifest.StreamName = "Myproject"
	manifest.GoModules = "github.com/go-gulfstream/myproject"
	manifest.Description = "My project short description"
	manifest.Contributors = []schema.Contributor{
		{
			Author: "author",
			Email:  "author@gmail.com",
		},
	}
	manifest.EventsPkgName = "myprojectevents"
	manifest.CommandsPkgName = "myprojectcommands"
	manifest.StreamPkgName = "myprojectstream"
	manifest.Mutations.Commands = []schema.CommandMutation{
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
			Create: schema.YesOp,
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
			Delete: schema.YesOp,
		},
	}
	manifest.StreamStorage.Name = schema.RedisStreamStorageAdapter.String()
	manifest.StreamStorage.AdapterID = schema.RedisStreamStorageAdapter
	manifest.StreamPublisher.Name = schema.KafkaStreamPublisherAdapter.String()
	manifest.StreamPublisher.AdapterID = schema.KafkaStreamPublisherAdapter
	manifest.Mutations.Events = []schema.EventMutation{{
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
	manifest.ImportEvents = []string{
		"github.com/go-gulfstream/counterevents",
	}
	return manifest
}

func writeManifestFile(path string, manifest *schema.Manifest, force bool) error {
	data, err := schema.EncodeManifest(manifest)
	if err != nil {
		return err
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
	manifestFile := filepath.Join(path, manifestFilename)
	if _, err := os.Stat(manifestFile); err == nil && !force {
		return fmt.Errorf("manifest file already exists")
	}
	return ioutil.WriteFile(manifestFile, buf.Bytes(), 0755)
}
