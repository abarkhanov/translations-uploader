# Go Options
MODULE   := github.com/abarkhanov/ttu
LDFLAGS  := -w -s
BINDIR   := $(CURDIR)/bin
GOBIN    := $(BINDIR)
PATH     := $(GOBIN):$(PATH)
NAME     := ttu
VERSION  := "unknown"
COVEROUT := ./coverage.out

# Tools as dependencies
TOOLS += github.com/mattn/goveralls
TOOLS += github.com/maxbrunsfeld/counterfeiter/v6
TOOLS += github.com/golangci/golangci-lint/cmd/golangci-lint

# Verbose output
ifdef VERBOSE
V = -v
else
.SILENT:
endif

# Git dependencies
HAS_GIT := $(shell command -v git;)
ifndef HAS_GIT
	$(error Please install git)
endif

# Git Status
GIT_SHA ?= $(shell git rev-parse --short HEAD)

# Default target
.DEFAULT_GOAL := all

# Make All targets
.PHONY: all
all: deps build

# Download dependencies to go module cache
.PHONY: deps
deps:
	$(info Installing dependencies)
	XDG_CONFIG_HOME=$(CURDIR)/configs go mod download

# Builds Binary
.PHONY: build
build: deps
build: LDFLAGS += -X $(MODULE)/pkg/version.Timestamp=$(shell date +%s)
build: LDFLAGS += -X $(MODULE)/pkg/version.Version=${VERSION}
build: LDFLAGS += -X $(MODULE)/pkg/version.GitSHA=${GIT_SHA}
build: LDFLAGS += -X $(MODULE)/pkg/version.ServiceName=${NAME}
build:
	$(info building binary to bin/$(NAME))
	@CGO_ENABLED=0 go build -o bin/$(NAME) -installsuffix cgo -ldflags '$(LDFLAGS)' ./cmd/ttu

# Builds and runs the binary with debug logging
.PHONY: run
run: build
	@LOG_LEVEL=debug ./bin/$(NAME) upload

# Clean all the things
.PHONY: clean
clean:
	@rm -f *.cov
	@rm bin/$(NAME)