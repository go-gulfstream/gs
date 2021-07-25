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
		return fmt.Errorf("go modules too short. got %d, expected > 2 symbols",
			len(m.GoModules))
	}
	if len(m.PackageName) > 128 {
		return fmt.Errorf("go modules too long. got %d, expected <= 128 symbols",
			len(m.GoModules))
	}
	return nil
}

func validateGoVersion(m *Manifest) error {
	if len(m.GoVersion) < 0 {
		return fmt.Errorf("undefined go version")
	}
	return nil
}

func validatePackageName(m *Manifest) error {
	if len(m.PackageName) < 2 {
		return fmt.Errorf("package name too short. got %d, expected > 2 symbols",
			len(m.PackageName))
	}
	if !packageNameRe.MatchString(m.PackageName) {
		return fmt.Errorf("invalid package name. got %s, expected only alphanumeric project name without space",
			m.PackageName)
	}
	if len(m.PackageName) > 36 {
		return fmt.Errorf("package name too long. got %d, expected <= 36 symbols",
			len(m.PackageName))
	}
	return nil
}

func validateStreamName(m *Manifest) error {
	if len(m.StreamName) < 2 {
		return fmt.Errorf("stream name too short. got %d, expected > 2 symbols",
			len(m.StreamName))
	}
	if !streamNameRe.MatchString(m.StreamName) {
		return fmt.Errorf("invalid stream name. got %s, expected only alphabet stream name without space",
			m.StreamName)
	}
	if len(m.StreamName) > 36 {
		return fmt.Errorf("stream name too long. got %d, expected <= 36 symbols",
			len(m.StreamName))
	}
	return nil
}

func validateProjectName(m *Manifest) error {
	if len(m.Name) < 2 {
		return fmt.Errorf("project name too short. got %d, expected > 2 symbols",
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
		return fmt.Errorf("invalid publisher adapter id. got %s, expected %v",
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
		return fmt.Errorf("invalid stream storage adapter id. got %s, expected %v",
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
			return fmt.Errorf("mutation name too short. index %d", i)
		}
		if len(cmd.Command.Name) < 2 {
			return fmt.Errorf("command name too short. index %d", i)
		}
		if len(cmd.Event.Name) < 2 {
			return fmt.Errorf("event name too short. index %d", i)
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
			return fmt.Errorf("mutation name too short. index %d", i)
		}
		if len(e.InEvent.Name) < 2 {
			return fmt.Errorf("inEvent name too short. index %d", i)
		}
		if len(e.OutEvent.Name) < 2 {
			return fmt.Errorf("outEvent name too short. index %d", i)
		}
	}
	return nil
}
