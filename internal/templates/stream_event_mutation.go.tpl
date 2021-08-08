package stream

import (
	"context"

	{{if $.Mutations.HasEvents}}
	   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
	   {{range $.ImportEvents}}
	       "{{.}}"
	   {{end}}
	{{end}}

	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
)

type EventMutation interface {
    {{range $.Mutations.Events -}}
       {{if .InEvent.Payload -}}
           {{.Mutation}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, s State, c *{{.InEvent.Payload}}) error
       {{else -}}
           {{.Mutation}}(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, s State) error
       {{end -}}
    {{end -}}
}

func NewEventMutation() EventMutation {
	return &eventMutation{}
}

type eventMutation struct {
    // indexes, clients, etc...
}

{{range $.Mutations.Events -}}
{{if .InEvent.Payload -}}
   func (m *eventMutation) {{.Mutation}}(
      ctx context.Context,
      streamID uuid.UUID,
      eventID uuid.UUID,
      s State,
      c *{{.InEvent.Payload}},
      ) error {
         return nil
      }
{{else -}}
   func (m *eventMutation) {{.Mutation}}(
   ctx context.Context,
   streamID uuid.UUID,
   eventID uuid.UUID,
   s State,
   ) error {
      return nil
   }
{{end -}}
{{end}}
