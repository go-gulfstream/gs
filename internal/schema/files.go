package schema

import (
	"fmt"
	"strings"
)

type File struct {
	IsDir        bool
	Path         string
	Template     string
	TemplateData []byte
	HasTemplate  bool
}

func (f File) String() string {
	return fmt.Sprintf("File{Path:%s, IsDir:%v}", f.Path, f.IsDir)
}

func (f File) IsGo() bool {
	return strings.HasSuffix(f.Path, ".go")
}

var files = []File{
	{
		Path:  "/cmd",
		IsDir: true,
	},
	{
		Path:  "/cmd/{package}",
		IsDir: true,
	},
	{
		Path:  "/cmd/{package}-projection",
		IsDir: true,
	},
	{
		Path:     "/cmd/{package}/main.go",
		Template: "cmd_main.go.tpl",
	},
	{
		Path:     "/cmd/{package}-projection/main.go",
		Template: "cmd_proj_main.go.tpl",
	},
	{
		Path:  "/internal",
		IsDir: true,
	},
	{
		Path:  "/internal/stream",
		IsDir: true,
	},
	{
		Path:     "/internal/stream/command_mutation.go",
		Template: "stream_command_mutation.go.tpl",
	},
	{
		Path:     "/internal/stream/command_controller.go",
		Template: "stream_command_controller.go.tpl",
	},
	{
		Path:     "/internal/stream/event_mutation.go",
		Template: "stream_event_mutation.go.tpl",
	},
	{
		Path:     "/internal/stream/event_controller.go",
		Template: "stream_event_controller.go.tpl",
	},
	{
		Path:     "/internal/stream/state.go",
		Template: "stream_state.go.tpl",
	},
	{
		Path:     "/internal/stream/state_encoding.go",
		Template: "stream_state_encoding.go.tpl",
	},
	{
		Path:  "/internal/projection",
		IsDir: true,
	},
	{
		Path:     "/internal/projection/projection.go",
		Template: "projection_projection.go.tpl",
	},
	{
		Path:     "/internal/projection/controller.go",
		Template: "projection_controller.go.tpl",
	},
	{
		Path:     "/internal/projection/storage.go",
		Template: "projection_storage.go.tpl",
	},
	{
		Path:     "/internal/projection/types.go",
		Template: "projection_types.go.tpl",
	},
	{
		Path:  "/internal/config",
		IsDir: true,
	},
	{
		Path:     "/internal/config/config.go",
		Template: "config_config.go.tpl",
	},
	{
		Path:  "/internal/api",
		IsDir: true,
	},
	{
		Path:  "/docs",
		IsDir: true,
	},
	{
		Path:     "/docs/README.md",
		Template: "docs_readme.tpl",
	},
	{
		Path:  "/pkg",
		IsDir: true,
	},
	{
		Path:  "/pkg/{commands_package}",
		IsDir: true,
	},
	{
		Path:     "/pkg/{commands_package}/commands.go",
		Template: "pkg_commands.go.tpl",
	},
	{
		Path:     "/pkg/{commands_package}/commands_encoding.go",
		Template: "pkg_commands_encoding.go.tpl",
	},
	{
		Path:  "/pkg/{events_package}",
		IsDir: true,
	},
	{
		Path:     "/pkg/{events_package}/events.go",
		Template: "pkg_events.go.tpl",
	},
	{
		Path:     "/pkg/{events_package}/events_encoding.go",
		Template: "pkg_events_encoding.go.tpl",
	},
	{
		Path:  "/pkg/{stream_package}",
		IsDir: true,
	},
	{
		Path:     "/pkg/{stream_package}/stream.go",
		Template: "pkg_stream.go.tpl",
	},
	{
		Path:  "/docker",
		IsDir: true,
	},
	{
		Path:     "/docker/stream.dockerfile",
		Template: "docker_stream.go.tpl",
	},
	{
		Path:     "/docker/projection.dockerfile",
		Template: "docker_projection.go.tpl",
	},
	{
		Path:     "/.gitignore",
		Template: "gitignore.tpl",
	},
	{
		Path:     "/Makefile",
		Template: "makefile.tpl",
	},
	{
		Path:     "/README.md",
		Template: "readme.tpl",
	},
	{
		Path:     "/go.mod",
		Template: "go-mod.tpl",
	},
}
