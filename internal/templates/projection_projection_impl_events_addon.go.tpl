package addon

import (
   "context"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
)

{{ if .OutEvent.Payload -}}
      func(p *projection){{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *{{$.EventsPkgName}}.{{.OutEvent.Payload}}) error {
          return nil
      }
{{else -}}
      func(p *projection){{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error {
          return nil
      }
{{end}}