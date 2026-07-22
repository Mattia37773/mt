watch:
	chmod +x ./bin/watch.sh
	./bin/watch.sh

build:
	go build

test:
	go clean -cache
	go test -v ./...