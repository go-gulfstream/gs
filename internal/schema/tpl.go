package schema

import (
	"bytes"

	"github.com/go-gulfstream/gs/internal/templates"
)

func renderTemplate(fileTpl string, m *Manifest) ([]byte, error) {
	tpl, err := templates.Parse(fileTpl)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tpl.Execute(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func renderCommandMutationAddon(fileTpl string, manifest *Manifest, m CommandMutation) ([]byte, error) {
	tpl, err := templates.Parse(fileTpl)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tpl.Execute(buf, struct {
		CommandMutation
		GoModules       string
		EventsPkgName   string
		CommandsPkgName string
		StreamPkgName   string
	}{
		CommandMutation: m,
		GoModules:       manifest.GoModules,
		EventsPkgName:   manifest.EventsPkgName,
		CommandsPkgName: manifest.CommandsPkgName,
		StreamPkgName:   manifest.StreamPkgName,
	}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type pkg struct {
	Path string
}

func renderEventMutationAddon(fileTpl string, manifest *Manifest, m EventMutation) ([]byte, error) {
	tpl, err := templates.Parse(fileTpl)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	var importEvents []pkg
	if len(manifest.ImportEvents) > 0 {
		importEvents = make([]pkg, len(manifest.ImportEvents))
		for i, ep := range manifest.ImportEvents {
			importEvents[i] = pkg{
				Path: ep,
			}
		}
	}

	if err := tpl.Execute(buf, struct {
		EventMutation
		GoModules       string
		EventsPkgName   string
		CommandsPkgName string
		StreamPkgName   string
		ImportEvents    []pkg
	}{
		EventMutation:   m,
		GoModules:       manifest.GoModules,
		EventsPkgName:   manifest.EventsPkgName,
		CommandsPkgName: manifest.CommandsPkgName,
		StreamPkgName:   manifest.StreamPkgName,
		ImportEvents:    importEvents,
	}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
