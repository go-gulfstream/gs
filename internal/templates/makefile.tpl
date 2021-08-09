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