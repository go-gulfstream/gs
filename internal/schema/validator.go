package schema

import (
	"fmt"
	"os"
	"path/filepath"
)

func rfiles(path string, m *Manifest) []File {
	gofiles := make([]File, 0, len(files))
	for _, file := range files {
		if !file.required {
			continue
		}
		file.Path = NormalizePath(file.Path, m)
		file.Path = filepath.Join(path, file.Path)
		gofiles = append(gofiles, file)
	}
	return gofiles
}

func Validate(path string, m *Manifest) error {
	if m == nil {
		return fmt.Errorf("schema: manifest file not found. path: %s/gulfstram.yml", path)
	}
	for _, file := range rfiles(path, m) {
		_, err := os.Stat(file.Path)
		if os.IsNotExist(err) {
			return fmt.Errorf("schema: required file - %v", err)
		}
	}
	return nil
}
