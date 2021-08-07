DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

GOOS    ?= linux
GOARCH  ?= amd64
BUILD = CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s" -o build/gs-$(GOOS)-$(GOARCH) ./cmd/gs

.PHONY: docker
docker:
	@docker build -t gulstream/gs:latest -f ./docker/gs.dockerfile .
	@docker image prune --filter label=stage=builder

.PHONY: build
build:
	$(BUILD)

.PHONY: build-linux
build-linux:
	 GOOS=linux GOARCH=amd64 $(BUILD)

.PHONY: build-mac
build-mac:
	 GOOS=darwin GOARCH=amd64 $(BUILD)

.PHONY: version
version:
	@echo $(VERSION)