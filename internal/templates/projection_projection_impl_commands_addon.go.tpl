package addon

import (
   "context"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
)

{{ if .Event.Payload -}}
      func(p *projection){{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *{{$.EventsPkgName}}.{{.Event.Payload}}) error {
          return nil
      }
{{else -}}
      func(p *projection){{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error {
          return nil
      }
{{end}}