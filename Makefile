APP_ENV         ?= dev
MODULE_PATH 	:= github.com/mattia37773
APP_NAME        := mt
BUILD_METHOD 	:= source
GIT_TAG 		:= $(shell git describe --tags --always 2>/dev/null || echo "dev")

.PHONY: test
.PHONY: watch
.PHONY: test-all
.PHONY: test-basic
.PHONY: test-cover
.PHONY: build
.PHONY: e2e

help: ## Shows help for all command
	@echo "Nutzung: make [target]"
	@echo ""
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

watch: ## Recompile on filechange
	chmod +x ./bin/watch.sh
	./bin/watch.sh

build: ## Build the binary
	GIT_TAG=$(GIT_TAG) go build -ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.Version=$(GIT_TAG)' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.BuildMethod=$(BUILD_METHOD)' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(APP_ENV)'"

test: ## Run the testsuite
	go clean -cache
	go test -v ./cmd/... -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'" ./...  | grep -v '\[no test files\]'

test-all: ## Run the testsuite with showing directories without tests
	go clean -cache
	go test -v ./cmd/... -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'" ./... 

test-basic: ## Run the testsuite wihout the verbose flag
	go clean -cache
	go test ./cmd/... -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'" ./...

test-cover: ## Show the test coverage
	go clean -cache
	go test  ./cmd/... -cover -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'" ./...

e2e: check-docker-compose  ## run the E2E tests
	go clean -cache
	go test -v ./tests/e2e/... -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=dev'" ./...  | grep -v '\[no test files\]'

check-docker-compose:
	@if ! docker compose version >/dev/null 2>&1; then \
		echo "Docker Compose isn't installed"; \
		echo "please install it"; \
		exit 1; \
	fi
	@if ! docker info >/dev/null 2>&1; then \
		echo "The docker engine isn't running"; \
		exit 1; \
	fi
	