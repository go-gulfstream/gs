.PHONY: mock
mock: docker-mock
	@bash ./scripts/mockgen.bash

.PHONY: docker-mock
docker-mock:
	@docker build -t github.com/go-gulfstream/mockgen:latest -f    \
           ./docker/mockgen.dockerfile .