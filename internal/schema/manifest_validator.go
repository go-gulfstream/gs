package schema

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	packageNameRe = regexp.MustCompile("^[a-zA-Z0-9]*$")
	streamNameRe  = regexp.MustCompile("^[a-zA-Z]*$")
	urlRe         = regexp.MustCompile(`^(?:[^%]|%[0-9A-Fa-f]{2})*$`)
)

var uidx = newUnique()

func CheckUnique(s string) bool {
	return uidx.has(s)
}

func IndexCommandMutation(c CommandMutation) {
	uidx.addCommandMutation(c)
}

func IndexEventMutation(c EventMutation) {
	uidx.addEventMutation(c)
}

func Index(m *Manifest) {
	for _, c := range m.Mutations.Commands {
		uidx.addCommandMutation(c)
	}
	for _, e := range m.Mutations.Events {
		uidx.addEventMutation(e)
	}
}

func ValidateManifest(m *Manifest) error {
	idx := newUnique()
	for _, fn := range []func(m *Manifest) error{
		validatePackageName,
		validateGoModules,
		validateGoVersion,
		validateStreamName,
		validateProjectName,
		validatePublisherAdapter,
		validateStorageAdapter,
	} {
		if err := fn(m); err != nil {
			return err
		}
	}
	if err := validateGoGetPackages(m, idx); err != nil {
		return err
	}
	if err := validatePkgs(m, idx); err != nil {
		return err
	}
	if err := validateCommands(m, idx); err != nil {
		return err
	}
	if err := validateEvents(m, idx); err != nil {
		return err
	}
	return nil
}

func validateGoModules(m *Manifest) error {
	if len(m.GoModules) < 2 {
		return fmt.Errorf("gulfstream.yml: go modules too short. got %d, expected > 2 symbols",
			len(m.GoModules))
	}
	if len(m.PackageName) > 128 {
		return fmt.Errorf("gulfstream.yml: go modules too long. got %d, expected <= 128 symbols",
			len(m.GoModules))
	}
	return nil
}

func validateGoVersion(m *Manifest) error {
	if len(m.GoVersion) == 0 {
		return fmt.Errorf("gulfstream.yml: undefined go version")
	}
	return nil
}

func validatePackageName(m *Manifest) error {
	if len(m.PackageName) < 2 {
		return fmt.Errorf("gulfstream.yml: package name too short. got %d, expected > 2 symbols",
			len(m.PackageName))
	}
	if !packageNameRe.MatchString(m.PackageName) {
		return fmt.Errorf("gulfstream.yml: invalid package name. got %s, expected only alphanumeric project name without space",
			m.PackageName)
	}
	if len(m.PackageName) > 36 {
		return fmt.Errorf("gulfstream.yml: package name too long. got %d, expected <= 36 symbols",
			len(m.PackageName))
	}
	return nil
}

func validateStreamName(m *Manifest) error {
	if len(m.StreamName) < 2 {
		return fmt.Errorf("gulfstream.yml: stream name too short. got %d, expected > 2 symbols",
			len(m.StreamName))
	}
	if !streamNameRe.MatchString(m.StreamName) {
		return fmt.Errorf("gulfstream.yml: invalid stream name. got %s, expected only alphabet stream name without space",
			m.StreamName)
	}
	if len(m.StreamName) > 36 {
		return fmt.Errorf("gulfstream.yml: stream name too long. got %d, expected <= 36 symbols",
			len(m.StreamName))
	}
	return nil
}

func validateProjectName(m *Manifest) error {
	if len(m.Name) < 2 {
		return fmt.Errorf("gulfstream.yml: project name too short. got %d, expected > 2 symbols",
			len(m.Name))
	}
	return nil
}

func validatePublisherAdapter(m *Manifest) error {
	switch m.StreamPublisher.AdapterID {
	case DefaultStreamPublisher,
		KafkaStreamPublisherAdapter,
		ConnectorStreamPublisherAdapter:
		return nil
	default:
		return fmt.Errorf("gulfstream.yml: invalid publisher adapter id. got %s, expected %v",
			m.StreamPublisher.AdapterID,
			strings.Join([]string{
				KafkaStreamPublisherAdapter.String(),
				ConnectorStreamPublisherAdapter.String(),
			}, " OR "))
	}
}

func validateGoGetPackages(m *Manifest, u *unique) error {
	if len(m.GoGetPackages) == 0 {
		return nil
	}
	for _, pkg := range m.GoGetPackages {
		if !urlRe.MatchString(pkg) {
			return fmt.Errorf("gulfstream.yml: invalid package url %s", pkg)
		}
		if u.has(pkg) {
			return fmt.Errorf("gulfstream.yml: package url %s already exists", pkg)
		}
		u.add(pkg)
	}
	return nil
}

func validatePkgs(m *Manifest, u *unique) error {
	if ok := u.has(m.StreamPkgName); ok {
		return fmt.Errorf("gulfstream.yml: go_stream_pkg_name: %s already exists in go_commands_pkg_name or go_events_pkg_name",
			m.StreamPkgName)
	}
	u.add(m.StreamPkgName)
	if ok := u.has(m.CommandsPkgName); ok {
		return fmt.Errorf("gulfstream.yml: go_commands_pkg_name: %s already exists in go_stream_pkg_name or go_events_pkg_name",
			m.CommandsPkgName)
	}
	u.add(m.CommandsPkgName)
	if ok := u.has(m.EventsPkgName); ok {
		return fmt.Errorf("gulfstream.yml: go_events_pkg_name: %s already exists in go_commands_pkg_name or go_stream_pkg_name",
			m.EventsPkgName)
	}
	u.add(m.EventsPkgName)
	return nil
}

