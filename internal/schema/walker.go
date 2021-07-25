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
		file.Path = replacePath(file.Path, m)
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

func replacePath(filePath string, m *Manifest) string {
	filePath = strings.ReplaceAll(filePath, "{package}", m.PackageName)
	filePath = strings.ReplaceAll(filePath, "{commands_package}", m.CommandsPkgName)
	filePath = strings.ReplaceAll(filePath, "{events_package}", m.EventsPkgName)
	filePath = strings.ReplaceAll(filePath, "{stream_package}", m.StreamPkgName)
	return filePath
}
