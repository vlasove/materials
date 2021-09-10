.PHONY: build, test, bench

build:
	go build -v -o app .

test:
	go test -v ./...

bench:
	go test -v -bench=. ./...

.DEFAULT_GOAL := build