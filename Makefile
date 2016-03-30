#
# environment

all: test

.PHONY: test fmt

fmt:
	go fmt ./...

test: deps
	go test ./...

deps:
	go get -v ./...
