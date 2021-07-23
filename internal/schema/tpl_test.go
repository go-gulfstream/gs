package schema

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplate_RenderProjectionInterface(t *testing.T) {
	manifest := new(Manifest)
	manifest.Project.Name = "myproject"
	manifest.Project.GoModules = "modules"
	manifest.Mutations.Commands = []CommandMutation{
		{
			Mutation: "myMethod",
			Command: Command{
				Name:    "Add",
				Payload: "AddPayload",
			},
			Event: Event{
				Name:    "Added",
				Payload: "AddedPayload",
			},
		},
		{
			Mutation: "myMethod2",
			Command: Command{
				Name:    "Mul",
				Payload: "MulPayload",
			},
			Event: Event{
				Name:    "Completed",
				Payload: "CompletedPayload",
			},
		},
		{
			Mutation: "myMethod3",
			Command: Command{
				Name: "Sub",
			},
		},
	}
	manifest.Sanitize()

	tpl, err := template.ParseFiles("./testdata/projection_interface.go.tpl")
	assert.NoError(t, err)
	buf := bytes.NewBuffer(nil)
	assert.NoError(t, tpl.Funcs(funcMap).Execute(buf, manifest))
	assert.Contains(t, buf.String(), "MyMethod(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *events.AddedPayload) error")
	assert.Contains(t, buf.String(), "MyMethod2(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int, e *events.CompletedPayload) error")
	assert.Contains(t, buf.String(), "MyMethod3(ctx context.Context, streamID uuid.UUID, eventID uuid.UUID, version int) error")
}
