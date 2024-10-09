# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
PROTO_SOURCES = $(call rwildcard,,*.proto)

.PHONY: lint
lint:
	clang-format -style="{BasedOnStyle: Google, IndentWidth: 2, AlignConsecutiveDeclarations: true, AlignConsecutiveAssignments: true, ColumnLimit: 0}" -i $(PROTO_SOURCES)
	buf lint

.PHONY: init
init:
	go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: update
update:
	buf dep update
