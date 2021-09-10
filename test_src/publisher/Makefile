.PHONY: build, lint, nats

build:
	go build -v ./cmd/app/

lint:
	golint ./... && golangci-lint run

nats:
	docker run -p 4222:4222 -ti nats:latest

.DEFAULT_GOAL := build