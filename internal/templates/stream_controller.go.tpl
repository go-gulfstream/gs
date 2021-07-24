package stream

import (
	"context"

    {{if $.Mutations.HasCommand}}
	   {{$.Project.Name}}commands  "{{$.Project.GoModules}}/pkg/commands"
	   {{$.Project.Name}}events "{{$.Project.GoModules}}/pkg/events"
	{{end}}

	{{if $.Mutations.HasEvent}}
	   "github.com/go-gulfstream/gulfstream/pkg/event"
	   {{range $.EventExternalPackages}}
	       "{{.}}"
	   {{end}}
	{{end}}

	gulfstreamcommand "github.com/go-gulfstream/gulfstream/pkg/command"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
)

func NewController(
	m Mutation,
	s gulfstream.Storage,
	p gulfstream.Publisher,
	o ...gulfstream.MutatorOption,
) *gulfstream.Mutator {
	controller := gulfstream.NewMutator(s, p, o...)

    {{range $.Mutations.Commands -}}
        {{if .Operations.Create -}}
        controller.AddCommandController(
        	{{$.Project.Name}}commands.{{.Command.Name}},
        	{{.Mutation}}CommandController(m),
        	gulfstream.WithCommandControllerCreateIfNotExists(),
        )
        {{else if .Operations.Delete}}
        controller.AddCommandController(
            {{$.Project.Name}}commands.{{.Command.Name}},
            {{.Mutation}}CommandController(m),
            gulfstream.WithCommandControllerDropStream(),
        )
        {{else -}}
         controller.AddCommandController(
             {{$.Project.Name}}commands.{{.Command.Name}},
             {{.Mutation}}CommandController(m),
         )
        {{end -}}
    {{end -}}

    {{if $.Mutations.HasEvent}}
        {{range $.Mutations.Commands -}}
            {{if .Operations.Create -}}
            controller.AddEventController(
            	events.{{.Event.Name}},
            	{{.Mutation}}EventController(m),
            	stream.WithEventControllerCreateIfNotExists(),
            )
            {{ else if .Operations.Delete -}}
            controller.AddEventController(
                events.{{.Event.Name}},
                {{.Mutation}}EventController(m),
                stream.WithEventControllerDropStream(),
            )
            {{else}}
            controller.AddEventController(
                 events.{{.Event.Name}},
                 {{.Mutation}}EventController(m),
            )
            {{end}}
        {{end}}
    {{end}}

	return controller
}

{{if $.Mutations.HasCommand}}
    {{range $.Mutations.Commands -}}
         func {{.Mutation}}CommandController(m Mutation) gulfstream.ControllerFunc {
         	return func(ctx context.Context, s *gulfstream.Stream, c *gulfstreamcommand.Command) (*gulfstreamcommand.Reply, error) {
                {{if .Command.Payload -}}
                    {{if .Event.Payload -}}
                        e, err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State(), c.Payload().(*{{$.Project.Name}}commands.{{.Command.Payload}}))
                        if err != nil {
                           return c.ReplyErr(err), nil
                        }
                        s.Mutate({{$.Project.Name}}events.{{.Event.Name}}, e)
                        return c.ReplyOk(s.Version()), nil
                    {{else -}}
                        err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State(), c.Payload().(*{{$.Project.Name}}commands.{{.Command.Payload}}))
                        if err != nil {
                           return c.ReplyErr(err), nil
                        }
                        s.Mutate({{$.Project.Name}}events.{{.Event.Name}}, nil)
                        return c.ReplyOk(s.Version()), nil
                    {{end -}}
                {{else -}}
                    {{if .Event.Payload -}}
                         e, err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State())
                         if err != nil {
                             return c.ReplyErr(err), nil
                         }
                         s.Mutate({{$.Project.Name}}events.{{.Event.Name}}, e)
                         return c.ReplyOk(s.Version()), nil
                    {{else -}}
                        err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State())
                        if err != nil {
                            return c.ReplyErr(err), nil
                        }
                        s.Mutate({{$.Project.Name}}events.{{.Event.Name}}, nil)
                        return c.ReplyOk(s.Version()), nil
                    {{end -}}
                {{end -}}
         	}
         }
    {{end -}}
{{end -}}

{{if $.Mutations.HasEvent}}
   {{range $.Mutations.Events -}}
      func {{.Mutation}}EventController(m Mutation) gulfstream.EventController {
      	return gulfstream.EventControllerFunc(
      		func(event *event.Event) stream.Picker {
      			return gulfstream.Picker{}
      		},
      		func(ctx context.Context, s *gulfstream.Stream, e *event.Event) error {
      			return nil
      		})
      }
   {{end}}
{{end}}

