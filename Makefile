watch:
	chmod +x ./bin/watch.sh
	./bin/watch.sh

build:
	go build

test:
	go clean -cache
	go test -v ./...

build-github:
	go build -ldflags="-X github.com/mattia37773/mt/config.Version=$$(git describe --tags --always) -X github.com/mattia37773/mt/config.BuildMethod=github"