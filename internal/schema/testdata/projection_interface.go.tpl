package test

import (
   {{ if $.Mutations.HasCommand -}}
      "{{$.Project.GoModules}}/pkg/events"
   {{end -}}
)

type Projection interface {
    {{range $.Mutations.Commands -}}
        {{ if .Event.Payload -}}
           {{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *events.{{.Event.Payload}}) error
        {{else -}}
           {{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error
        {{end -}}
    {{end -}}
}