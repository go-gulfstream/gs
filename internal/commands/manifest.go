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

type manifestFlags struct {
	numCommands int
	numEvents   int
	withData    bool
}

func manifestCommand() *cobra.Command {
	var flags manifestFlags
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
			drawBanner()
			manifest := blankManifest(&flags)
			if err := writeManifestFile(args[0], manifest, false); err != nil {
				return err
			}
			fmt.Printf("New manifest file created successfully %s/%s\n", args[0], manifestFilename)
			return nil
		},
	}

	command.Flags().BoolVarP(&flags.withData, "data", "d", false, "with data example")
	command.Flags().IntVarP(&flags.numCommands, "commands", "c", 1, "number of template commands")
	command.Flags().IntVarP(&flags.numEvents, "events", "e", 1, "number of template events")

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

func blankManifest(f *manifestFlags) *schema.Manifest {
	goVersion := goutil.Version()

	manifest := new(schema.Manifest)
	manifest.GoVersion = goVersion
	tpl := func(n int, name string) string {
		return fmt.Sprintf("$template_%d_%s", n, name)
	}

	for i := 0; i < f.numEvents; i++ {
		manifest.Mutations.Events = append(manifest.Mutations.Events,
			schema.EventMutation{
				Mutation: tpl(i, "EventMutation"),
				InEvent: schema.Event{
					Name:    tpl(i, "InEvent"),
					Payload: tpl(i, "InEventPayload"),
				},
				OutEvent: schema.Event{
					Name:    tpl(i, "OutEvent"),
					Payload: tpl(i, "OutEventPayload"),
				},
			})
	}

	for i := 0; i < f.numCommands; i++ {
		manifest.Mutations.Commands = append(manifest.Mutations.Commands,
			schema.CommandMutation{
				Mutation: tpl(i, "CommandMutation"),
				Command: schema.Command{
					Name:    tpl(i, "Command"),
					Payload: tpl(i, "CommandPayload"),
				},
				Event: schema.Event{
					Name:    tpl(i, "Event"),
					Payload: tpl(i, "EventPayload"),
				},
			})
	}

	if !f.withData {
		manifest.Contributors = []schema.Contributor{{}}
		return manifest
	}

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
