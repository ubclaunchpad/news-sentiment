build:
	go build -o bin/news-sentiment *.go

run:
	go run main.go routes.go

## lint: run static analysis for entire project
.PHONY: lint
lint: SHELL:=/bin/bash
lint:
	go vet ./...