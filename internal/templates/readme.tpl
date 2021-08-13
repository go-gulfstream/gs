# {{$.Name}} - {{$.PackageName}}
{{$.Description}}


```shell script
$ go run ./cmd/{{$.PackageName}} -config config/stream.config.yml
$ go run ./cmd/{{$.PackageName}}-projection -config config/projection.config.yml
```