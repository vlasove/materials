.PHONY: test, lint, doc, cover, bench, race
test:
	go test -v ./...

lint:
	golint ./... && golangci-lint run

doc:
	godoc -http :8000

cover:
	go test -cover ./...

bench:
	go test -bench=. ./...
race:
	go test -race ./...
.DEFAULT_GOAL:= test