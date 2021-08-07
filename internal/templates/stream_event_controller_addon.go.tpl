package addon

import (
  "context"

   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
   {{range $.ImportEvents}}
      "{{.Path}}"
   {{end}}

	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
)

func render() {
{{if eq .Create "yes" -}}
                controller.AddEventController(
                	{{.ControllerName}}(mutation),
                	{{.InEvent.Name}},
                	gulfstream.WithEventControllerCreateIfNotExists(),
                )
{{ else if eq .Delete "yes" -}}
                controller.AddEventController(
                    {{.InEvent.Name}},
                    {{.ControllerName}}(mutation),
                    gulfstream.WithEventControllerDropStream(),
                )
 {{else -}}
                controller.AddEventController(
                     {{.InEvent.Name}},
                     {{.ControllerName}}(mutation),
                )
{{end -}}
}

 func {{.ControllerName}}(m EventMutation) gulfstream.EventController {
      	return gulfstream.EventControllerFunc(
      		func(event *gulfstreamevent.Event) gulfstream.Picker {
      			return gulfstream.Picker{}
      		},
      		func(ctx context.Context, s *gulfstream.Stream, e *gulfstreamevent.Event) (err error) {
      		    {{if .InEvent.Payload -}}
                    return m.{{.Mutation}}(ctx, s.ID(), e.ID(), s.State(), e.Payload().(*{{.InEvent.Payload}}))
                {{else -}}
                    return m.{{.Mutation}}(ctx, s.ID(), e.ID(), s.State())
                {{end -}}
      		})
      }