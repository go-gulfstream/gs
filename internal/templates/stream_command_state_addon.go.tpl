package addon

import (
   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"

   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
)

func (s *root) Mutate(e *gulfstreamevent.Event) {
    switch e.Name() {
         case {{$.EventsPkgName}}.{{.Event.Name}}:
              {{if .Event.Payload -}}
                  payload := e.Payload().(*{{$.EventsPkgName}}.{{.Event.Payload}})
                  s.apply{{.Event.Name}}(payload)
              {{else -}}
                  s.apply{{.Event.Name}}()
              {{end -}}
    }
}

{{if .Event.Payload -}}
     func (s *root) apply{{.Event.Name}}(p *{{$.EventsPkgName}}.{{.Event.Payload}}) {
     }
  {{else}}
     func (s *root) apply{{.Event.Name}}() {
     }
{{end}}