package templates

import (
	"embed"
	"text/template"
)

//go:embed *.tpl
var tpls embed.FS

func Parse(tpl string) (*template.Template, error) {
	return template.ParseFS(tpls, tpl)
}
