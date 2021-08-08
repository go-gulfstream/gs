package addon

import (
	"context"

	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	"{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
)

func render() {
   projection.AddEventController(
       {{$.EventsPkgName}}.{{.OutEvent.Name}},
       {{.OutEvent.LcFirstName}}Controller(p),
   )
}

{{if .OutEvent.Payload}}
   func {{.OutEvent.LcFirstName}}Controller(p Projection) gulfstream.EventHandlerFunc {
       return func(ctx context.Context, e *gulfstreamevent.Event) error {
           return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version(), e.Payload().(*{{$.EventsPkgName}}.{{.OutEvent.Payload}}))
       }
   }
{{else}}
   func {{.OutEvent.LcFirstName}}Controller(p Projection) gulfstream.EventHandlerFunc {
          return func(ctx context.Context, e *gulfstreamevent.Event) error {
              return p.{{.Mutation}}(ctx, e.StreamID(), e.ID(), e.Version())
          }
   }
{{end}}