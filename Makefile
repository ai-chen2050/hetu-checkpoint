.PHONY: build clean deps test docker-dispatcher docker-validator

GO := go
GOFLAGS := -v
BINARY_DIR := build
DISPATCHER_BINARY := $(BINARY_DIR)/dispatcher
VALIDATOR_BINARY := $(BINARY_DIR)/validator

all: build

build: build-dispatcher build-validator

# Build binaries
build-dispatcher:
	@mkdir -p $(BINARY_DIR)
	$(GO) build $(GOFLAGS) -o $(DISPATCHER_BINARY) ./cmd/dispatcher

build-validator:
	@mkdir -p $(BINARY_DIR)
	$(GO) build $(GOFLAGS) -o $(VALIDATOR_BINARY) ./cmd/validator

deps:
	$(GO) mod download

# Clean build artifacts
clean:
	rm -rf $(BINARY_DIR)
	$(GO) clean

# Run tests
test:
	$(GO) test -v ./...

# Build Docker images
docker-dispatcher:
	docker build -t hetuproject/checkpoint-dispatcher:latest -f docker/dispatcher/Dockerfile .

docker-validator:
	docker build -t hetuproject/checkpoint-validator:latest -f docker/validator/Dockerfile .

# Start dispatcher with Docker Compose
start-dispatcher:
	docker-compose -f docker-compose-dispatcher.yml up -d

# Start validator with Docker Compose
start-validator:
	docker-compose -f docker-compose-validator.yml up -d

# Stop all containers
stop:
	docker-compose -f docker-compose-dispatcher.yml down
	docker-compose -f docker-compose-validator.yml down 