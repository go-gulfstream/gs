package commands

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/schema"
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

func loadManifestFromFile(projectPath string) (*schema.Manifest, error) {
	manifestFile := filepath.Join(projectPath, manifestFilename)
	data, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return nil, err
	}
	manifest, err := schema.DecodeManifest(data)
	schema.SanitizeManifest(manifest)
	if err := schema.ValidateManifest(manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}
