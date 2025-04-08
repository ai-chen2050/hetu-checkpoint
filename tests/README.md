# Hetu Checkpoint Tests

This document provides guidelines for running tests and important test cases for the Hetu Checkpoint project.

## Project Setup

### Prerequisites
- Go 1.23.4
- Docker and Docker Compose
- Make
- Git

### Environment Setup

1. Clone the repository:
```bash
git clone https://github.com/hetuproject/checkpoint.git
cd checkpoint
```

2. Install dependencies:
```bash
make deps
```

3. Build the project:
```bash
make build
```

## Test Categories

### 1. Unit Tests
Run unit tests:
```bash
make test
```

### 2. Integration Tests
Run integration tests:
```bash
make test-integration
```

### 3. Smoke Tests
Run smoke tests:
```bash
make test-smoke
```

## Core Test Cases

### Dispatcher Service

1. **Checkpoint Creation**
   - Test successful checkpoint creation with valid data
   - Test checkpoint creation with invalid data
   - Test checkpoint creation with different data sizes

2. **Checkpoint Validation**
   - Test successful validation of valid checkpoints
   - Test validation of corrupted checkpoints
   - Test validation of expired checkpoints
   - Test validation with different validation rules

3. **Storage Integration**
   - Test checkpoint storage in database backends
   - Test checkpoint retrieval
   - Test storage cleanup
   - Test storage capacity limits
   - Test storage failure scenarios

4. **API Endpoints**
   - Test all REST API endpoints
   - Test API rate limiting
   - Test API authentication/whitelist

### Validator Service

1. **Validation Logic**
   - Test signature verification
   - Test timestamp validation
   - Test data integrity checks
   - Test custom validation rules
   - Test validation performance

2. **Integration Tests**
   - Test validator-dispatcher communication
   - Test validator-storage integration
   - Test validator-API integration
   - Test validator scaling
   - Test validator failover

### Performance Tests

1. **Load Testing**
   - Test system under high concurrent requests
   - Test system with large data volumes
   - Test system with different request patterns
   - Test system resource usage
   - Test system recovery after load

2. **Stress Testing**
   - Test system under extreme conditions
   - Test system with network latency
   - Test system with storage delays
   - Test system with partial failures
   - Test system recovery mechanisms

### Security Tests

1. **Authentication & Authorization**
   - Test API key validation
   - Test role-based access control
   - Test permission boundaries
   - Test token expiration
   - Test security headers

2. **Data Protection**
   - Test data encryption
   - Test data integrity
   - Test data privacy
   - Test secure communication
   - Test audit logging

## Test Environment

### Local Development
```bash
# Start local development environment
make dev-up

# Run all tests
make test-all

# Clean up
make dev-down
```

### CI/CD Pipeline
Tests are automatically run in the CI/CD pipeline for:
- Pull requests
- Main branch commits
- Release tags

## Test Results

Test results are available in:
- Console output
- `test-results/` directory
- CI/CD pipeline artifacts

## Contributing

When adding new features:
1. Add corresponding unit tests
2. Add integration tests if applicable
3. Update this documentation if needed
4. Ensure all tests pass before submitting PR 