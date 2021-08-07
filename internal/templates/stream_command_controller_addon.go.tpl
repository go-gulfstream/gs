package addon

import (
	"context"

	"{{$.GoModules}}/pkg/{{$.CommandsPkgName}}"
	"{{$.GoModules}}/pkg/{{$.EventsPkgName}}"

	gulfstreamcommand "github.com/go-gulfstream/gulfstream/pkg/command"
	gulfstream "github.com/go-gulfstream/gulfstream/pkg/stream"
)

func render() {
{{if eq .Create "yes" -}}
        controller.AddCommandController(
        	{{$.CommandsPkgName}}.{{.Command.Name}},
        	{{.ControllerName}}(mutation),
        	gulfstream.WithCommandControllerCreateIfNotExists(),
        )
{{else if eq .Delete "yes" -}}
        controller.AddCommandController(
            {{$.CommandsPkgName}}.{{.Command.Name}},
            {{.ControllerName}}(mutation),
            gulfstream.WithCommandControllerDropStream(),
        )
{{else -}}
         controller.AddCommandController(
             {{$.CommandsPkgName}}.{{.Command.Name}},
             {{.ControllerName}}(mutation),
         )
{{end -}}
}

func {{ .ControllerName }}(m CommandMutation) gulfstream.ControllerFunc {
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
