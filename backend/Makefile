build:
	go build -o bin/news-sentiment *.go

run:
	go run main.go routes.go

test:
	go test
	go test github.com/ubclaunchpad/news-sentiment/db

## lint: run static analysis for entire project
.PHONY: lint
lint: SHELL:=/bin/bash
lint:
	go vet ./...