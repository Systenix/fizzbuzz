REPO_ROOT ?= $(shell pwd)

all: clean build

all_debug: clean build run-debug

rebuild:
	rm -rf bin/
	mkdir bin
	build

build:
	rm -rf bin/fizzbuzz
	# dep ensure
	go mod download && go mod verify
	go build -a --ldflags '-extldflags "-static"' -o bin/fizzbuzz ${REPO_ROOT}/cmd/main.go

# Run lint
lint:
	golangci-lint run ./cmd/...

clean:
	rm -rf bin/

run:
	./bin/fizzbuzz

test:
	go get github.com/stretchr/testify
	go test -v -cover ./...

.PHONY: clean all build run