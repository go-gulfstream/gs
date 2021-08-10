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
	Addon        string
	required     bool
}

func (f File) String() string {
	return fmt.Sprintf("File{Path:%s, IsDir:%v}", f.Path, f.IsDir)
}

func (f File) IsGo() bool {
	return strings.HasSuffix(f.Path, ".go")
}

var files = []File{
	{
		Path:     "/cmd",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/cmd/{package}",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/cmd/{package}-projection",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/cmd/{package}/main.go",
		Template: "cmd_main.go.tpl",
		required: true,
	},
	{
		Path:     "/cmd/{package}-projection/main.go",
		Template: "cmd_proj_main.go.tpl",
		required: true,
	},
	{
		Path:     "/internal",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/internal/stream",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/internal/stream/command_mutation.go",
		Template: "stream_command_mutation.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/stream/command_mutation_test.go",
		Template: "stream_command_mutation_test.go.tpl",
	},
	{
		Path:     "/internal/stream/command_controller.go",
		Template: "stream_command_controller.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/stream/event_mutation.go",
		Template: "stream_event_mutation.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/stream/event_mutation_test.go",
		Template: "stream_event_mutation_test.go.tpl",
	},
	{
		Path:     "/internal/stream/event_controller.go",
		Template: "stream_event_controller.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/stream/state.go",
		Template: "stream_state.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/stream/state_encoding.go",
		Template: "stream_state_encoding.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/projection",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/internal/projection/projection.go",
		Template: "projection_projection.go.tpl",
		required: true,
	},
	{
		Path:     "/internal/projection/controller.go",
		Template: "projection_controller.go.tpl",
		required: true,
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
		Path:     "/internal/projection/projection_test.go",
		Template: "projection_projection_test.go.tpl",
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
		Path:     "/internal/config/stream.go",
		Template: "config_stream.go.tpl",
	},
	{
		Path:     "/internal/config/projection.go",
		Template: "config_projection.go.tpl",
	},
	{
		Path:  "/internal/api",
		IsDir: true,
	},
	{
		Path:     "/internal/api/service.go",
		Template: "api_service.go.tpl",
	},
	{
		Path:     "/internal/api/service_test.go",
		Template: "api_service_test.go.tpl",
	},
	{
		Path:     "/internal/api/endpoints.go",
		Template: "api_endpoints.go.tpl",
	},
	{
		Path:     "/internal/api/middleware.go",
		Template: "api_middleware.go.tpl",
	},
	{
		Path:     "/internal/api/http.go",
		Template: "api_http.go.tpl",
	},
	{
		Path:     "/internal/api/grpc.go",
		Template: "api_grpc.go.tpl",
	},
	{
		Path:  "/docs",
		IsDir: true,
	},
	{
		Path:  "/config",
		IsDir: true,
	},
	{
		Path:     "/config/stream.config.yml",
		Template: "config_stream.config.yml.tpl",
	},
	{
		Path:     "/config/projection.config.yml",
		Template: "config_projection.config.yml.tpl",
	},
	{
		Path:     "/docs/README.md",
		Template: "docs_readme.tpl",
	},
	{
		Path:     "/pkg",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/pkg/{commands_package}",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/pkg/{commands_package}/commands.go",
		Template: "pkg_commands.go.tpl",
		required: true,
	},
	{
		Path:     "/pkg/{commands_package}/commands_encoding.go",
		Template: "pkg_commands_encoding.go.tpl",
		required: true,
	},
	{
		Path:     "/pkg/{events_package}",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/pkg/{events_package}/events.go",
		Template: "pkg_events.go.tpl",
		required: true,
	},
	{
		Path:     "/pkg/{events_package}/events_encoding.go",
		Template: "pkg_events_encoding.go.tpl",
		required: true,
	},
	{
		Path:     "/pkg/{stream_package}",
		IsDir:    true,
		required: true,
	},
	{
		Path:     "/pkg/{stream_package}/stream.go",
		Template: "pkg_stream.go.tpl",
		required: true,
	},
	{
		Path:  "/pkg/proto",
		IsDir: true,
	},
	{
		Path:  "/pkg/client",
		IsDir: true,
	},
	{
		Path:     "/pkg/client/http.go",
		Template: "pkg_client_http.go.tpl",
	},
	{
		Path:     "/pkg/client/grpc.go",
		Template: "pkg_client_grpc.go.tpl",
	},
	{
		Path:     "/pkg/proto/service.proto",
		Template: "pkg_proto_service.tpl",
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
		Path:     "/docker/mockgen.dockerfile",
		Template: "docker_mockgen.tpl",
	},
	{
		Path:     "/docker/protoc.dockerfile",
		Template: "docker_protoc.tpl",
	},
	{
		Path:  "/scripts",
		IsDir: true,
	},
	{
		Path:     "/scripts/mockgen.bash",
		Template: "scripts_mockgen_bash.tpl",
	},
	{
		Path:     "/scripts/protoc.bash",
		Template: "scripts_protoc_bash.tpl",
	},
	{
		Path:  "/mocks",
		IsDir: true,
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
