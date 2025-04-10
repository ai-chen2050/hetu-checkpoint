# Hetu Checkpoint Test Cases Documentation

## Overview

This document outlines test cases for the Hetu Checkpoint system, focusing on the Dispatcher and Validator components. These test cases are designed to verify the functionality, reliability, and security of the checkpoint mechanism.

## Prerequisites

Before running the tests, ensure you have the following:

1. Clone the repository:
   ```
   git clone https://github.com/hetuproject/checkpoint.git
   ```

2. Set up the required environment:
   - PostgreSQL database
   - Ethereum development network (local or testnet)
   - Hetu chain node (CometBFT node)


3. Configure the test environment:
   - Update `dis_config.json` and `val_config.json` with appropriate test values
   - Deploy necessary contracts (CKPTValStaking and Staking Token)

## Dispatcher Test Cases

### 1. Initialization and Configuration

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-001 | Basic Initialization | Start dispatcher with valid configuration | Dispatcher starts successfully and logs initialization |
| DISP-002 | Database Connection | Verify database connection on startup | Connection established, tables created if not exist |
| DISP-003 | Configuration Validation | Start with invalid configuration parameters | Clear error messages, graceful failure |
| DISP-004 | Ethereum Connection | Verify connection to Ethereum RPC | Connection established, chain ID verified |
| DISP-005 | Contract Connection | Verify connection to CKPTValStaking contract | Contract instance created successfully |

### 2. Validator Management

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-101 | Validator Registration | Process incoming validator registration | Registration stored in database |
| DISP-102 | Validator List Query | Query active validators from the chain | Correct validator set retrieved |
| DISP-103 | Validator Connection | Handle validator TCP connections | Connections accepted and tracked |
| DISP-104 | Validator Disconnection | Handle validator disconnection | Connection closed properly, state updated |
| DISP-105 | Multiple Validators | Test with multiple validators connected | All connections maintained correctly |

### 3. Checkpoint Generation

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-201 | Epoch Detection | Detect new epochs from blockchain | New epochs identified correctly |
| DISP-202 | Checkpoint Creation | Create checkpoint for new epoch | Checkpoint created with correct data |
| DISP-203 | Request Distribution | Distribute checkpoint requests to validators | Requests sent to all active validators |
| DISP-204 | Timeout Handling | Handle validator response timeout | Timeout correctly identified and logged |
| DISP-205 | Request Persistence | Store checkpoint requests in database | Requests stored correctly with status |

### 4. Signature Aggregation

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-301 | Response Collection | Collect responses from validators | All valid responses collected |
| DISP-302 | Signature Verification | Verify individual BLS signatures | Valid signatures accepted, invalid rejected |
| DISP-303 | Threshold Verification | Check if enough validators responded | Threshold check works correctly |
| DISP-304 | BLS Aggregation | Aggregate BLS signatures | Signatures aggregated correctly |
| DISP-305 | Power Calculation | Calculate total voting power | Power calculated correctly based on validators |
| DISP-306 | Bitmap Generation | Generate bitmap of participating validators | Bitmap correctly represents participants |

### 5. Checkpoint Reporting

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-401 | Ethereum Transaction | Submit checkpoint to Ethereum | Transaction submitted successfully |
| DISP-402 | Transaction Confirmation | Wait for transaction confirmation | Transaction confirmed, receipt stored |
| DISP-403 | Retry Mechanism | Handle submission failure with retry | Retry works with backoff strategy |
| DISP-404 | Gas Price Adjustment | Adjust gas price for transaction | Gas price adjusted according to network conditions |
| DISP-405 | Report Persistence | Store reporting status in database | Reporting status stored correctly |

### 6. Reward Distribution

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-501 | Distribution Triggering | Trigger distribution after interval | Distribution triggered after configured epochs |
| DISP-502 | Contract Interaction | Call distributeCheckpointRewards | Function called with correct parameters |
| DISP-503 | Distribution Status | Track distribution status | Status correctly tracked in database |
| DISP-504 | Multiple Epoch Distribution | Distribute rewards for multiple epochs | All epochs processed correctly |
| DISP-505 | Error Handling | Handle distribution errors | Errors logged, retries implemented |

