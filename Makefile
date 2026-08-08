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
	@echo "Usage: make [target]"
	@echo ""
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

watch: ## Recompile on filechange
	@echo "watch the code"
	chmod +x ./bin/watch.sh
	./bin/watch.sh

build: ## Build the binary
	@echo "Building..."
	GIT_TAG=$(GIT_TAG) go build -ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=$(GIT_TAG)' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.BuildMethod=$(BUILD_METHOD)' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(APP_ENV)'"

test: ## Run the testsuite
	@echo "Testing with unittests..."
	go clean -cache
	go test -v ./cmd/... -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./...  | grep -v '\[no test files\]'

test-all: ## Run the entire test suite
	@echo "Run the entire testsuite"
	@echo ""
	$(MAKE) test
	@echo ""
	@echo ""
	$(MAKE) e2e

test-basic: ## Run the testsuite wihout the verbose flag
	@echo "Testing with unittests without verbose mode..."
	go clean -cache
	go test ./cmd/... -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./... | grep -v '\[no test files\]'

test-cover: ## Show the test coverage
	@echo "Shows unittest coverage..."
	go clean -cache
	go test  ./cmd/... -cover -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=test'\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./...

e2e: check-docker-compose  ## run the E2E test
	@echo "Run E2E tests..."
	go clean -cache
	go test -v ./tests/e2e/... -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=dev'\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./...  | grep -v '\[no test files\]'

clear-docker: ## This remvoes everything in docker! from all namespaces
	docker system prune -a --volumes -f

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
	