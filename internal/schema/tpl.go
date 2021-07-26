package schema

import (
	"bytes"

	"github.com/go-gulfstream/gs/internal/templates"
)

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
