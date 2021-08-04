package addon

const (
   {{.OutEvent.Name}} = "{{.OutEvent.Name}}Event"
)

{{ if .OutEvent.Payload -}}
         type {{.OutEvent.Payload}} struct {
         }
{{end}}