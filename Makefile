MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
CMD_NAME = "demo"
DOCKER_TAG_VERSION= "0.0.1"

.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:  ## build the cli binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o $(CMD_NAME) $(MODULE)/

.PHONY: test
test:  ## testing the cli binary
	go test -json -covermode=atomic -coverpkg=../../demo/...  ./...

.PHONY: run-query
run-query: ## run the cli
	go run ${LDFLAGS} . query

.PHONY: run-query
run-command: ## run the cli
	go run ${LDFLAGS} . command

.PHONY: build-docker
build-docker: ## build the cli as a docker image
	docker build -f Dockerfile -t $(CMD_NAME) .

.PHONY: version
version: ## display the version of the cli
	@echo $(VERSION)


.PHONY: dev-env-start
dev-env-start: ## use this command to start what you need

	docker run --rm --name demo-mongo \
		-e MONGO_INITDB_ROOT_USERNAME=root -e  MONGO_INITDB_ROOT_PASSWORD=example -d \
		-p 27017:27017 mongo

	docker run --rm --name demo-mongo-express \
		-e ME_CONFIG_MONGODB_SERVER=172.17.0.2 -e ME_CONFIG_MONGODB_PORT=27017 \
		-e ME_CONFIG_MONGODB_ADMINUSERNAME=root -e ME_CONFIG_MONGODB_ADMINPASSWORD=example -d \
		-p 8081:8081 mongo-express

.PHONY: dev-env-stop
dev-env-stop: ## stop the services
	docker stop demo-mongo demo-mongo-express