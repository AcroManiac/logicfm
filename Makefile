PROJECT = $(shell basename "$(PWD)")

# Use the v0.0.0 tag for testing, it shouldn't clobber any release builds
RELEASE ?= v0.0.0

PWD = $(shell pwd)

REPO_INFO=$(shell git config --get remote.origin.url)
REPO_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
RELEASE_DATE=$(shell date +%FT%T%Z)

ifndef REPO_COMMIT
REPO_COMMIT = git-$(shell git rev-parse --short HEAD)
endif

HAS_LINT := $(shell command -v golangci-lint;)

LDFLAGS = "-s -w \
	-X $(PROJECT)/pkg/version.RELEASE=$(RELEASE) \
	-X $(PROJECT)/pkg/version.DATE=$(RELEASE_DATE) \
	-X $(PROJECT)/pkg/version.REPO=$(REPO_INFO) \
	-X $(PROJECT)/pkg/version.COMMIT=$(REPO_COMMIT) \
	-X $(PROJECT)/pkg/version.BRANCH=$(REPO_BRANCH)"

GO_PACKAGES=$(shell go list $(PROJECT)/...)

all: build

bootstrap:
ifndef HAS_LINT
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint@v1.20.0
endif

vendor: bootstrap
	@echo "+ $@"
	@go mod tidy

fmt:
	@echo "+ $@"
	@gofmt -w ./pkg/

test:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -race -cover {{.Dir}}"{{end}}' $(GO_PACKAGES) | xargs -L 1 sh -c

lint: bootstrap fmt
	@echo "+ $@"
	@golangci-lint run ./...

build: vendor lint test
	@echo "+ $@"
	@go build -a -ldflags $(LDFLAGS) -o bin/$(BIN) $(PROJECT)/cmd/builder

.PHONY: all fmt bootstrap build lint vendor
