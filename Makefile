#! /usr/bin/make
#
# Makefile for github.com/eric-isakson/goa-playground
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "gen" invokes the "goa" tool to generate the examples source code
# - "build" compiles the example microservices and client CLIs
# - "clean" deletes the output of "build"
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
MODULE=$(shell go list -m)
MODULE_DIR=$(shell go list -m -f "{{.Dir}}")
GO_FILES=$(shell find . -type f -name '*.go')

export GO111MODULE=on

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	github.com/hashicorp/go-getter \
	github.com/cheggaaa/pb \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/golang/protobuf/proto \
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	honnef.co/go/tools/cmd/staticcheck

.PHONY: all travis depend lint gen build clean test check-freshness

all: gen lint test build
	@echo DONE!

travis: depend all check-freshness

# Install protoc
GOOS=$(shell go env GOOS)
GOPATH?=$(shell go env GOPATH)
PROTOC_VERSION=3.6.1
ifeq ($(GOOS),linux)
PROTOC=protoc-$(PROTOC_VERSION)-linux-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
GOBIN=$(GOPATH)/bin
else
	ifeq ($(GOOS),darwin)
PROTOC=protoc-$(PROTOC_VERSION)-osx-x86_64
PROTOC_EXEC=$(PROTOC)/bin/protoc
GOBIN=$(GOPATH)/bin
	else
		ifeq ($(GOOS),windows)
PROTOC=protoc-$(PROTOC_VERSION)-win32
PROTOC_EXEC="$(PROTOC)\bin\protoc.exe"
GOBIN="$(GOPATH)\bin"
		endif
	endif
endif
depend:
	@echo INSTALLING DEPENDENCIES...
	@env GO111MODULE=off go get -v $(DEPEND)
	@go get -v goa.design/goa/v3
	@go get -v goa.design/goa/v3/...
	@env GO111MODULE=off go install github.com/hashicorp/go-getter/cmd/go-getter && \
		go-getter https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC).zip $(PROTOC) && \
		cp $(PROTOC_EXEC) $(GOBIN) && \
		rm -r $(PROTOC)
	@env GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -v ./...

lint:
	@echo LINTING CODE...
	@if [ "`goimports -l $(GO_FILES) | grep -v .pb.go | tee /dev/stderr`" ]; then \
		echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
	fi
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi
	@if [ "`staticcheck -checks all,-ST1000,-ST1001,-ST1021 ./... | grep -v ".pb.go" | tee /dev/stderr`" ]; then \
		echo "^ - staticcheck errors!" && echo && exit 1; \
	fi

gen:
	@echo GENERATING CODE...
	@rm -rf "$(MODULE_DIR)/cmd" && \
	goa gen     $(MODULE)/design  && \
	goa example $(MODULE)/design 

build:
	@go build ./cmd/playground && \
	go build ./cmd/playground-cli

clean:
	@rm -f playground playground-cli

test:
	@echo TESTING...
	@go test ./... > /dev/null

check-freshness:
	@if [ "`git diff | wc -l`" -gt "0" ]; then \
		echo "[ERROR] generated code not in-sync with design:"; \
		echo; \
		git status -s; \
		git --no-pager diff; \
		echo; \
		exit 1; \
	fi
