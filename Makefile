watch:
	chmod +x ./bin/watch.sh
	./bin/watch.sh

build:
	go build

test-c:
	go clean -cache
	go test -ldflags="-X 'github.com/mattia37773/mt/config.Environment=test'" ./...

test:
	go clean -cache
	go test -v -ldflags="-X 'github.com/mattia37773/mt/config.Environment=test'" ./...

build-github:
	go build -ldflags="-X github.com/mattia37773/mt/config.Version=$$(git describe --tags --always) -X github.com/mattia37773/mt/config.BuildMethod=github"
