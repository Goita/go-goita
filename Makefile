export GO111MODULE=on

GIT_COMMIT_HASH=$(shell git rev-parse --short HEAD)
ifeq ($(VERSION),)
VERSION := $(GIT_COMMIT_HASH)
endif

lint:
	go list ./... | xargs golint -set_exit_status
	go vet ./...

test:
	CGO_ENABLED=0 go test -cover -coverprofile=./coverage.out ./...
	go tool cover -func=coverage.out | grep "total:"

