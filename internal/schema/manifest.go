package schema

import "time"

type Manifest struct {
	CreatedAt   time.Time  `yml:"created_at"`
	ProjectName string     `yml:"project_name"`
	StreamName  string     `yml:"stream_name"`
	GoMod       string     `yml:"go_mod"`
	Mutations   []Mutation `yml:"mutations"`
	Author      string     `yml:"author"`
	Email       string     `yml:"email"`
	Description string     `yml:"description"`
}

type Mutation struct {
	Name           string
	Command        string
	Event          string
	WithOutPayload bool
}
