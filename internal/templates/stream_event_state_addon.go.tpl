package stream

import (
   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"

   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"

   {{range .Imports}}
       {{.}}
   {{end}}
)

 {{if .OutEvent.Payload -}}
     func (s *root) apply{{.OutEvent.Name}}(p *{{$.EventsPkgName}}.{{.OutEvent.Payload}}) {
     }
  {{else}}
     func (s *root) apply{{.OutEvent.Name}}() {
     }
  {{end}}

  func (s *root) Mutate(e *gulfstreamevent.Event) {
      switch e.Name() {
         case {{$.EventsPkgName}}.{{.OutEvent.Name}}:
           {{if .OutEvent.Payload -}}
               payload := e.Payload().(*{{$.EventsPkgName}}.{{.OutEvent.Payload}})
               s.apply{{.OutEvent.Name}}(payload)
           {{else -}}
               s.apply{{.OutEvent.Name}}()
           {{end -}}
      }
  }