.PHONY: test, lint, doc, cover, bench
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

.DEFAULT_GOAL:= test