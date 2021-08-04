package addon

import (
	"encoding/json"
	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
)

func init() {
       {{ if .Event.Payload -}}
           gulfstreamevent.RegisterCodec({{.Event.Name}}, &{{.Event.Payload}}{})
       {{end -}}
}


{{ if .Event.Payload -}}
           func (c *{{.Event.Payload}}) MarshalBinary() ([]byte, error) {
                return json.Marshal(c)
           }

           func (c *{{.Event.Payload}}) UnmarshalBinary(data []byte) error {
                return json.Unmarshal(data, c)
           }
{{end -}}