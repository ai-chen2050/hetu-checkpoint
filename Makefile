.PHONY: build deps test clean

GO := go
GOFLAGS := -v
BINARY_DIR := build
DISPATCHER_BINARY := $(BINARY_DIR)/dispatcher
VALIDATOR_BINARY := $(BINARY_DIR)/validator

all: build

build: build-dispatcher build-validator

build-dispatcher:
	@mkdir -p $(BINARY_DIR)
	$(GO) build $(GOFLAGS) -o $(DISPATCHER_BINARY) ./cmd/dispatcher

build-validator:
	@mkdir -p $(BINARY_DIR)
	$(GO) build $(GOFLAGS) -o $(VALIDATOR_BINARY) ./cmd/validator

deps:
	$(GO) mod download

test:
	$(GO) test ./...

clean:
	rm -rf $(BINARY_DIR)
	$(GO) clean 