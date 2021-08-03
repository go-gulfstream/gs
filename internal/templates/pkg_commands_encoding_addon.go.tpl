package addon

import (
	"encoding/json"
	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"
	gulfstreamcommand "github.com/go-gulfstream/gulfstream/pkg/command"
	"github.com/google/uuid"
)

func init() {
   {{ if .Command.Payload -}}
       gulfstreamcommand.RegisterCodec({{.Command.Name}}, &{{.Command.Payload}}{})
   {{end -}}
}

{{ if .Command.Payload -}}
     func (c *{{.Command.Payload}}) MarshalBinary() ([]byte, error) {
     	return json.Marshal(c)
     }

     func (c *{{.Command.Payload}}) UnmarshalBinary(data []byte) error {
     	return json.Unmarshal(data, c)
     }
{{end}}