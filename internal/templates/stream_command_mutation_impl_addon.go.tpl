package addon

import (
   "context"
   "github.com/google/uuid"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
   "{{$.GoModules}}/pkg/{{$.CommandsPkgName}}"
)

{{if .Command.Payload}}
    {{if .Event.Payload -}}
                func (m *commandMutation) {{.Mutation}}(
                  ctx context.Context,
                  streamID uuid.UUID,
                  commandID uuid.UUID,
                  s State, c *{{$.CommandsPkgName}}.{{.Command.Payload}},
                  ) (*{{$.EventsPkgName}}.{{.Event.Payload}}, error) {
                    return &{{$.EventsPkgName}}.{{.Event.Payload}}{}, nil
                }
    {{else -}}
                func (m *mutation) {{.Mutation}}(
                ctx context.Context,
                streamID uuid.UUID,
                commandID uuid.UUID,
                s State,
                c *{{$.CommandsPkgName}}.{{.Command.Payload}},
                )  error {
                    return nil
                }
    {{end}}
{{else -}}
    {{if .Event.Payload -}}
                func (m *mutation) {{.Mutation}}(
                ctx context.Context,
                streamID uuid.UUID,
                commandID uuid.UUID,
                s State,
                ) (*{{$.EventsPkgName}}.{{.Event.Payload}}, error) {
                    return &{{$.EventsPkgName}}.{{.Event.Payload}}{}, nil
                }
    {{else -}}
                func (m *mutation) {{.Mutation}}(
                 ctx context.Context,
                 streamID uuid.UUID,
                 commandID uuid.UUID,
                 s State,
                 )  error {
                    return nil
                }
    {{end}}
{{end}}