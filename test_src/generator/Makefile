.PHONY: build, lint

build:
	go build -v ./cmd/app/

lint:
	golint ./... && golangci-lint run
	
.DEFAULT_GOAL := build