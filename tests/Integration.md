# Integration Test Documents

# Environment compilation and startup

## Prerequisites

- Go 1.23.4
- Docker and Docker Compose
- Make
- Git
- PostgreDB

## Repository

### **Checkpoint**

1. Clone the repository:

```bash
git clone https://github.com/hetuproject/checkpoint.git
cd checkpoint
```

1. Install dependencies:

```bash
make deps
```

1. Build the project:

```bash
make
```

### Hetu-chain

1. Clone the repository:

```bash
git clone https://github.com/hetu-project/hetu-chain.git
git checkout checkpoints
cd hetu-chain
```

1. Build the project:

```bash
make install
```

## Setup

### Hetu-chain

- **Local Deployment**

To deploy locally, use the `local_node.sh` script. This script will set up a local environment for running the Hetu Chain node.

```bash
./local_node.sh

# export privatekey for signing tx
hetud  keys unsafe-export-eth-key dev0  --home "$HOME/.tmp-hetud"  --keyring-backend test
```

### Contract

```bash
# 0. install deps
cd checkpoint/contracts
npm i # install dep

# 1. compile
cp .hardhat.config.js hardhat.config.js
vim hardhat.config.js 
# put privatekey to accounts

# 2. run test and deploy
npx hardhat compile
npx hardhat test
npx hardhat run scripts/deployHGT.js --network local

# Update the latest erc20 token
vim deployCKPT.js
# const Token = "Latest Token Address";

# Or local chain
npx hardhat run scripts/deployCKPT.js --network local

# 3. regist ckptStaking contract to hetu-chain
hetud tx checkpointing regist-ckpt-stake-contract **contract-address**  --from dev0  --home "$HOME/.tmp-hetud"  --keyring-backend test --gas 3000000 --gas-prices 100gas
```

### Checkpoint Networks

```bash
# 0. start postgredb
docker run -d  --name postgres-1024  -e POSTGRES_USER=postgres  -e POSTGRES_PASSWORD=hetu  -e POSTGRES_DB=checkpoint  -p 5432:5432  postgres:latest  -c "max_connections=1024"

# 1. gen keypairs
cd checkpoint
chmod +x ./scripts/start.sh ./scripts/stop.sh
./scripts/start.sh -n 128 -g

# 2. transfer gas fee & hetu token to all validators, and grant roles
cd contracts
vim scripts/initial_valstaking.js
# replace privateKey,  valStakingAddress, stakingTokenAddress(hetu token)
npx hardhat run scripts/initial_valstaking.js  --network local

# 3. update config file
 cd ..   # chdir to root dir of project
 vim docs/config/dis_config.json
 # replace **ValidatorStakingAddress** with CKPTStaking contract address
 vim docs/config/val_config.json
 # replace **StakingTokenAddress** with Hetu Erc20 token contract address
 # replace **ValidatorStakingAddress** with CKPTStaking contract address

# 4. start a dispatcher, and 128 validator, and stake 500 hetu
./scripts/start.sh -n 128

# 5. watch dispatcher logs
tail -f logs/dispatcher.log
```

## Configurations

## Contracts

The CKPTValStaking contract requires several configuration parameters that control its behavior:

### Constructor Parameters

1. **stakingToken** (address)
    - The ERC20 token used for staking
    - Must be a valid ERC20 contract address
2. **rewardRateByLockTime** (uint256)
    - Base reward rate for staking (tokens per second)
    - Used in formula: `rewards = timeElapsed * rewardRateByLockTime * stakedAmount / 1e18`
3. **minimumStake** (uint256)
    - Minimum amount of tokens required to become a validator
    - Enforced during stake operations

### Configurable Parameters (Owner-only)

1. **unstakeLockPeriod** (uint256)
    - Time period validators must wait to withdraw tokens after initiating unstake
    - Default: 7 days
    - Configured via: `setUnstakeLockPeriod()`
2. **stakeLockPeriod** (uint256)
    - Time period before staked tokens become active
    - Default: 1 day
    - Configured via: `setStakeLockPeriod()`
3. **CKPTRewardScalingFactor** (uint256)
    - Scaling factor for checkpoint rewards
    - Default: 1
    - Must be greater than 0
    - Configured via: `setCKPTRewardScalingFactor()`

### Role-based Access Control

1. **VALIDATOR_ROLE**
    - Role required to register and operate as a validator
    - Managed via: `grantValidatorRole()`, `revokeValidatorRole()`
2. **DISPATCHER_ROLE**
    - Role required to submit checkpoints
    - Managed via: `grantDispatcherRole()`, `revokeDispatcherRole()`
3. **DISTRIBUTER_ROLE**
    - Role required to distribute checkpoint rewards
    - Managed via: `grantDistributerRole()`, `revokeDistributerRole()`

## Checkpoints

### Configuration Files Overview

The Hetu Checkpoint system contains two main configuration files:

- `dis_config.json` - Dispatcher node configuration
- `val_config.json` - Validator node configuration

### Dispatcher Configuration (dis_config.json)

| Parameter | Type | Description |
| --- | --- | --- |
| `DBhost` | String | PostgreSQL database host address |
| `DBport` | Integer | PostgreSQL database port number |
| `DBuser` | String | PostgreSQL database username |
| `DBpassword` | String | PostgreSQL database password |
| `DBname` | String | PostgreSQL database name |
| `Httpport` | Integer | HTTP server listening port for API and Validator communication |
| `Tcpport` | Integer | TCP server listening port for Validator connections |
| `EnableReport` | Boolean | Whether to enable reporting checkpoints to Ethereum |
| `ChainGRpcURL` | String | Hetu chain gRPC service endpoint address |
| `CometBFTSvr` | String | CometBFT service endpoint address |
| `ChainID` | String | Hetu blockchain network ID |
| `EthRpcURL` | String | Ethereum RPC service endpoint address |
| `EthChainID` | Integer | Ethereum chain ID |
| `ValidatorStakingAddress` | String | CKPTValStaking contract address |
| `RewardDistrInterval` | Integer | Reward distribution time interval (in epochs) |

### Validator Configuration (val_config.json)

| Parameter | Type | Description |
| --- | --- | --- |
| `DBhost` | String | PostgreSQL database host address |
| `DBport` | Integer | PostgreSQL database port number |
| `DBuser` | String | PostgreSQL database username |
| `DBpassword` | String | PostgreSQL database password |
| `DBname` | String | PostgreSQL database name |
| `dispatchertcp` | String | Dispatcher TCP service address |
| `ListenAddr` | String | Validator listening address |
| `port` | Integer | Validator listening port (0 means random assignment) |
| `ChainRpcURL` | String | Ethereum RPC service endpoint address |
| `StakingTokenAddress` | String | Staking token contract address |
| `ValidatorStakingAddress` | String | CKPTValStaking contract address |
| `DispatcherURL` | String | Dispatcher HTTP service address |
| `ChainGRpcURL` | String | Hetu chain gRPC service endpoint address |
| `CometBFTSvr` | String | CometBFT service endpoint address |
| `ChainID` | String | Hetu blockchain network ID |

# **Test Cases**

## Hetu-chain

### **Checkpoint Creation & Query**

```bash
hetud q  checkpointing raw-checkpoint 100

hetud q checkpointing raw-checkpoint-list CKPT_STATUS_ACCUMULATING
```

## Contracts

### Test Cases

Please refer to [validatorStaking.test.js](../contracts/test/validatorStaking.test.js) scripts.

## Checkpoints

Please refer to [Checkpoints](./Checkpoints.md) document.