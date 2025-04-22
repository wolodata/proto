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
	buf format -w
	buf lint

.PHONY: init
init:
	go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: update
update:
	buf dep update
