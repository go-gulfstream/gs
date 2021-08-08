package projection

import (
    "testing"
)

{{range $.Mutations.Commands -}}
func TestProjection_{{.Mutation}}(t *testing.T) {
}
{{end}}

{{range $.Mutations.Events -}}
func TestProjection_{{.Mutation}}(t *testing.T) {
}
{{end}}