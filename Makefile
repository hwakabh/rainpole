MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

# all targets are phony
.PHONY: $(shell egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')

COMPILED_FILEPATH="./cmd/rainpole"

build: ## Compile single binary for rainpole app
	@echo "Compiling single binary for rainple application ..."
	@GOOS=linux GOARCH=amd64 go build -o ${COMPILED_FILEPATH} .

test: ## Run unittest
	@echo "Run Go test ..."
	@go test -v

help: ## Print this help
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
