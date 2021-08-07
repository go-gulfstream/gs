module {{$.GoModules}}

go {{$.GoVersion}}

require (
	github.com/go-gulfstream/gulfstream v0.0.0-20210711164342-4e9fbadbbe03 // indirect
	github.com/google/uuid v1.2.0
	{{range $.ImportEvents}}
       "{{.Path}}"
    {{end}}
)