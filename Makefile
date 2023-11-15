# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

GO_MODULE_NAME = github.com/wolodata/api
VERSION_FLAG=-X '$(GO_MODULE_NAME)/internal/pkg/version.gitBranch=`git branch --show-current`' \
-X '$(GO_MODULE_NAME)/internal/pkg/version.gitCommit=`git rev-parse HEAD`' \
-X '$(GO_MODULE_NAME)/internal/pkg/version.gitTag=`git describe --always`' \
-X '$(GO_MODULE_NAME)/internal/pkg/version.buildUser=`whoami`' \
-X '$(GO_MODULE_NAME)/internal/pkg/version.buildDate=`date +'%Y-%m-%dT%H:%M:%SZ'`'

GO_LDFLAGS=-ldflags "-s $(VERSION_FLAG)"

.PHONY: apiserver
apiserver:
	rm -rf ./dist
	GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o ./dist/wolodata-apiserver-linux-x86 main.go
	GOOS=darwin GOARCH=arm64 go build $(GO_LDFLAGS) -o ./dist/wolodata-apiserver-darwin-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build $(GO_LDFLAGS) -o ./dist/wolodata-apiserver-darwin-x86 main.go

.PHONY: lint
lint:
	goimports-reviser -format -set-alias -project-name ${GO_MODULE_NAME} -rm-unused ./...
	test -z "$$(buf format -d . | tee /dev/stderr)"
	go vet ./...
	golangci-lint run
	staticcheck ./...
	buf lint

.PHONY: generate
generate:
	rm -rf internal/gen
	buf generate

.PHONY: init
init:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
	go install github.com/incu6us/goimports-reviser/v3@latest
