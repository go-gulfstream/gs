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

func MakeEventControllers(
	mutation EventMutation,
    controller *gulfstream.Mutator,
) {

   {{if $.Mutations.HasEvents}}
        {{range $.Mutations.Events -}}
            {{if eq .Create "yes" -}}
            controller.AddEventController(
            	{{.InEvent.Name}},
            	{{.ControllerName}}(mutation),
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
            {{end}}
        {{end}}
    {{end -}}
}

{{if $.Mutations.HasEvents}}
   {{range $.Mutations.Events -}}
      func {{.ControllerName}}(m EventMutation) gulfstream.EventController {
      	return gulfstream.EventControllerFunc(
      		func(event *gulfstreamevent.Event) gulfstream.Picker {
      			return gulfstream.Picker{}
      		},
      		func(ctx context.Context, s *gulfstream.Stream, e *gulfstreamevent.Event) error {
      			return nil
      		})
      }
   {{end}}
{{end}}