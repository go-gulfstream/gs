package addon

const (
  {{.Event.Name}} = "{{.Event.Name}}Event"
)

{{ if .Event.Payload -}}
         type {{.Event.Payload}} struct {
         }
{{end}}