package stream

import (
   gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
   {{if $.Mutations.HasCommands}}
       "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   {{end}}
   {{if $.Mutations.HasEvents}}
          "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
    {{end}}
)

type State interface {
}

type root struct {
   // domain state
}

{{range $.Mutations.Commands -}}
  {{if .Event.Payload -}}
     func (s *root) apply{{.Event.Name}}(p *{{$.EventsPkgName}}.{{.Event.Payload}}) {
     }
  {{else}}
     func (s *root) apply{{.Event.Name}}() {
     }
  {{end}}
{{end}}
{{range $.Mutations.Events -}}
  {{if .OutEvent.Payload -}}
     func (s *root) apply{{.OutEvent.Name}}(p *{{$.EventsPkgName}}.{{.OutEvent.Payload}}) {
     }
  {{else}}
     func (s *root) apply{{.OutEvent.Name}}() {
     }
  {{end}}
{{end}}

func New() gulfstream.State {
	return new(root)
}

func (s *root) Mutate(e *gulfstreamevent.Event) {
    switch e.Name() {
    {{if $.Mutations.HasCommands -}}
        {{range $.Mutations.Commands -}}
           case {{$.EventsPkgName}}.{{.Event.Name}}:
              {{if .Event.Payload -}}
              payload := e.Payload().(*{{$.EventsPkgName}}.{{.Event.Payload}})
              s.apply{{.Event.Name}}(payload)
              {{else -}}
              s.apply{{.Event.Name}}()
              {{end -}}
        {{end -}}
        {{range $.Mutations.Events -}}
            case {{$.EventsPkgName}}.{{.OutEvent.Name}}:
               {{if .OutEvent.Payload -}}
                  payload := e.Payload().(*{{$.EventsPkgName}}.{{.OutEvent.Payload}})
                  s.apply{{.OutEvent.Name}}(payload)
               {{else -}}
                  s.apply{{.OutEvent.Name}}()
               {{end -}}
        {{end -}}
    {{end -}}
    }
}