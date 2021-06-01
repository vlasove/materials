.PHONY: test, lint, doc

test:
	go test -v ./...

lint:
	golint ./... && golangci-lint run

doc:
	godoc -http :8000

.DEFAULT_GOAL:= test