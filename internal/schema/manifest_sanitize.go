package schema

import (
	"strings"
)

const tplSymbol = "$"

func SanitizeManifest(m *Manifest) {
	m.CommandsPkgName = SanitizePackageName(m.CommandsPkgName)
	m.GoModules = SanitizePackageName(m.GoModules)
	m.StreamPkgName = SanitizePackageName(m.StreamPkgName)
	m.EventsPkgName = SanitizePackageName(m.EventsPkgName)
	m.PackageName = SanitizePackageName(m.PackageName)
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

	m.Mutations.Commands = filterCommandMutations(m.Mutations.Commands)
	m.Mutations.Events = filterEventMutations(m.Mutations.Events)

	sanitizeCommands(m.Mutations.Commands)
	sanitizeEvents(m.Mutations.Events)
}

func SanitizePackageName(name string) string {
	name = strings.ToLower(name)
	return strings.ReplaceAll(name, " ", "")
}

func sanitizeStreamName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "")
	return strings.Title(name)
}

func SanitizeName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	return strings.Title(name)
}

func trim(name string) string {
	return strings.ReplaceAll(name, " ", "")
}

func filterCommandMutations(commands []CommandMutation) []CommandMutation {
	result := make([]CommandMutation, 0, len(commands))
	for _, m := range commands {
		if strings.HasPrefix(m.Mutation, tplSymbol) {
			continue
		}
		if strings.HasPrefix(m.Command.Name, tplSymbol) {
			continue
		}
		if strings.HasPrefix(m.Event.Name, tplSymbol) {
			continue
		}
		result = append(result, m)
	}
	return result
}

func filterEventMutations(events []EventMutation) []EventMutation {
	result := make([]EventMutation, 0, len(events))
	for _, m := range events {
		if strings.HasPrefix(m.Mutation, tplSymbol) {
			continue
		}
		if strings.HasPrefix(m.InEvent.Name, tplSymbol) {
			continue
		}
		if strings.HasPrefix(m.OutEvent.Name, tplSymbol) {
			continue
		}
		result = append(result, m)
	}
	return result
}

func sanitizeCommands(commands []CommandMutation) {
	for i, cmd := range commands {
		if cmd.Delete == YesOp && cmd.Create == YesOp {
			commands[i].Create = NoOp
			commands[i].Delete = NoOp
		}
		cmd.Mutation = SanitizeName(cmd.Mutation)
		cmd.Command.Name = SanitizeName(cmd.Command.Name)
		cmd.Command.Payload = SanitizeName(cmd.Command.Payload)
		cmd.Event.Name = SanitizeName(cmd.Event.Name)
		cmd.Event.Payload = SanitizeName(cmd.Event.Payload)
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
		// template data
		if strings.HasPrefix(e.Mutation, "_") {
			continue
		}
		if e.Delete == YesOp && e.Create == YesOp {
			events[i].Create = NoOp
			events[i].Delete = NoOp
		}
		e.Mutation = SanitizeName(e.Mutation)
		e.InEvent.Name = trim(e.InEvent.Name)
		e.InEvent.Payload = trim(e.InEvent.Payload)
		e.OutEvent.Name = SanitizeName(e.OutEvent.Name)
		e.OutEvent.Payload = SanitizeName(e.OutEvent.Payload)
		if e.OutEvent.Name == e.OutEvent.Payload {
			e.OutEvent.Payload = e.OutEvent.Payload + "Payload"
		}
		events[i] = e
	}
}
