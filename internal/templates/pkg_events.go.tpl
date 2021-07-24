package events

{{if $.Mutations.HasCommand }}
import (
	"encoding/json"
	"github.com/go-gulfstream/gulfstream/pkg/event"
)

const (
   {{range $.Mutations.Commands -}}
       {{.Event.Name}} = "{{.Event.Name}}Event"
   {{end}}
)

func init() {
    {{range $.Mutations.Commands -}}
        {{ if .Event.Payload -}}
            event.RegisterCodec({{.Event.Name}}, &{{.Event.Payload}}{})
        {{end -}}
    {{end -}}
}

{{range $.Mutations.Commands -}}
   {{ if .Event.Payload -}}
     type {{.Event.Payload}} struct {
     }

     func (c *{{.Event.Payload}}) MarshalBinary() ([]byte, error) {
     	return json.Marshal(c)
     }

     func (c *{{.Event.Payload}}) UnmarshalBinary(data []byte) error {
     	return json.Unmarshal(data, c)
     }
   {{end}}
{{end}}

{{end}}