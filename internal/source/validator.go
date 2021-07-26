package source

import (
	"path/filepath"

	"github.com/go-gulfstream/gs/internal/schema"
)

var streamFiles = []string{
	"/internal/projection/projection.go",
	"/internal/projection/controller.go",
	"/internal/stream/mutation.go",
	"/internal/stream/controller.go",
	"/internal/stream/state.go",
	"/pkg/{events_package}/events.go",
	"/pkg/{commands_package}/commands.go",
}

func iterStreamFiles(path string, m *schema.Manifest, fn func(filename string) error) error {
	for _, filename := range streamFiles {
		filename = schema.NormalizePath(filename, m)
		filename = filepath.Join(path, filename)
		if err := fn(filename); err != nil {
			return err
		}
	}
	return nil
}

func FilesValidation(path string, m *schema.Manifest) error {
	return nil
}
