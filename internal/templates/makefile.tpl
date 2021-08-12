.PHONY: mock
mock: docker-mock
	@bash ./scripts/mockgen.bash

.PHONY: docker-mock
docker-mock:
	@docker build -t github.com/go-gulfstream/mockgen:latest -f    \
           ./docker/mockgen.dockerfile .

.PHONY: proto
proto: docker-protoc
	@bash ./scripts/protoc.bash

.PHONY: docker-protoc
docker-protoc:
	@docker build -t github.com/go-gulfstream/protoc:latest -f   \
           ./docker/protoc.dockerfile .

.PHONY: docker-stream
docker-stream:
	@docker build --build-arg BUILD_VERSION=$(VERSION) -t {{$.GoModules}}/stream:$(VERSION) -f ./docker/stream.dockerfile .
	@docker image prune --filter label=stage=builder

.PHONY: docker-projection
docker-projection:
	@docker build --build-arg BUILD_VERSION=$(VERSION) -t {{$.GoModules}}/projection:$(VERSION) -f ./docker/projection.dockerfile .
	@docker image prune --filter label=stage=builder

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$(VERSION)'" -o build/{{$.PackageName}}-stream-linux-amd64 ./cmd/{{$.PackageName}}
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$(VERSION)'" -o build/{{$.PackageName}}-projection-linux-amd64 ./cmd/{{$.PackageName}}-projection

.PHONY: build-mac
build-mac:
    GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$(VERSION)'" -o build/{{$.PackageName}}-stream-darwin-amd64 ./cmd/{{$.PackageName}}
    GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -ldflags="-w -s -X 'main.buildVersion=$(VERSION)'" -o build/{{$.PackageName}}-projection-darwin-amd64 ./cmd/{{$.PackageName}}-projection

.PHONE: clean
clean:
	@rm -rfv build

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
	@docker run --rm -v $(PWD):/lint --workdir /lint golangci/golangci-lint:latest golangci-lint run -v --timeout 300s