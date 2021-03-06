package addon

import (
	"context"

	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
	"{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
)

func render() {
     projection.AddEventController(
       	{{$.EventsPkgName}}.{{.Event.Name}},
       	{{.Event.LcFirstName}}Controller(p),
     )
}

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
{{end -}}