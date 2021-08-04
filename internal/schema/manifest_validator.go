package schema

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

var (
	packageNameRe = regexp.MustCompile("^[a-zA-Z0-9]*$")
	streamNameRe  = regexp.MustCompile("^[a-zA-Z]*$")
)

func ValidateManifest(m *Manifest) error {
	for _, fn := range []func(m *Manifest) error{
		validatePackageName,
		validateGoModules,
		validateGoVersion,
		validateStreamName,
		validateProjectName,
		validateCommands,
		validateEvents,
		validatePublisherAdapter,
		validateStorageAdapter,
	} {
		if err := fn(m); err != nil {
			return err
		}
	}
	return validator.New().Struct(m)
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
	if len(m.GoVersion) < 0 {
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
	case KafkaStreamPublisherAdapter,
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

func validateStorageAdapter(m *Manifest) error {
	switch m.StreamStorage.AdapterID {
	case RedisStreamStorageAdapter,
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

func validateCommands(m *Manifest) error {
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
	}
	return nil
}

func validateEvents(m *Manifest) error {
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
	}
	return nil
}
