package addon

import (
   "context"
   "github.com/google/uuid"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName -}}"
   {{range $.ImportEvents}}
       "{{.Path}}"
   {{end}}
)

type EventMutation interface {
   {{if .InEvent.Payload -}}
       {{.Mutation}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, s State, c *{{.InEvent.Payload}}) error
   {{else -}}
       {{.Mutation}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, s State) error
    {{end -}}
}