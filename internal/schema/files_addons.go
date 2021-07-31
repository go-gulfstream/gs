package schema

const (
	CommandMutationAddon   = "CommandMutation"
	CommandControllerAddon = "CommandController"
	EventMutationAddon     = "EventMutation"
	EventControllerAddon   = "EventController"
	StateAddon             = "State"
	CommandsAddon          = "Commands"
	EventsAddon            = "Events"
)

var commandMutationAddons = []File{
	{
		Path:     "/internal/stream/command_mutation.go",
		Template: "stream_command_mutation_addon.go.tpl",
		Addon:    CommandMutationAddon,
	},
	{
		Path:     "/internal/stream/command_mutation_test.go",
		Template: "stream_command_mutation_test_addon.go.tpl",
		Addon:    CommandMutationAddon,
	},
	{
		Path:     "/internal/stream/command_controller.go",
		Template: "stream_command_controller_addon.go.tpl",
		Addon:    CommandControllerAddon,
	},
	{
		Path:     "/internal/stream/state.go",
		Template: "stream_state_addon.go.tpl",
		Addon:    StateAddon,
	},
	{
		Path:     "/pkg/{commands_package}/commands.go",
		Template: "pkg_commands_addon.go.tpl",
		Addon:    CommandsAddon,
	},
	{
		Path:     "/pkg/{events_package}/events.go",
		Template: "pkg_events_addon.go.tpl",
		Addon:    EventsAddon,
	},
}

var eventMutationAddons = []File{
	{
		Path:     "/internal/stream/event_mutation.go",
		Template: "stream_event_mutation_addon.go.tpl",
		Addon:    EventMutationAddon,
	},
	{
		Path:     "/internal/stream/event_mutation.go",
		Template: "stream_event_mutation_test_addon.go.tpl",
		Addon:    EventMutationAddon,
	},
	{
		Path:     "/internal/stream/event_controller.go",
		Template: "stream_event_controller_addon.go.tpl",
		Addon:    EventControllerAddon,
	},
	{
		Path:     "/internal/stream/state.go",
		Template: "stream_state_addon.go.tpl",
		Addon:    StateAddon,
	},
	{
		Path:     "/pkg/{events_package}/events.go",
		Template: "pkg_events_addon.go.tpl",
		Addon:    EventsAddon,
	},
}
