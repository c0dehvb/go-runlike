all: help

## build: Build binary executable file.
build:
	CGO_ENABLED=0 go build -o bin/gorunlike cmd/main.go

## clean: rm -rf bin/*
.PHONY: image clean
clean:
	rm -rf bin/*

## help: Show help message.
help:
	@echo "Choose a command run:"
	@sed -n 's/^##//p' Makefile | sed -e 's/:/  \t/' |  sed -e 's/^/  /'
	@echo -n "\nExamples:\n"
	@echo "  Build binary executable file:"
	@echo "    make build"
