# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

.PHONY: lint
lint:
	buf lint

.PHONY: generate
export PATH := $(PWD)/node_modules/.bin:$(PATH)
generate:
	rm -rf gen
	buf generate

.PHONY: init
init:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	npm install --save-dev @bufbuild/buf @bufbuild/protoc-gen-es @connectrpc/protoc-gen-connect-es
    npm install @connectrpc/connect @bufbuild/protobuf
