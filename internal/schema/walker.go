package schema

import (
	"path/filepath"
	"strings"
)

func Walk(path string, m *Manifest, fn func(File) error) error {
	if m == nil {
		return nil
	}
	for _, file := range files {
		file.Path = strings.ReplaceAll(file.Path, "%s", m.PackageName)
		file.Path = filepath.Join(path, file.Path)
		if len(file.Template) > 0 {
			file.HasTemplate = true
			tplData, err := renderTemplate(file.Template, m)
			if err != nil {
				return err
			}
			file.TemplateData = tplData
		}
		if err := fn(file); err != nil {
			return err
		}
	}
	return nil
}
