TEST_ENV        ?= test
APP_ENV         ?= dev
TEST  			:= cmd
MODULE_PATH 	:= github.com/mattia37773
APP_NAME        := mt
BUILD_METHOD 	:= source
GIT_TAG 		:= $(shell git describe --tags --always 2>/dev/null || echo "unknown")

.PHONY: test
.PHONY: watch
.PHONY: test-all
.PHONY: test-basic
.PHONY: test-cover
.PHONY: build

help: ## Shows help for all command
	@echo "Nutzung: make [target]"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

watch: ## Recompile on filechange
	chmod +x ./bin/watch.sh
	./bin/watch.sh

test: ## Run the testsuite
	go clean -cache
	go test -v ./$(TEST)/... -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(TEST_ENV)'" ./...  | grep -v '\[no test files\]'

test-all: ## Run the testsuite with showing directories without tests
	go clean -cache
	go test -v -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(TEST_ENV)'" ./... 

test-basic: ## Run the testsuite wihout the verbose flag
	go clean -cache
	go test -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(TEST_ENV)'" ./...

test-cover: ## Show the test coverage
	go clean -cache
	go test -cover -ldflags="-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(TEST_ENV)'" ./...

build: ## Build the binary
	GIT_TAG=$(GIT_TAG) go build -ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.Version=$(GIT_TAG)' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.BuildMethod=$(BUILD_METHOD)' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.Environment=$(APP_ENV)'"
