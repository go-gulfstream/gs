package schema

import "strings"

func SanitizeManifest(m *Manifest) {
	m.CommandsPkgName = sanitizePackageName(m.CommandsPkgName)
	m.StreamPkgName = sanitizePackageName(m.StreamPkgName)
	m.EventsPkgName = sanitizePackageName(m.EventsPkgName)
	m.PackageName = sanitizePackageName(m.PackageName)
	m.StreamName = sanitizeStreamName(m.StreamName)
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

//func sanitizeCommands(m *Manifest) {
//	index := make(map[string]struct{})
//	commands := make([]CommandMutation, 0)
//	for _, mc := range m.Mutations.Commands {
//		if len(mc.Mutation) == 0 || !mc.Command.Validate() {
//			continue
//		}
//
//		_, foundCommand := index[mc.Mutation]
//		if foundCommand {
//			continue
//		}
//
//		index[mc.Mutation] = struct{}{}
//
//		mc.Event.Name = strings.Title(mc.Event.Name)
//		mc.Event.Payload = strings.Title(mc.Event.Payload)
//		mc.Command.Name = strings.Title(mc.Command.Name)
//		mc.Command.Payload = strings.Title(mc.Command.Payload)
//		mc.Mutation = strings.Title(mc.Mutation)
//
//		commands = append(commands, mc)
//	}
//	m.Mutations.Commands = commands
//}
//
//func sanitizeEvents(m *Manifest) {
//	index := make(map[string]struct{})
//	events := make([]EventMutation, 0)
//	for _, ec := range m.Mutations.Events {
//		if len(ec.Mutation) == 0 || ec.Event.Validate() {
//			continue
//		}
//		_, foundEvent := index[ec.Mutation]
//		if foundEvent {
//			continue
//		}
//
//		index[ec.Mutation] = struct{}{}
//
//		ec.Mutation = strings.Title(ec.Mutation)
//		ec.Event.Name = strings.Title(ec.Event.Name)
//		ec.Event.Payload = strings.Title(ec.Event.Payload)
//		events = append(events, ec)
//	}
//	m.Mutations.Events = events
//}
