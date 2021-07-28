package projection

import (
	"context"

	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	"{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
)

func NewController(p Projection) *gulfstream.Projection {
	projection := gulfstream.NewProjection()

    {{range $.Mutations.Commands -}}
        projection.AddEventController(
       	   {{$.EventsPkgName}}.{{.Event.Name}},
       	   {{.Event.LcFirstName}}Controller(p),
       	)

    {{end}}

    {{range $.Mutations.Events -}}
        projection.AddEventController(
           {{$.EventsPkgName}}.{{.OutEvent.Name}},
           {{.OutEvent.LcFirstName}}Controller(p),
        )

    {{end}}

	return projection
}

{{range $.Mutations.Commands -}}
   {{if .Event.Payload}}
   func {{.Event.LcFirstName}}Controller(p Projection) gulfstream.EventHandlerFunc {
       return func(ctx context.Context, e *gulfstreamevent.Event) error {
           return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version(), e.Payload().(*{{$.EventsPkgName}}.{{.Event.Payload}}))
       }
   }
   {{else}}
   func {{.Event.LcFirstName}}Controller(p Projection) gulfstream.EventHandlerFunc {
          return func(ctx context.Context, e *gulfstreamevent.Event) error {
              return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version())
          }
      }
   {{end}}
{{end}}

{{range $.Mutations.Events -}}
   {{if .OutEvent.Payload}}
   func {{.OutEvent.LcFirstName}}Controller(p Projection) gulfstream.EventHandlerFunc {
       return func(ctx context.Context, e *gulfstreamevent.Event) error {
           return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version(), e.Payload().(*{{$.EventsPkgName}}.{{.Event.Payload}}))
       }
   }
   {{else}}
   func {{.OutEvent.LcFirstName}}Controller(p Projection) gulfstream.EventHandlerFunc {
          return func(ctx context.Context, e *gulfstreamevent.Event) error {
              return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version())
          }
      }
   {{end}}
{{end}}