package schema

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/go-gulfstream/gs/internal/templates"
)

func Walk(path string, m *Manifest, fn func(File) error) error {
	for _, file := range files {
		file.Path = strings.ReplaceAll(file.Path, "%s", m.Project.Name)
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

func renderTemplate(fileTpl string, m *Manifest) ([]byte, error) {
	tpl, err := templates.Parse(fileTpl)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tpl.Execute(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
