package projection

import (
   "context"
   {{ if $.Mutations.HasCommand -}}
      "{{$.Project.GoModules}}/pkg/events"
   {{end -}}
   "github.com/google/uuid"
)

type Projection interface {
    {{range $.Mutations.Commands -}}
        {{ if .Event.Payload -}}
           {{ .Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *events.{{.Event.Payload}}) error
        {{else -}}
           {{ .Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error
        {{end -}}
    {{end -}}
}

func New(
	storage *Storage,
) Projection {
	return &projection{
		storage: storage,
	}
}

type projection struct {
	storage *Storage
}

{{range $.Mutations.Commands -}}
   {{ if .Event.Payload -}}
      func(p *projection){{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *events.{{.Event.Payload}}) error {
          return nil
      }
   {{else -}}
      func(p *projection){{.Mutation -}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error {
          return nil
      }
   {{end}}
{{end -}}