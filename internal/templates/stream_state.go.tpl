package stream

import (
   gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
   {{if $.Mutations.HasCommands}}
       "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   {{end}}
   "encoding/json"
)

type State interface {
}

type root struct {
   // domain state
}

func New() gulfstream.State {
	return new(root)
}

func (s *root) Mutate(e *gulfstreamevent.Event) {
    {{if $.Mutations.HasCommands -}}
        switch e.Name() {
        {{range $.Mutations.Commands -}}
           case {{$.EventsPkgName}}.{{.Event.Name}}:
              {{if .Event.Payload -}}
              payload := e.Payload().(*{{$.EventsPkgName}}.{{.Event.Payload}})
              _ = payload
              {{end -}}
        {{end -}}
        }
    {{end -}}
}

func (s *root) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *root) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}