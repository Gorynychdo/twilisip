.PHONY: build
build:
	go build -v ./cmd/twiliserver

.DEFAULT_GOAL := build
