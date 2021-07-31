package {{$.EventsPkgName}}

{{if or $.Mutations.HasCommands $.Mutations.HasEvents }}
import (
	"encoding/json"
	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
)

const (
   {{range $.Mutations.Commands -}}
       {{.Event.Name}} = "{{.Event.Name}}Event"
   {{end}}
   {{range $.Mutations.Events -}}
       {{.OutEvent.Name}} = "{{.OutEvent.Name}}Event"
   {{end}}
)

{{range $.Mutations.Commands -}}
   {{ if .Event.Payload -}}
     type {{.Event.Payload}} struct {
     }
   {{end}}
{{end}}

{{range $.Mutations.Events -}}
   {{ if .OutEvent.Payload -}}
     type {{.OutEvent.Payload}} struct {
     }
   {{end}}
{{end}}

{{end}}
