GOHOSTOS := $(shell go env GOHOSTOS)
GOHOSTARCH := $(shell go env GOHOSTARCH)

ifndef GOOS
GOOS := $(GOHOSTOS)
endif

ifndef GOARCH
GOARCH := $(GOHOSTARCH)
endif

ifeq ($(GOOS),windows)
BIN_EXT := .exe
else
BIN_EXT :=
endif

DIR_BIN := bin
NAME := $(DIR_BIN)/fullstack-$(GOOS)-$(GOARCH)$(BIN_EXT)

GO_HOST := GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) CGO_ENABLED=1 go
GO_TARGET := GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=1 go

.PHONY: all build generate test

all: build

build: generate
	@mkdir -p -- $(dir $(NAME))
	$(GO_TARGET) build -v -tags $@ -o $(NAME) ./cmd/main

generate:
	$(GO_HOST) generate -v ./...

test:
	$(GO_HOST) test ./...