func validateStorageAdapter(m *Manifest) error {
	switch m.StreamStorage.AdapterID {
	case DefaultStreamStorage,
		RedisStreamStorageAdapter,
		PostgresStreamStorageAdapter:
		return nil
	default:
		return fmt.Errorf("gulfstream.yml: invalid stream storage adapter id. got %s, expected %v",
			m.StreamStorage.AdapterID, strings.Join([]string{
				RedisStreamStorageAdapter.String(),
				PostgresStreamStorageAdapter.String(),
			}, " OR "))
	}
}

func validateCommands(m *Manifest, idx *unique) error {
	if len(m.Mutations.Commands) == 0 {
		return nil
	}
	for i, cmd := range m.Mutations.Commands {
		if len(cmd.Mutation) < 2 {
			return fmt.Errorf("gulfstream.yml: command mutation too short. index[%d]", i)
		}
		if len(cmd.Command.Name) < 2 {
			return fmt.Errorf("gulfstream.yml: Mutations.%s{InCommand.###EMPTY###} => command name too short. index[%d]",
				cmd.Mutation, i)
		}
		if len(cmd.Event.Name) < 2 {
			return fmt.Errorf("gulfstream.yml: Mutations.%s{InCommand.%s, OutEvent.###EMPTY###} => event name too short. index[%d]",
				cmd.Mutation, cmd.Command.Name, i)
		}
		if ok := idx.hasCommandMutation(cmd); ok {
			return fmt.Errorf("gulfstream.yml: there are duplicates in command mutations index[%d]", i)
		}
		idx.addCommandMutation(cmd)
	}
	return nil
}

func validateEvents(m *Manifest, idx *unique) error {
	if len(m.Mutations.Events) == 0 {
		return nil
	}
	for i, e := range m.Mutations.Events {
		if len(e.Mutation) < 2 {
			return fmt.Errorf("gulfstream.yml: event mutation too short. index[%d]", i)
		}
		if len(e.InEvent.Name) < 2 {
			return fmt.Errorf("gulfstream.yml: Mutations.%s{InEvent.###EMPTY###} => in event name too short. index[%d]",
				e.Mutation, i)
		}
		if len(e.OutEvent.Name) < 2 {
			return fmt.Errorf("gulfstream.yml: Mutations.%s{InEvent.%s, OutEvent.###EMPTY} => out event name too short. index[%d]",
				e.Mutation, e.InEvent.Name, i)
		}
		if ok := idx.hasEventMutation(e); ok {
			return fmt.Errorf("gulfstream.yml: there are duplicates in event mutations index[%d]", i)
		}
		idx.addEventMutation(e)
	}
	return nil
}

type unique struct {
	mutations map[string]struct{}
}

func newUnique() *unique {
	return &unique{
		mutations: make(map[string]struct{}),
	}
}

func (u *unique) add(s string) {
	u.mutations[s] = struct{}{}
}

func (u *unique) has(val string) bool {
	if _, ok := u.mutations[val]; ok {
		return true
	}
	return false
}

func (u *unique) addCommandMutation(c CommandMutation) {
	u.mutations[c.Mutation] = struct{}{}
	u.mutations[c.Command.Name] = struct{}{}
	if len(c.Command.Payload) > 0 {
		u.mutations[c.Command.Payload] = struct{}{}
	}
	u.mutations[c.Event.Name] = struct{}{}
	if len(c.Event.Payload) > 0 {
		u.mutations[c.Event.Payload] = struct{}{}
	}
}

func (u *unique) hasCommandMutation(c CommandMutation) bool {
	if _, ok := u.mutations[c.Mutation]; ok {
		return true
	}
	if _, ok := u.mutations[c.Command.Name]; ok {
		return true
	}
	if len(c.Command.Payload) > 0 {
		if _, ok := u.mutations[c.Command.Payload]; ok {
			return true
		}
	}
	if _, ok := u.mutations[c.Event.Name]; ok {
		return true
	}
	if len(c.Event.Payload) > 0 {
		if _, ok := u.mutations[c.Event.Payload]; ok {
			return true
		}
	}
	return false
}

func (u *unique) addEventMutation(e EventMutation) {
	u.mutations[e.Mutation] = struct{}{}
	u.mutations[e.InEvent.Name] = struct{}{}
	if len(e.InEvent.Payload) > 0 {
		u.mutations[e.InEvent.Payload] = struct{}{}
	}
	u.mutations[e.OutEvent.Name] = struct{}{}
	if len(e.InEvent.Payload) > 0 {
		u.mutations[e.OutEvent.Payload] = struct{}{}
	}
}

func (u *unique) hasEventMutation(e EventMutation) bool {
	if _, ok := u.mutations[e.Mutation]; ok {
		return true
	}
	if _, ok := u.mutations[e.InEvent.Name]; ok {
		return true
	}
	if len(e.InEvent.Payload) > 0 {
		if _, ok := u.mutations[e.InEvent.Payload]; ok {
			return true
		}
	}
	if _, ok := u.mutations[e.OutEvent.Name]; ok {
		return true
	}
	if len(e.InEvent.Payload) > 0 {
		if _, ok := u.mutations[e.OutEvent.Payload]; ok {
			return true
		}
	}
	return false
}
