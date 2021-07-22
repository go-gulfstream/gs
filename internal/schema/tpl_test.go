package schema

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplate_RenderProjectionInterface(t *testing.T) {
	manifest := new(Manifest)
	manifest.Mutations.Commands = []CommandMutation{
		{
			Name:    "myMethod",
			Command: "Add",
			Event:   "Added",
		},
		{
			Name:    "myMethod2",
			Command: "Mul",
			Event:   "Completed",
		},
	}
	manifest.Sanitize()

	tpl, err := template.ParseFiles("./testdata/projection_interface.go.tpl")
	assert.NoError(t, err)
	buf := bytes.NewBuffer(nil)
	assert.NoError(t, tpl.Funcs(funcMap).Execute(buf, manifest))
	assert.Contains(t, buf.String(), "MyMethod(ctx context.Context, e *event.Event) error")
	assert.Contains(t, buf.String(), "MyMethod2(ctx context.Context, e *event.Event) error")
}
