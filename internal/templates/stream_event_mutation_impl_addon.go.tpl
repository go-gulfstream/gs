package addon

import (
   "context"
   "github.com/google/uuid"
   "{{$.GoModules}}/pkg/{{$.EventsPkgName}}"
    {{range $.ImportEvents}}
        "{{.Path -}}"
    {{end}}
)

{{if .InEvent.Payload -}}
   func (m *eventMutation) {{.Mutation}}(
      ctx context.Context,
      streamID uuid.UUID,
      eventID uuid.UUID,
      s State,
      c *{{.InEvent.Payload}},
      ) error {
         return nil
      }
{{else -}}
   func (m *eventMutation) {{.Mutation}}(
   ctx context.Context,
   streamID uuid.UUID,
   eventID uuid.UUID,
   s State,
   ) error {
      return nil
   }
{{end -}}