### 7. API and Integration

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| DISP-601 | HTTP API Availability | Check HTTP API endpoints | All endpoints accessible |
| DISP-602 | Status Endpoint | Query status information | Correct status information returned |
| DISP-603 | Metrics Collection | Collect performance metrics | Metrics collected and available |
| DISP-604 | Chain Integration | Verify integration with Hetu chain | Chain data correctly processed |
| DISP-605 | CometBFT Integration | Verify integration with CometBFT | Block data correctly retrieved |

## Validator Test Cases

### 1. Initialization and Configuration

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-001 | Basic Initialization | Start validator with valid configuration | Validator starts successfully and logs initialization |
| VAL-002 | Database Connection | Verify database connection on startup | Connection established, tables created if not exist |
| VAL-003 | Configuration Validation | Start with invalid configuration parameters | Clear error messages, graceful failure |
| VAL-004 | Dispatcher Connection | Connect to dispatcher TCP service | Connection established successfully |
| VAL-005 | Key Loading | Load validator keys | Keys loaded correctly from storage |

### 2. Registration and Staking

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-101 | Role Verification | Verify validator has required roles | Role verification works correctly |
| VAL-102 | Validator Registration | Register as validator in the contract | Registration transaction successful |
| VAL-103 | Token Approval | Approve tokens for staking | Approval transaction successful |
| VAL-104 | Staking Process | Stake tokens in the contract | Staking transaction successful |
| VAL-105 | Info Update | Update validator information | Update transaction successful |

### 3. Checkpoint Signing

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-201 | Request Reception | Receive checkpoint request from dispatcher | Request received and parsed correctly |
| VAL-202 | Checkpoint Validation | Validate checkpoint data | Validation logic works correctly |
| VAL-203 | BLS Signature Generation | Generate BLS signature for checkpoint | Signature generated correctly |
| VAL-204 | Response Sending | Send response to dispatcher | Response sent successfully |
| VAL-205 | Request Persistence | Store request and response in database | Data stored correctly with status |

### 4. BLS Key Management

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-301 | Key Generation | Generate new BLS key pair | Key pair generated with correct format |
| VAL-302 | Key Storage | Store BLS keys securely | Keys stored with proper encryption |
| VAL-303 | Key Recovery | Recover keys from storage | Keys recovered correctly |
| VAL-304 | Key Rotation | Rotate BLS keys | Key rotation process works correctly |
| VAL-305 | Key Registration | Register new BLS public key | Registration transaction successful |

### 5. Reward Management

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-401 | Reward Calculation | Calculate expected rewards | Calculation matches contract implementation |
| VAL-402 | Reward Claiming | Claim rewards from contract | Claim transaction successful |
| VAL-403 | Reward Tracking | Track claimed rewards | Rewards tracked correctly in database |
| VAL-404 | Unstaking Process | Initiate and complete unstaking | Unstaking process works correctly |
| VAL-405 | Balance Verification | Verify token balances after operations | Balances match expected values |

### 6. Error Handling and Recovery

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-501 | Connection Loss | Handle dispatcher connection loss | Reconnection attempts made |
| VAL-502 | Invalid Request | Handle invalid checkpoint request | Error logged, no signature generated |
| VAL-503 | Transaction Failure | Handle Ethereum transaction failure | Proper error handling and retry |
| VAL-504 | Database Error | Handle database connection issues | Graceful degradation, recovery attempts |
| VAL-505 | Restart Recovery | Recover state after restart | State correctly recovered from database |

### 7. Security Tests

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| VAL-601 | Authentication | Verify dispatcher authentication | Only authorized dispatchers accepted |
| VAL-602 | Request Verification | Verify request integrity | Tampered requests rejected |
| VAL-603 | Double Signing Prevention | Prevent signing same checkpoint twice | Duplicate signing prevented |
| VAL-604 | Key Protection | Protect private keys from exposure | Keys never exposed in logs or responses |
| VAL-605 | Data Validation | Validate all incoming data | Invalid data rejected with appropriate errors |

