package addon

import (
   "context"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   "github.com/google/uuid"
)

type Projection interface {
   {{ if .Event.Payload -}}
       {{ .Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *{{$.EventsPkgName}}.{{.Event.Payload}}) error
   {{else -}}
       {{ .Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error
   {{end -}}
}