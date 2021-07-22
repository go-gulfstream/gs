package schema

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/go-gulfstream/gs/internal/templates"
)

var funcMap = template.FuncMap{
	"renderProjectionMethod": renderProjectionMethod,
}

func renderProjectionMethod(m CommandMutation) string {
	return fmt.Sprintf("")
}

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
