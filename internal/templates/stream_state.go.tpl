package stream

import (
   gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
   "github.com/go-gulfstream/gulfstream/pkg/event"
   {{if $.Mutations.HasCommand}}
       {{$.Project.Name}}events "{{$.Project.GoModules}}/pkg/events"
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

func (s *root) Mutate(e *event.Event) {
    {{if $.Mutations.HasCommand -}}
        switch e.Name() {
        {{range $.Mutations.Commands -}}
           case {{$.Project.Name}}events.{{.Event.Name}}:
              {{if .Event.Payload -}}
              payload := e.Payload().(*{{$.Project.Name}}events.{{.Event.Payload}})
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


