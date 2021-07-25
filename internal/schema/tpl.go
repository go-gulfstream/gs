package schema

import (
	"bytes"
	"text/template"

	"github.com/go-gulfstream/gs/internal/templates"
)

var funcMap = template.FuncMap{}

func renderTemplate(fileTpl string, m *Manifest) ([]byte, error) {
	tpl, err := templates.Parse(fileTpl)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tpl.Funcs(funcMap).Execute(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
