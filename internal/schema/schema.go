package schema

import "fmt"

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

var files = []File{
	{
		Path:  "/cmd",
		IsDir: true,
	},
	{
		Path:  "/cmd/%s",
		IsDir: true,
	},
	{
		Path:  "/cmd/%s-projection",
		IsDir: true,
	},
	{
		Path:     "/cmd/%s/main.go",
		Template: "cmd_main.go.tpl",
	},
	{
		Path:     "/cmd/%s-projection/main.go",
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
		Path:     "/internal/stream/mutation.go",
		Template: "stream_mutation.go.tpl",
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
		Path:     "/internal/stream/controller.go",
		Template: "stream_controller.go.tpl",
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
		Path:  "/pkg",
		IsDir: true,
	},
	{
		Path:  "/pkg/commands",
		IsDir: true,
	},
	{
		Path:  "/pkg/events",
		IsDir: true,
	},
	{
		Path:  "/pkg/stream",
		IsDir: true,
	},
	{
		Path:     "/pkg/stream/%s",
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
