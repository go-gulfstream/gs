package schema

import (
	"path/filepath"
	"strings"
)

func Walk(path string, m *Manifest, fn func(File) error) error {
	if m == nil {
		return nil
	}
	for _, file := range files {
		file.Path = NormalizePath(file.Path, m)
		file.Path = filepath.Join(path, file.Path)
		if len(file.Template) > 0 {
			file.HasTemplate = true
			tplData, err := renderTemplate(file.Template, m)
			if err != nil {
				return err
			}
			file.TemplateData = tplData
		}
		if err := fn(file); err != nil {
			return err
		}
	}
	return nil
}

func WalkCommandMutationAddons(path string, m *Manifest, fn func(CommandMutation, File) error) error {
	if m == nil {
		return nil
	}
	for _, mutation := range m.Mutations.Commands {
		for _, addon := range commandMutationAddons {
			addon.Path = NormalizePath(addon.Path, m)
			addon.Path = filepath.Join(path, addon.Path)
			addon.HasTemplate = true
			tplData, err := renderCommandMutationAddon(addon.Template, m, mutation)
			if err != nil {
				return err
			}
			addon.TemplateData = tplData
			if err := fn(mutation, addon); err != nil {
				return err
			}
		}
	}
	return nil
}

func WalkEventMutationAddons(path string, m *Manifest, fn func(EventMutation, File) error) error {
	if m == nil {
		return nil
	}
	for _, mutation := range m.Mutations.Events {
		for _, addon := range eventMutationAddons {
			addon.Path = NormalizePath(addon.Path, m)
			addon.Path = filepath.Join(path, addon.Path)
			addon.HasTemplate = true
			tplData, err := renderEventMutationAddon(addon.Template, m, mutation)
			if err != nil {
				return err
			}
			addon.TemplateData = tplData
			if err := fn(mutation, addon); err != nil {
				return err
			}
		}
	}
	return nil
}

func NormalizePath(filePath string, m *Manifest) string {
	filePath = strings.ReplaceAll(filePath, "{package}", m.PackageName)
	filePath = strings.ReplaceAll(filePath, "{commands_package}", m.CommandsPkgName)
	filePath = strings.ReplaceAll(filePath, "{events_package}", m.EventsPkgName)
	filePath = strings.ReplaceAll(filePath, "{stream_package}", m.StreamPkgName)
	return filePath
}
