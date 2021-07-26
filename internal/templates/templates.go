package templates

import (
	"embed"
	"text/template"
)

//go:embed *.tpl
var tpls embed.FS

func Parse(tpl string) (*template.Template, error) {
	newTpl, err := template.ParseFS(tpls, tpl)
	if err != nil {
		return nil, err
	}
	return newTpl, err
}
