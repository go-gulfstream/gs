package projection

import (
	"context"

	"github.com/go-gulfstream/gulfstream/pkg/event"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	{{$.Project.Name}}events "{{$.Project.GoModules}}/pkg/events"
)

func NewController(p Projection) *gulfstream.Projection {
	projection := gulfstream.NewProjection()

    {{range $.Mutations.Commands -}}
        projection.AddEventController(
       	   {{$.Project.Name}}events.{{.Event.Name}},
       	   {{.Event.Name}}Controller(p),
       	)
    {{end}}

	return projection
}

{{range $.Mutations.Commands -}}
   {{if .Event.Payload}}
   func {{.Event.Name}}Controller(p Projection) gulfstream.EventHandlerFunc {
       return func(ctx context.Context, e *event.Event) error {
           return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version(), e.Payload().(*{{$.Project.Name}}events.{{.Event.Payload}}))
       }
   }
   {{else}}
   func {{.Event.Name}}Controller(p Projection) gulfstream.EventHandlerFunc {
          return func(ctx context.Context, e *event.Event) error {
              return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version())
          }
      }
   {{end}}
{{end}}