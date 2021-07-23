package commands

{{if $.Mutations.HasCommand }}
import (
	"encoding/json"
	{{$.Project.Name}}stream "{{$.Project.GoModules}}/pkg/stream"
	"github.com/go-gulfstream/gulfstream/pkg/command"
	"github.com/google/uuid"
)

const (
   {{range $.Mutations.Commands -}}
       {{.Command.Name}} = "{{.Command.Name}}Command"
   {{end}}
)

func init() {
    {{range $.Mutations.Commands -}}
        {{ if .Command.Payload -}}
            command.RegisterCodec({{.Command.Name}}, &{{.Command.Payload}}{})
        {{end -}}
    {{end -}}
}

{{range $.Mutations.Commands -}}
   {{ if .Command.Payload -}}
     type {{.Command.Payload}} struct {
     }

     func (c *{{.Command.Payload}}) MarshalBinary() ([]byte, error) {
     	return json.Marshal(c)
     }

     func (c *{{.Command.Payload}}) UnmarshalBinary(data []byte) error {
     	return json.Unmarshal(data, c)
     }
   {{end}}
{{end}}

{{range $.Mutations.Commands -}}
    {{ if .Command.Payload -}}
       func New{{.Command.Name}}(streamID uuid.UUID, c *{{.Command.Payload}}) *command.Command {
       	   return command.New({{.Command.Name}}, {{$.Project.Name}}stream.Name, streamID, c)
       }
    {{else}}
       func New{{.Command.Name}}(streamID uuid.UUID) *command.Command {
           return command.New(RegisterSession, {{$.Project.Name}}stream.Name, streamID, nil)
       }
    {{end}}
{{end}}

{{end}}