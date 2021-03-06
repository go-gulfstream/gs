# gs generator
Standard Tooling for Go-Gulfstream Development

Make golang microservices based on event-driven architecture

# Status
*Not ready for production. 
In the process of testing and fixing in internal projects.*

## Table of contents
- [Requirements](#requirements)
- [Installation](#installation)
- [Create a new manifest file](#1-create-a-new-manifest-file-for-project)
- [Initialize a new project](#2-initialize-a-new-project)
- [Add mutations](#3-add-mutations)
- [Apply changes](#4-apply-changes-to-the-project)
- [A short example](#a-short-example)
- [Manifest](docs/manifest.md)

### Requirements
[Install](https://golang.org/doc/install) golang is recommended 

### Installation
#### From source
```shell script
$ git clone git@github.com:go-gulfstream/gs.git
$ cd gs 
$ make build-linux 
$ make build-mac
```

#### Binary
```shell script
$ go install github.com/go-gulfstream/gs/cmd/gs@latest
```

#### Docker 
[Repository](https://hub.docker.com/r/gulstream/gs)
```shell script
$ docker pull gulstream/gs:latest
$ docker run --rm docker.io/gulstream/gs:latest version
```

### 1. Create a new manifest file for project
With empty manifest file
```shell script
$ gs manifest path/to/project

$ docker run -v path/to/project:/gs -w /gs docker.io/gulstream/gs:latest manifest /gs
```

With interactive mode  
```shell script
$ gs manifest -i path/to/project
```

With data example
```shell script
$ gs manifest -d path/to/project 

$ docker run -v path/to/project:/gs -w /gs docker.io/gulstream/gs:latest manifest -d /gs
```

### 2. Initialize a new project
```shell script
$ gs init path/to/project

$ docker run -v path/to/project:/gs -w /gs docker.io/gulstream/gs:latest init /gs
```

### 3. Add mutations
```shell script
$ gs add path/to/project
$ gs apply path/to/project
```
OR short entry
```shell script
$ gs add -a path/to/project 
```
OR [See step 4](#4-apply-changes-to-the-project)

### 4. Apply changes to the project 
Edit the ```path/to/project/gulfstream.yml``` manifest file 

Add [command mutations](docs/add_command_mutation.md) OR/AND [event mutations](docs/add_event_mutation.md) 

Then execute apply command:
```shell script
$ gs apply path/to/project  

$ docker run -v path/to/project:/gs -w /gs docker.io/gulstream/gs:latest apply /gs
```

### A short example
```shell script
$ mkdir ~/myproject
$ gs manifest -d ~/myproject
$ gs init ~/myproject
$ gs add -a ~/myproject
```