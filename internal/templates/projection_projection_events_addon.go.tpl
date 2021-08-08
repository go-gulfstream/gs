package addon

import (
   "context"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   "github.com/google/uuid"
)

type Projection interface {
   {{ if .OutEvent.Payload -}}
       {{ .Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *{{$.EventsPkgName}}.{{.OutEvent.Payload}}) error
   {{else -}}
       {{ .Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error
   {{end -}}
}