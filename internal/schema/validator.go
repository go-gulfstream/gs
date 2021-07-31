package schema

import (
	"fmt"
	"os"
	"path/filepath"
)

func Validate(path string, m *Manifest) error {
	if m == nil {
		return fmt.Errorf("schema: manifest file not found. path: %s/gulfstram.yml", path)
	}
	for _, file := range files {
		if !file.required {
			continue
		}
		file.Path = NormalizePath(file.Path, m)
		file.Path = filepath.Join(path, file.Path)

		_, err := os.Stat(file.Path)
		if os.IsNotExist(err) {
			return fmt.Errorf("schema: required file - %v", err)
		}
	}
	return nil
}
