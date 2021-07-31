package stream

import (
    "testing"
)

{{range $.Mutations.Events -}}
func TestStreamEventMutation_{{.Mutation}}(t *testing.T) {
}

{{end}}