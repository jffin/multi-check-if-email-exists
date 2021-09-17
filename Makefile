.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

test: vet
	go test ./... -cover
.PHONY: test

build: test
	go build -o bin/multi-check-email-exists cmd/multi-check-email-exists/main.go
.PHONY:build
