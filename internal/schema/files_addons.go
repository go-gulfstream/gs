package schema

import "path/filepath"

const (
	// Events
	EventMutationAddon        = "EventMutation"
	EventMutationTestAddon    = "EventMutationTest"
	EventMutationImplAddon    = "EventMutationImpl"
	EventControllerAddon      = "EventController"
	EventStateAddon           = "EventState"
	EventsEventsAddon         = "EventsEventsAddon"
	EventsEventsEncodingAddon = "EventsEventsEncodingAddon"

	// Commands
	CommandMutationAddon        = "CommandMutation"
	CommandMutationImplAddon    = "CommandMutationImpl"
	CommandMutationTestAddon    = "CommandMutationTest"
	CommandControllerAddon      = "CommandController"
	CommandStateAddon           = "CommandState"
	CommandsAddon               = "Commands"
	CommandsEncodingAddon       = "CommandsEncoding"
	CommandsEventsAddon         = "CommandEvents"
	CommandsEventsEncodingAddon = "CommandEventsEncoding"

	// Projection
	ProjectionTestAddon                = "ProjectionTest"
	CommandMutationProjectionAddon     = "CommandMutationProjection"
	CommandMutationImplProjectionAddon = "CommandMutationImplProjection"
	CommandControllerProjectionAddon   = "CommandControllerProjection"
	EventMutationProjectionAddon       = "EventMutationProjection"
	EventMutationImplProjectionAddon   = "EventMutationImplProjection"
	EventControllerProjectionAddon     = "EventControllerProjection"
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
	{
		Path:     "/internal/projection/projection.go",
		Template: "projection_projection_commands_addon.go.tpl",
		Addon:    CommandMutationProjectionAddon,
	},
	{
		Path:     "/internal/projection/projection_test.go",
		Template: "projection_projection_test_addon.go.tpl",
		Addon:    ProjectionTestAddon,
	},
	{
		Path:     "/internal/projection/projection.go",
		Template: "projection_projection_impl_commands_addon.go.tpl",
		Addon:    CommandMutationImplProjectionAddon,
	},
	{
		Path:     "/internal/projection/controller.go",
		Template: "projection_controller_commands_addon.go.tpl",
		Addon:    CommandControllerProjectionAddon,
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
		Path:     "/internal/projection/projection.go",
		Template: "projection_projection_events_addon.go.tpl",
		Addon:    EventMutationProjectionAddon,
	},
	{
		Path:     "/internal/projection/projection.go",
		Template: "projection_projection_impl_events_addon.go.tpl",
		Addon:    EventMutationImplProjectionAddon,
	},
	{
		Path:     "/internal/projection/projection_test.go",
		Template: "projection_projection_test_addon.go.tpl",
		Addon:    ProjectionTestAddon,
	},
	{
		Path:     "/internal/projection/controller.go",
		Template: "projection_controller_events_addon.go.tpl",
		Addon:    EventControllerProjectionAddon,
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
