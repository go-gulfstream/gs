package {{$.CommandsPkgName}}

{{if $.Mutations.HasCommands }}
import (
	"encoding/json"
	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"
	gulfstreamcommand "github.com/go-gulfstream/gulfstream/pkg/command"
	"github.com/google/uuid"
)

func init() {
  {{range $.Mutations.Commands -}}
       {{ if .Command.Payload -}}
           gulfstreamcommand.RegisterCodec({{.Command.Name}}, &{{.Command.Payload}}{})
       {{end -}}
  {{end -}}
}

{{range $.Mutations.Commands -}}
   {{ if .Command.Payload -}}
     func (c *{{.Command.Payload}}) MarshalBinary() ([]byte, error) {
     	return json.Marshal(c)
     }

     func (c *{{.Command.Payload}}) UnmarshalBinary(data []byte) error {
     	return json.Unmarshal(data, c)
     }
   {{end}}
{{end}}

{{end}}