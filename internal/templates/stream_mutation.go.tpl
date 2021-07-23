package stream

import (
   "context"
   "github.com/google/uuid"
   {{if $.Mutations.HasCommand}}
       {{$.Project.Name}}events "{{$.Project.GoModules}}/pkg/events"
       {{$.Project.Name}}commands "{{$.Project.GoModules}}/pkg/commands"
   {{end}}
)

type Mutation interface {
    {{range $.Mutations.Commands -}}
        {{if .Command.Payload}}
            {{if .Event.Payload -}}
                {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *{{$.Project.Name}}commands.{{.Command.Payload}}) (*{{$.Project.Name}}events.{{.Event.Payload}}, error)
            {{else -}}
                {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *{{$.Project.Name}}commands.{{.Command.Payload}})  error
            {{end}}
        {{else -}}
            {{if .Event.Payload -}}
                {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State) (*{{$.Project.Name}}events.{{.Event.Payload}}, error)
            {{else -}}
                {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State)  error
            {{end -}}
        {{end -}}
    {{end -}}
}

func NewMutation() Mutation {
	return &mutation{}
}

type mutation struct {
    // indexes, clients, etc...
}

 {{range $.Mutations.Commands -}}
        {{if .Command.Payload}}
            {{if .Event.Payload -}}
                func (m *mutation) {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *{{$.Project.Name}}commands.{{.Command.Payload}}) (*{{$.Project.Name}}events.{{.Event.Payload}}, error) {
                    return &{{$.Project.Name}}events.{{.Event.Payload}}{}, nil
                }
            {{else -}}
                func (m *mutation) {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State, c *{{$.Project.Name}}commands.{{.Command.Payload}})  error {
                    return nil
                }
            {{end}}
        {{else -}}
            {{if .Event.Payload -}}
                func (m *mutation) {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State) (*{{$.Project.Name}}events.{{.Event.Payload}}, error) {
                    return &{{$.Project.Name}}events.{{.Event.Payload}}{}, nil
                }
            {{else -}}
                func (m *mutation) {{.Mutation}}(ctx context.Context, streamID uuid.UUID, commandID uuid.UUID, s State)  error {
                    return nil
                }
            {{end -}}
        {{end -}}
    {{end -}}



