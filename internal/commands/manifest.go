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
	manifest.Name = "User sessions store"
	manifest.PackageName = "session"
	manifest.StreamName = "Session"
	manifest.GoModules = "github.com/go-gulfstream/session"
	manifest.Description = ""
	manifest.Contributors = []schema.Contributor{
		{
			Author: "author",
			Email:  "author@gmail.com",
		},
	}
	manifest.EventsPkgName = "sessionevents"
	manifest.CommandsPkgName = "sessioncommands"
	manifest.StreamPkgName = "sessionstream"
	manifest.Mutations.Commands = []schema.CommandMutation{
		{
			Mutation: "CreateSessionMutation",
			Command: schema.Command{
				Name:    "CreateSession",
				Payload: "CreateSessionPayload",
			},
			Event: schema.Event{
				Name:    "SessionCreated",
				Payload: "SessionCreatedPayload",
			},
			Create: schema.YesOp,
		},
		{
			Mutation: "UpdateSessionMutation",
			Command: schema.Command{
				Name:    "UpdateSession",
				Payload: "UpdateSessionPayload",
			},
			Event: schema.Event{
				Name:    "SessionUpdated",
				Payload: "SessionUpdatedPayload",
			},
		},
		{
			Mutation: "RemoveSessionMutation",
			Command: schema.Command{
				Name:    "RemoveSession",
				Payload: "RemoveSessionPayload",
			},
			Event: schema.Event{
				Name:    "SessionRemoved",
				Payload: "SessionRemovedPayload",
			},
			Delete: schema.YesOp,
		},
	}
	manifest.StreamStorage.Name = schema.RedisStreamStorageAdapter.String()
	manifest.StreamStorage.AdapterID = schema.RedisStreamStorageAdapter
	manifest.StreamPublisher.Name = schema.KafkaStreamPublisherAdapter.String()
	manifest.StreamPublisher.AdapterID = schema.KafkaStreamPublisherAdapter
	manifest.ImportEvents = []string{}
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
