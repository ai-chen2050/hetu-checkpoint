FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make gcc musl-dev

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o validator ./cmd/validator

# Create a minimal image
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from the builder stage
COPY --from=builder /app/validator .

# Create directories for configuration and data
RUN mkdir -p /app/config /app/keys

# Set executable permissions
RUN chmod +x /app/validator

# Expose ports
EXPOSE 9081

# Set the entrypoint
ENTRYPOINT ["/app/validator"]
CMD ["run", "--config", "/app/config/val_config.json"] 