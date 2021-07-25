package stream

import (
	"context"

    {{if $.Mutations.HasCommands}}
	   "{{$.GoModules}}/pkg/{{$.CommandsPkgName}}"
	   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
	{{end}}

	{{if $.Mutations.HasEvents}}
	   gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
	   {{range $.ImportEvents}}
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
        {{if eq .Create "yes" -}}
        controller.AddCommandController(
        	{{$.CommandsPkgName}}.{{.Command.Name}},
        	{{.Mutation}}CommandController(m),
        	gulfstream.WithCommandControllerCreateIfNotExists(),
        )
        {{else if eq .Delete "yes" -}}
        controller.AddCommandController(
            {{$.CommandsPkgName}}.{{.Command.Name}},
            {{.Mutation}}CommandController(m),
            gulfstream.WithCommandControllerDropStream(),
        )
        {{else -}}
         controller.AddCommandController(
             {{$.CommandsPkgName}}.{{.Command.Name}},
             {{.Mutation}}CommandController(m),
         )
        {{end -}}
    {{end -}}

    {{if $.Mutations.HasEvents}}
        {{range $.Mutations.Commands -}}
            {{if eq .Create "yes" -}}
            controller.AddEventController(
            	{{$.EventsPkgName}}.{{.Event.Name}},
            	{{.Mutation}}EventController(m),
            	gulfstream.WithEventControllerCreateIfNotExists(),
            )
            {{ else if eq .Delete "yes" -}}
            controller.AddEventController(
                {{$.EventsPkgName}}.{{.Event.Name}},
                {{.Mutation}}EventController(m),
                gulfstream.WithEventControllerDropStream(),
            )
            {{else -}}
            controller.AddEventController(
                 {{$.EventsPkgName}}.{{.Event.Name}},
                 {{.Mutation}}EventController(m),
            )
            {{end}}
        {{end}}
    {{end}}

	return controller
}

{{if $.Mutations.HasCommands}}
    {{range $.Mutations.Commands -}}
         func {{.Mutation}}CommandController(m Mutation) gulfstream.ControllerFunc {
         	return func(ctx context.Context, s *gulfstream.Stream, c *gulfstreamcommand.Command) (*gulfstreamcommand.Reply, error) {
                {{if .Command.Payload -}}
                    {{if .Event.Payload -}}
                        e, err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State(), c.Payload().(*{{$.CommandsPkgName}}.{{.Command.Payload}}))
                        if err != nil {
                           return c.ReplyErr(err), nil
                        }
                        s.Mutate({{$.EventsPkgName}}.{{.Event.Name}}, e)
                        return c.ReplyOk(s.Version()), nil
                    {{else -}}
                        err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State(), c.Payload().(*{{$.CommandsPkgName}}.{{.Command.Payload}}))
                        if err != nil {
                           return c.ReplyErr(err), nil
                        }
                        s.Mutate({{$.EventsPkgName}}.{{.Event.Name}}, nil)
                        return c.ReplyOk(s.Version()), nil
                    {{end -}}
                {{else -}}
                    {{if .Event.Payload -}}
                         e, err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State())
                         if err != nil {
                             return c.ReplyErr(err), nil
                         }
                         s.Mutate({{$.EventsPkgName}}.{{.Event.Name}}, e)
                         return c.ReplyOk(s.Version()), nil
                    {{else -}}
                        err := m.{{.Mutation}}(ctx, c.StreamID(), c.ID(), s.State())
                        if err != nil {
                            return c.ReplyErr(err), nil
                        }
                        s.Mutate({{$.EventsPkgName}}.{{.Event.Name}}, nil)
                        return c.ReplyOk(s.Version()), nil
                    {{end -}}
                {{end -}}
         	}
         }
    {{end -}}
{{end -}}

{{if $.Mutations.HasEvents}}
   {{range $.Mutations.Events -}}
      func {{.Mutation}}EventController(m Mutation) gulfstream.EventController {
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

