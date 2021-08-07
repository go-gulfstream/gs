package {{$.EventsPkgName}}

import (
	"encoding/json"
	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"

)

func init() {
  {{range $.Mutations.Commands -}}
       {{ if .Event.Payload -}}
           gulfstreamevent.RegisterCodec({{.Event.Name}}, &{{.Event.Payload}}{})
       {{end -}}
  {{end -}}
    {{range $.Mutations.Events -}}
         {{ if .OutEvent.Payload -}}
             gulfstreamevent.RegisterCodec({{.OutEvent.Name}}, &{{.OutEvent.Payload}}{})
         {{end -}}
    {{end -}}
}

{{range $.Mutations.Commands -}}
   {{ if .Event.Payload -}}
     func (c *{{.Event.Payload}}) MarshalBinary() ([]byte, error) {
     	return json.Marshal(c)
     }

     func (c *{{.Event.Payload}}) UnmarshalBinary(data []byte) error {
     	return json.Unmarshal(data, c)
     }
   {{end}}
{{end}}

{{range $.Mutations.Events -}}
   {{ if .OutEvent.Payload -}}
     func (c *{{.OutEvent.Payload}}) MarshalBinary() ([]byte, error) {
     	return json.Marshal(c)
     }

     func (c *{{.OutEvent.Payload}}) UnmarshalBinary(data []byte) error {
     	return json.Unmarshal(data, c)
     }
   {{end}}
{{end}}