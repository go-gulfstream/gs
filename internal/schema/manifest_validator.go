package schema

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

var (
	packageNameRe = regexp.MustCompile("^[a-zA-Z0-9]*$")
	streamNameRe  = regexp.MustCompile("^[a-zA-Z]*$")
)

func ValidateManifest(m *Manifest) error {
	for _, fn := range []func(m *Manifest) error{
		validatePackageName,
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
	case kafkaStreamPublisherAdapter,
		connectorStreamPublisherAdapter:
		return nil
	default:
		return fmt.Errorf("invalid publisher adapter id. got %s, expected %v",
			m.StreamPublisher.AdapterID, PublisherAdapters)
	}
}

func validateStorageAdapter(m *Manifest) error {
	switch m.StreamStorage.AdapterID {
	case redisStreamStorageAdapter,
		postgresStreamStorageAdapter:
		return nil
	default:
		return fmt.Errorf("invalid stream storage adapter id. got %s, expected %v",
			m.StreamStorage.AdapterID, StorageAdapters)
	}
}

func validateCommands(m *Manifest) error {
	//if len(m.Mutations.Commands) == 0 {
	//	return nil
	//}
	//for _, e := range m.Mutations.Events {
	//	if err := e.Validate(); err != nil {
	//		return err
	//	}
	//}
	return nil
}

func validateEvents(m *Manifest) error {
	//if len(m.Mutations.Events) == 0 {
	//	return nil
	//}
	//for _, c := range m.Mutations.Commands {
	//	if err := c.Validate(); err != nil {
	//		return err
	//	}
	//}
	return nil
}
