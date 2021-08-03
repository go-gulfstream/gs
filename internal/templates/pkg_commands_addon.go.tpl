package addon

import (
	"{{$.GoModules}}/pkg/{{$.StreamPkgName}}"
	gulfstreamcommand "github.com/go-gulfstream/gulfstream/pkg/command"
	googleuuid "github.com/google/uuid"
)

const (
   {{.Command.Name}} = "{{.Command.Name}}Command"
)

{{ if .Command.Payload -}}
     type {{.Command.Payload}} struct {
     }
{{end}}

{{ if .Command.Payload -}}
       func New{{.Command.Name}}(streamID googleuuid.UUID, c *{{.Command.Payload}}) *gulfstreamcommand.Command {
       	   return gulfstreamcommand.New({{.Command.Name}}, {{$.StreamPkgName}}.Name, streamID, c)
       }
{{else}}
       func New{{.Command.Name}}(streamID googleuuid.UUID) *gulfstreamcommand.Command {
           return gulfstreamcommand.New({{.Command.Name}}, {{$.StreamPkgName}}.Name, streamID, nil)
       }
{{end}}
