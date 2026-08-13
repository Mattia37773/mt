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
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=$(BUILD_METHOD)'"

test: check-docker-compose  ## Run the testsuite
	@echo "Testing with unittests..."
	@echo

	@echo "Removing the binaries"
	@rm -f tests/bin/homebrew-install
	@rm -f tests/bin/go-install
	@rm -f tests/bin/source-install

	@echo "Building binarie for testing update"
	@echo
# source install
	GIT_TAG=$(GIT_TAG) go build \
		-ldflags="\
			-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
			-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=source'" \
		-o tests/bin/source-install
# go install
	GIT_TAG=$(GIT_TAG) go build \
	-ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0'" \
	-o tests/bin/go-install
# homebrew install
	GIT_TAG=$(GIT_TAG) go build \
	-ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=homebrew'" \
	-o tests/bin/homebrew-install

	go clean -cache
	go test -v -p 1 ./tests/cmd/... -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./...  | grep -v '\[no test files\]'

# removing the binaries
	@echo "Removing the binaries"
	@rm -f tests/bin/homebrew-install
	@rm -f tests/bin/go-install
	@rm -f tests/bin/source-install

test-basic: ## Run the testsuite wihout the verbose flag
	@echo "Testing with unittests without verbose mode..."
		@echo

	@echo "Removing the binaries"
	@rm -f tests/bin/homebrew-install
	@rm -f tests/bin/go-install
	@rm -f tests/bin/source-install

	@echo "Building binarie for testing update"
	@echo
# source install
	GIT_TAG=$(GIT_TAG) go build \
		-ldflags="\
			-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
			-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=source'" \
		-o tests/bin/source-install
# go install
	GIT_TAG=$(GIT_TAG) go build \
	-ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0'" \
	-o tests/bin/go-install
# homebrew install
	GIT_TAG=$(GIT_TAG) go build \
	-ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=homebrew'" \
	-o tests/bin/homebrew-install

	go clean -cache
	go test -p 1 ./tests/cmd/... -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./...  | grep -v '\[no test files\]'

	# removing the binaries
	@echo "Removing the binaries"
	@rm -f tests/bin/homebrew-install
	@rm -f tests/bin/go-install
	@rm -f tests/bin/source-install

test-cover: ## Show the test coverage
	@echo "Shows unittest coverage..."
		@echo "Testing with unittests..."
	@echo

	@echo "Removing the binaries"
	@rm -f tests/bin/homebrew-install
	@rm -f tests/bin/go-install
	@rm -f tests/bin/source-install

	@echo "Building binarie for testing update"
	@echo
# source install
	GIT_TAG=$(GIT_TAG) go build \
		-ldflags="\
			-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
			-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=source'" \
		-o tests/bin/source-install
# go install
	GIT_TAG=$(GIT_TAG) go build \
	-ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0'" \
	-o tests/bin/go-install
# homebrew install
	GIT_TAG=$(GIT_TAG) go build \
	-ldflags="\
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
		-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideBuildMethod=homebrew'" \
	-o tests/bin/homebrew-install

	go clean -cache
	go test -v -p 1 ./tests/cmd/... -ldflags="\
	-X '$(MODULE_PATH)/$(APP_NAME)/config.OverrideCurrentVersion=v1.0.0' \
	" ./...  | grep -v '\[no test files\]'

	# removing the binaries
	@echo "Removing the binaries"
	@rm -f tests/bin/homebrew-install
	@rm -f tests/bin/go-install
	@rm -f tests/bin/source-install

clear-docker: ## This remvoes everything in docker! from all namespaces
	docker rm -f $$(docker ps -aq) 2>/dev/null || true
	docker system prune -a --volumes -f

docker-show: ## Show all containers, images & volumes
	@echo "All running docker containers"	
	docker ps
	@echo ""
	@echo "All docker images"
	docker image ls
	@echo ""
	@echo "All docker volumes"
	docker volume ls

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
	

