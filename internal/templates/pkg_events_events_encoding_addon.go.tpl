package addon

import (
   "encoding/json"
	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
)

func init() {
       {{ if .OutEvent.Payload -}}
           gulfstreamevent.RegisterCodec({{.OutEvent.Name}}, &{{.OutEvent.Payload}}{})
       {{end -}}
}

{{ if .OutEvent.Payload -}}
    func (c *{{.OutEvent.Payload}}) MarshalBinary() ([]byte, error) {
      return json.Marshal(c)
    }

    func (c *{{.OutEvent.Payload}}) UnmarshalBinary(data []byte) error {
       return json.Unmarshal(data, c)
    }
{{end}}