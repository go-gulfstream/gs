package schema

import "time"

type Manifest struct {
	CreatedAt        time.Time         `yml:"created_at"`
	ProjectName      string            `yml:"project_name"`
	GoMod            string            `yml:"go_mod"`
	CommandMutations []CommandMutation `yml:"command_mutations"`
	EventMutations   []EventMutation   `yml:"event_mutations"`
	Author           string            `yml:"author"`
	Email            string            `yml:"email"`
	Description      string            `yml:"description"`
}

type CommandMutation struct {
	Name    string `yml:"name"`
	Command string `yml:"command"`
	Event   string `yml:"event"`
}

type EventMutation struct {
	Name  string `yml:"name"`
	Event string `yml:"event"`
}
