# {{$.Name}} - {{$.PackageName}}
{{$.Description}}

devenv
```shell script
$ git clone https://github.com/go-gulfstream/devenv.git
$ cd devenv
$ make up
```

configuration
```shell script
$ vim config/stream.config.yml
$ vim config/projection.config.yml
```

run
```shell script
$ go run ./cmd/{{$.PackageName}} -config config/stream.config.yml
$ go run ./cmd/{{$.PackageName}}-projection -config config/projection.config.yml
```