package test

type Projection interface {
    {{range $.Mutations.Commands -}}
        {{.Name -}}(ctx context.Context, e *event.Event) error
    {{end}}
}