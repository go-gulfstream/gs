DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

.PHONY: docker
docker:
	@docker build --build-arg BUILD_VERSION=$(VERSION) -t gulstream/gs:latest -f ./docker/gs.dockerfile .
	@docker image prune --filter label=stage=builder

.PHONE: clean
clean:
	@rm -rfv build

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$(VERSION)'" -o build/gs-linux-amd64 ./cmd/gs

.PHONY: build-mac
build-mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$(VERSION)'" -o build/gs-darwin-amd64 ./cmd/gs

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: lint
test:
	@go test ./... -bench=. -benchtime=5s -count 5 -benchmem

.PHONY: test-cover-html
test-cover:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm -f coverage.out

.PHONY: test-cover-func
test-cover-func:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out
	@rm -f coverage.out

.PHONY: lint
lint:
	@docker run --rm -v $(PWD):/gs --workdir /gs golangci/golangci-lint:latest golangci-lint run -v --timeout 300s