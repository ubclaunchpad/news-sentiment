build:
	go build -o bin/news-sentiment *.go

run:
	go run *.go

## lint: run static analysis for entire project
.PHONY: lint
lint: SHELL:=/bin/bash
lint:
	go vet ./...