## Integration Test Cases

### 1. End-to-End Checkpoint Flow

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| INT-001 | Complete Checkpoint Cycle | Run through entire checkpoint cycle | Checkpoint submitted to Ethereum successfully |
| INT-002 | Multiple Validators | Test with multiple validators | Signatures correctly aggregated from all validators |
| INT-003 | Validator Subset | Test with subset of validators responding | Threshold mechanism works correctly |
| INT-004 | Sequential Epochs | Process multiple sequential epochs | All epochs processed correctly in order |
| INT-005 | Parallel Requests | Handle multiple checkpoint requests in parallel | All requests processed correctly |

### 2. Performance and Stress Tests

| Test ID | Test Name | Description | Expected Result |
|---------|-----------|-------------|----------------|
| PERF-001 | High Validator Count | Test with high number of validators | System scales correctly |
| PERF-002 | Rapid Epoch Changes | Process rapidly changing epochs | All epochs processed without missing any |
| PERF-003 | Network Latency | Test with simulated network latency | System handles latency gracefully |
| PERF-004 | Resource Constraints | Test under limited CPU/memory | System degrades gracefully |
| PERF-005 | Long-Running Test | Run system for extended period | Stability maintained over time |

## Test Environment Setup

### Local Development Environment

```bash
# Start PostgreSQL
docker run -d --name postgres-checkpoint -e POSTGRES_PASSWORD=hetu -p 5432:5432 postgres

# Create required databases
psql -h localhost -U postgres -c "CREATE DATABASE bls_db3;"
psql -h localhost -U postgres -c "CREATE DATABASE validator_db;"

# Start local Ethereum node
npx hardhat node

# Deploy contracts
cd contracts
npx hardhat run scripts/deploy.js --network localhost

# Start Dispatcher
cd cmd/dispatcher
go run main.go

# Start Validator
cd cmd/validator
go run main.go
```

### Test Configuration Examples

**Dispatcher Configuration (dis_config.json):**
```json
{
    "DBhost": "localhost",
    "DBport": 5432,
    "DBuser": "postgres",
    "DBpassword": "hetu",
    "DBname": "bls_db3",
    "Httpport": 8080,
    "Tcpport": 8081,
    "EnableReport": true,
    "ChainGRpcURL": "localhost:9090",
    "CometBFTSvr": "localhost:26657",
    "ChainID": "hetu_560002-1",
    "EthRpcURL": "http://localhost:8545",
    "EthChainID": 31337,
    "ValidatorStakingAddress": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    "RewardDistrInterval": 3
}
```

**Validator Configuration (val_config.json):**
```json
{
    "DBhost": "localhost",
    "DBport": 5432,
    "DBuser": "postgres",
    "DBpassword": "hetu",
    "DBname": "validator_db",
    "dispatchertcp": "localhost:8081",
    "ListenAddr": "localhost",
    "port": 0,
    "ChainRpcURL": "http://localhost:8545",
    "StakingTokenAddress": "0xe7f6c60d813a17e2fa16869371351d8c800c7fc1",
    "ValidatorStakingAddress": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    "DispatcherURL": "http://localhost:8080",
    "ChainGRpcURL": "localhost:9090",
    "CometBFTSvr": "localhost:26657",
    "ChainID": "hetu_560002-1"
}
```

## Test Execution Guidelines

1. **Unit Tests**: Run individual component tests
   ```
   go test ./... -v
   ```

2. **Integration Tests**: Run through complete checkpoint flow
   ```
   go test ./tests/integration -v
   ```

3. **Performance Tests**: Measure system performance under load
   ```
   go test ./tests/performance -v -benchtime=10s
   ```

4. **Manual Tests**: Follow test cases in this document with the provided test environment

## Test Reporting

For each test, document:
1. Test case ID
2. Test date and environment
3. Test result (Pass/Fail)
4. Observations and issues encountered
5. Performance metrics (if applicable)
