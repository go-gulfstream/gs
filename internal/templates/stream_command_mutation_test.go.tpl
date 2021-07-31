package stream

import (
    "testing"
)

{{range $.Mutations.Commands -}}
func TestStreamCommandMutation_{{.Mutation}}(t *testing.T) {
}

{{end}}