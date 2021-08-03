package schema

import (
	"strings"
)

func SanitizeManifest(m *Manifest) {
	m.CommandsPkgName = sanitizePackageName(m.CommandsPkgName)
	m.GoModules = sanitizePackageName(m.GoModules)
	m.StreamPkgName = sanitizePackageName(m.StreamPkgName)
	m.EventsPkgName = sanitizePackageName(m.EventsPkgName)
	m.PackageName = sanitizePackageName(m.PackageName)
	m.StreamName = sanitizeStreamName(m.StreamName)
	if len(m.StreamPkgName) == 0 {
		m.StreamPkgName = m.PackageName + "stream"
	}
	if len(m.EventsPkgName) == 0 {
		m.EventsPkgName = m.PackageName + "events"
	}
	if len(m.CommandsPkgName) == 0 {
		m.CommandsPkgName = m.PackageName + "commands"
	}
	sanitizeCommands(m.Mutations.Commands)
	sanitizeEvents(m.Mutations.Events)
}

func sanitizePackageName(name string) string {
	name = strings.ToLower(name)
	return strings.ReplaceAll(name, " ", "")
}

func sanitizeStreamName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "")
	return strings.Title(name)
}

func sanitizeName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	return strings.Title(name)
}

func trim(name string) string {
	return strings.ReplaceAll(name, " ", "")
}

func sanitizeCommands(commands []CommandMutation) {
	for i, cmd := range commands {
		if cmd.Delete == YesOp && cmd.Create == YesOp {
			commands[i].Create = NoOp
			commands[i].Delete = NoOp
		}
		cmd.Mutation = sanitizeName(cmd.Mutation)
		cmd.Command.Name = sanitizeName(cmd.Command.Name)
		cmd.Command.Payload = sanitizeName(cmd.Command.Payload)
		cmd.Event.Name = sanitizeName(cmd.Event.Name)
		cmd.Event.Payload = sanitizeName(cmd.Event.Payload)
		if cmd.Command.Name == cmd.Command.Payload {
			cmd.Command.Payload = cmd.Command.Payload + "Payload"
		}
		if cmd.Event.Name == cmd.Event.Payload {
			cmd.Event.Payload = cmd.Event.Payload + "Payload"
		}
		commands[i] = cmd
	}
}

func sanitizeEvents(events []EventMutation) {
	for i, e := range events {
		if e.Delete == YesOp && e.Create == YesOp {
			events[i].Create = NoOp
			events[i].Delete = NoOp
		}
		e.Mutation = sanitizeName(e.Mutation)
		e.InEvent.Name = trim(e.InEvent.Name)
		e.InEvent.Payload = trim(e.InEvent.Payload)
		e.OutEvent.Name = sanitizeName(e.OutEvent.Name)
		e.OutEvent.Payload = sanitizeName(e.OutEvent.Payload)
		if e.OutEvent.Name == e.OutEvent.Payload {
			e.OutEvent.Payload = e.OutEvent.Payload + "Payload"
		}
		events[i] = e
	}
}
