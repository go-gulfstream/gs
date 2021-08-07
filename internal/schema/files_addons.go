package schema

import "path/filepath"

const (
	CommandMutationAddon           = "CommandMutation"
	CommandMutationProjectionAddon = "CommandMutationProjection"
	CommandMutationImplAddon       = "CommandMutationImpl"
	CommandMutationTestAddon       = "CommandMutationTest"
	CommandControllerAddon         = "CommandController"
	CommandStateAddon              = "CommandState"
	EventMutationAddon             = "EventMutation"
	EventMutationProjectionAddon   = "EventMutationProjection"
	EventMutationTestAddon         = "EventMutationTest"
	EventMutationImplAddon         = "EventMutationImpl"
	EventControllerAddon           = "EventController"
	EventStateAddon                = "EventState"
	CommandsAddon                  = "Commands"
	CommandsEncodingAddon          = "CommandsEncoding"
	CommandsEventsAddon            = "CommandEvents"
	CommandsEventsEncodingAddon    = "CommandEventsEncoding"
	EventsEventsAddon              = "EventsEventsAddon"
	EventsEventsEncodingAddon      = "EventsEventsEncodingAddon"
)

var commandMutationAddons = []File{
	{
		Path:     "/internal/stream/command_mutation.go",
		Template: "stream_command_mutation_addon.go.tpl",
		Addon:    CommandMutationAddon,
		required: true,
	},
	{
		Path:     "/internal/stream/command_mutation.go",
		Template: "stream_command_mutation_impl_addon.go.tpl",
		Addon:    CommandMutationImplAddon,
		required: true,
	},
	{
		Path:     "/internal/stream/command_mutation_test.go",
		Template: "stream_command_mutation_test_addon.go.tpl",
		Addon:    CommandMutationTestAddon,
	},
	{
		Path:     "/internal/stream/command_controller.go",
		Template: "stream_command_controller_addon.go.tpl",
		Addon:    CommandControllerAddon,
		required: true,
	},
	{
		Path:     "/internal/stream/state.go",
		Template: "stream_command_state_addon.go.tpl",
		Addon:    CommandStateAddon,
		required: true,
	},
	{
		Path:     "/pkg/{commands_package}/commands.go",
		Template: "pkg_commands_addon.go.tpl",
		Addon:    CommandsAddon,
	},
	{
		Path:     "/pkg/{commands_package}/commands_encoding.go",
		Template: "pkg_commands_encoding_addon.go.tpl",
		Addon:    CommandsEncodingAddon,
	},
	{
		Path:     "/pkg/{events_package}/events.go",
		Template: "pkg_events_commands_addon.go.tpl",
		Addon:    CommandsEventsAddon,
	},
	{
		Path:     "/pkg/{events_package}/events_encoding.go",
		Template: "pkg_events_commands_encoding_addon.go.tpl",
		Addon:    CommandsEventsEncodingAddon,
	},
}

var eventMutationAddons = []File{
	{
		Path:     "/internal/stream/event_mutation.go",
		Template: "stream_event_mutation_addon.go.tpl",
		Addon:    EventMutationAddon,
		required: true,
	},
	{
		Path:     "/internal/stream/event_mutation.go",
		Template: "stream_event_mutation_impl_addon.go.tpl",
		Addon:    EventMutationImplAddon,
		required: true,
	},
	{
		Path:     "/internal/stream/event_mutation_test.go",
		Template: "stream_event_mutation_test_addon.go.tpl",
		Addon:    EventMutationTestAddon,
	},
	{
		Path:     "/internal/stream/event_controller.go",
		Template: "stream_event_controller_addon.go.tpl",
		Addon:    EventControllerAddon,
		required: true,
	},
	{
		Path:     "/internal/stream/state.go",
		Template: "stream_event_state_addon.go.tpl",
		Addon:    EventStateAddon,
	},
	{
		Path:     "/pkg/{events_package}/events.go",
		Template: "pkg_events_events_addon.go.tpl",
		Addon:    EventsEventsAddon,
		required: true,
	},
	{
		Path:     "/pkg/{events_package}/events_encoding.go",
		Template: "pkg_events_events_encoding_addon.go.tpl",
		Addon:    EventsEventsEncodingAddon,
		required: true,
	},
}

func AddonFiles(path string, m *Manifest, required bool) []File {
	result := make([]File, 0, len(commandMutationAddons)+len(eventMutationAddons))
	for _, file := range commandMutationAddons {
		if required && !file.required {
			continue
		}
		file.Path = NormalizePath(file.Path, m)
		file.Path = filepath.Join(path, file.Path)
		result = append(result, file)
	}
	for _, file := range eventMutationAddons {
		if required && !file.required {
			continue
		}
		file.Path = NormalizePath(file.Path, m)
		file.Path = filepath.Join(path, file.Path)
		result = append(result, file)
	}
	return result
}
