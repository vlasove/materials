.PHONY: test, lint, doc, cover
test:
	go test -v ./...

lint:
	golint ./... && golangci-lint run

doc:
	godoc -http :8000

cover:
	go test -cover ./...

.DEFAULT_GOAL:= test