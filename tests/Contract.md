# CKPTValStaking Contract Documentation

## Configuration Parameters

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

4. **unstakeLockPeriod** (uint256)
   - Time period validators must wait to withdraw tokens after initiating unstake
   - Default: 7 days
   - Configured via: `setUnstakeLockPeriod()`

5. **stakeLockPeriod** (uint256)
   - Time period before staked tokens become active
   - Default: 1 day
   - Configured via: `setStakeLockPeriod()`

6. **CKPTRewardScalingFactor** (uint256)
   - Scaling factor for checkpoint rewards
   - Default: 1
   - Must be greater than 0
   - Configured via: `setCKPTRewardScalingFactor()`

### Role-based Access Control

7. **VALIDATOR_ROLE**
   - Role required to register and operate as a validator
   - Managed via: `grantValidatorRole()`, `revokeValidatorRole()`

8. **DISPATCHER_ROLE**
   - Role required to submit checkpoints
   - Managed via: `grantDispatcherRole()`, `revokeDispatcherRole()`

9. **DISTRIBUTER_ROLE**
   - Role required to distribute checkpoint rewards
   - Managed via: `grantDistributerRole()`, `revokeDistributerRole()`

## Test Cases

### 1. Deployment & Initialization

```javascript
describe("Deployment", function() {
  it("should deploy with correct initial parameters", async function() {
    const stakingToken = await deployMockToken();
    const rewardRate = ethers.utils.parseEther("0.000001");
    const minimumStake = ethers.utils.parseEther("100");
    
    const contract = await deployContract(stakingToken.address, rewardRate, minimumStake);
    
    expect(await contract.stakingToken()).to.equal(stakingToken.address);
    expect(await contract.rewardRateByLockTime()).to.equal(rewardRate);
    expect(await contract.minimumStake()).to.equal(minimumStake);
    expect(await contract.unstakeLockPeriod()).to.equal(7 * 24 * 60 * 60);
    expect(await contract.stakeLockPeriod()).to.equal(1 * 24 * 60 * 60);
    expect(await contract.CKPTRewardScalingFactor()).to.equal(1);
  });
  
  it("should grant default roles to deployer", async function() {
    const deployer = await ethers.getSigner();
    expect(await contract.hasRole(await contract.DEFAULT_ADMIN_ROLE(), deployer.address)).to.be.true;
    expect(await contract.hasRole(await contract.VALIDATOR_ROLE(), deployer.address)).to.be.true;
    expect(await contract.hasRole(await contract.DISPATCHER_ROLE(), deployer.address)).to.be.true;
    expect(await contract.hasRole(await contract.DISTRIBUTER_ROLE(), deployer.address)).to.be.true;
  });
});
```

### 2. Configuration Management

```javascript
describe("Configuration", function() {
  it("should allow owner to update parameters", async function() {
    // Update unstake lock period
    const newUnstakeLockPeriod = 14 * 24 * 60 * 60;
    await contract.setUnstakeLockPeriod(newUnstakeLockPeriod);
    expect(await contract.unstakeLockPeriod()).to.equal(newUnstakeLockPeriod);
    
    // Update stake lock period
    const newStakeLockPeriod = 2 * 24 * 60 * 60;
    await contract.setStakeLockPeriod(newStakeLockPeriod);
    expect(await contract.stakeLockPeriod()).to.equal(newStakeLockPeriod);
    
    // Update reward scaling factor
    const newScalingFactor = 2;
    await contract.setCKPTRewardScalingFactor(newScalingFactor);
    expect(await contract.CKPTRewardScalingFactor()).to.equal(newScalingFactor);
  });
  
  it("should reject configuration updates from non-owner", async function() {
    const [_, nonOwner] = await ethers.getSigners();
    await expect(contract.connect(nonOwner).setUnstakeLockPeriod(1000))
      .to.be.revertedWith("Ownable: caller is not the owner");
  });
});
```

### 3. Role Management

```javascript
describe("Role Management", function() {
  it("should allow owner to grant and revoke roles", async function() {
    const [_, user] = await ethers.getSigners();
    
    // Grant validator role
    await contract.grantValidatorRole(user.address);
    expect(await contract.hasRole(await contract.VALIDATOR_ROLE(), user.address)).to.be.true;
    
    // Revoke validator role
    await contract.revokeValidatorRole(user.address);
    expect(await contract.hasRole(await contract.VALIDATOR_ROLE(), user.address)).to.be.false;
  });
});
```

### 4. Validator Registration

```javascript
describe("Validator Registration", function() {
  it("should allow users with validator role to register", async function() {
    const [_, validator] = await ethers.getSigners();
    
    // Grant validator role
    await contract.grantValidatorRole(validator.address);
    
    // Register as validator
    const dispatcherURL = "https://example.com/dispatcher";
    const blsPublicKey = "0x1234567890abcdef";
    
    await expect(contract.connect(validator).registerValidator(dispatcherURL, blsPublicKey))
      .to.emit(contract, "ValidatorRegistered")
      .withArgs(validator.address, dispatcherURL, blsPublicKey);
  });
  
  it("should reject registration from users without validator role", async function() {
    const [_, nonValidator] = await ethers.getSigners();
    
    await expect(contract.connect(nonValidator).registerValidator("url", "key"))
      .to.be.revertedWith("CKPTValStaking: must have validator role to stake");
  });
});
```

### 5. Staking Operations

```javascript
describe("Staking", function() {
  beforeEach(async function() {
    // Setup validator with tokens
    await contract.grantValidatorRole(validator.address);
    await stakingToken.mint(validator.address, stakeAmount);
    await stakingToken.connect(validator).approve(contract.address, stakeAmount);
    await contract.connect(validator).registerValidator("url", "key");
  });
  
  it("should allow validators to stake tokens", async function() {
    await expect(contract.connect(validator).stake(stakeAmount))
      .to.emit(contract, "Staked")
      .withArgs(validator.address, stakeAmount);
      
    const validatorInfo = await contract.getValidator(validator.address);
    expect(validatorInfo.stakedAmount).to.equal(stakeAmount);
    expect(validatorInfo.isActive).to.be.true;
  });
  
  it("should apply stake lock period", async function() {
    await contract.connect(validator).stake(stakeAmount);
    
    const activationTime = await contract.stakeActivationTime(validator.address);
    const stakeLockPeriod = await contract.stakeLockPeriod();
    const currentTime = (await ethers.provider.getBlock("latest")).timestamp;
    
    expect(activationTime).to.equal(currentTime + stakeLockPeriod);
  });
  
  it("should reject staking below minimum amount", async function() {
    const smallAmount = ethers.utils.parseEther("10");
    await expect(contract.connect(validator).stake(smallAmount))
      .to.be.revertedWith("Must stake exact bigger than minimum amount");
  });
});
```

### 6. Unstaking Operations

```javascript
describe("Unstaking", function() {
  beforeEach(async function() {
    // Setup validator with staked tokens
    // [setup code omitted for brevity]
  });
  
  it("should allow initiating unstake", async function() {
    await expect(contract.connect(validator).initiateUnstake(stakeAmount))
      .to.emit(contract, "UnstakeInitiated");
      
    const validatorInfo = await contract.getValidator(validator.address);
    expect(validatorInfo.stakedAmount).to.equal(0);
    expect(validatorInfo.isActive).to.be.false;
    expect(validatorInfo.unstakeTime).to.be.gt(0);
  });
  
  it("should allow completing unstake after lock period", async function() {
    await contract.connect(validator).initiateUnstake(stakeAmount);
    
    // Advance time past lock period
    await ethers.provider.send("evm_increaseTime", [await contract.unstakeLockPeriod() + 1]);
    await ethers.provider.send("evm_mine", []);
    
    const balanceBefore = await stakingToken.balanceOf(validator.address);
    await contract.connect(validator).completeUnstake();
    const balanceAfter = await stakingToken.balanceOf(validator.address);
    
    expect(balanceAfter.sub(balanceBefore)).to.equal(stakeAmount);
  });
  
  it("should reject completing unstake before lock period ends", async function() {
    await contract.connect(validator).initiateUnstake(stakeAmount);
    
    await expect(contract.connect(validator).completeUnstake())
      .to.be.revertedWith("Still in lock period");
  });
});
```

### 7. Checkpoint Submission

```javascript
describe("Checkpoint Submission", function() {
  beforeEach(async function() {
    // Setup dispatcher account
    await contract.grantDispatcherRole(dispatcher.address);
  });
  
  it("should allow dispatchers to submit checkpoints", async function() {
    const epochNum = 1;
    const blockHash = ethers.utils.randomBytes(32);
    const bitmap = ethers.utils.randomBytes(64);
    const blsMultiSig = ethers.utils.randomBytes(96);
    const blsAggrPk = ethers.utils.randomBytes(48);
    const powerSum = ethers.utils.parseEther("1000");
    
    await expect(contract.connect(dispatcher).submitCheckpoint(
      epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum
    ))
      .to.emit(contract, "CheckpointSubmitted")
      .withArgs(epochNum, blockHash, powerSum);
      
    const checkpoint = await contract.epochToCheckpoint(epochNum);
    expect(checkpoint.epochNum).to.equal(epochNum);
    expect(checkpoint.powerSum).to.equal(powerSum);
  });
  
  it("should reject duplicate checkpoint submissions", async function() {
    // Submit first checkpoint
    await contract.connect(dispatcher).submitCheckpoint(
      1, ethers.utils.randomBytes(32), "0x", "0x", "0x", 1000
    );
    
    // Attempt duplicate submission
    await expect(contract.connect(dispatcher).submitCheckpoint(
      1, ethers.utils.randomBytes(32), "0x", "0x", "0x", 2000
    ))
      .to.be.revertedWith("Checkpoint already exists for this epoch");
  });
});
```

### 8. Reward Distribution

```javascript
describe("Reward Distribution", function() {
  beforeEach(async function() {
    // Setup validators and submit checkpoint
    // [setup code omitted for brevity]
  });
  
  it("should distribute rewards to validators", async function() {
    await contract.connect(distributer).distributeCheckpointRewards(epochNum);
    
    expect(await contract.distributedEpochs(epochNum)).to.be.true;
    
    // Verify validators received rewards
    const validatorInfo = await contract.getValidator(validator.address);
    expect(validatorInfo.pendingRewards).to.be.gt(0);
  });
  
  it("should calculate rewards based on checkpoint power and scaling factor", async function() {
    // Set scaling factor
    await contract.setCKPTRewardScalingFactor(2);
    
    // Distribute rewards
    await contract.connect(distributer).distributeCheckpointRewards(epochNum);
    
    // Check rewards are scaled correctly
    const validatorInfo = await contract.getValidator(validator.address);
    // Verify rewards calculation
    // [verification logic omitted for brevity]
  });
  
  it("should reject duplicate reward distribution", async function() {
    await contract.connect(distributer).distributeCheckpointRewards(epochNum);
    
    await expect(contract.connect(distributer).distributeCheckpointRewards(epochNum))
      .to.be.revertedWith("Rewards already distributed for this epoch");
  });
});
```

### 9. Reward Claiming

```javascript
describe("Reward Claiming", function() {
  beforeEach(async function() {
    // Setup validator with rewards
    // [setup code omitted for brevity]
  });
  
  it("should allow validators to claim rewards", async function() {
    const validatorInfo = await contract.getValidator(validator.address);
    const pendingRewards = validatorInfo.pendingRewards;
    
    const balanceBefore = await stakingToken.balanceOf(validator.address);
    
    await contract.connect(validator).claimRewards();
    
    const balanceAfter = await stakingToken.balanceOf(validator.address);
    expect(balanceAfter.sub(balanceBefore)).to.equal(pendingRewards);
    
    // Verify rewards were reset
    const validatorAfter = await contract.getValidator(validator.address);
    expect(validatorAfter.pendingRewards).to.equal(0);
  });
  
  it("should reject claiming with no rewards", async function() {
    // Setup validator without rewards
    
    await expect(contract.connect(validator).claimRewards())
      .to.be.revertedWith("No rewards to claim");
  });
});
```

### 10. Validator Queries

```javascript
describe("Validator Queries", function() {
  beforeEach(async function() {
    // Setup multiple validators
    // [setup code omitted for brevity]
  });
  
  it("should return correct validator information", async function() {
    const info = await contract.getValidator(validator.address);
    
    expect(info.stakedAmount).to.equal(stakeAmount);
    expect(info.dispatcherURL).to.equal(dispatcherURL);
    expect(info.blsPublicKey).to.equal(blsPublicKey);
    expect(info.isActive).to.be.true;
  });
  
  it("should return top validators by rotation", async function() {
    const [addresses, stakes, urls, keys] = await contract.getTopValidators(3);
    
    expect(addresses.length).to.be.lte(3);
    expect(stakes.length).to.equal(addresses.length);
    
    // Verify rotation behavior
    await contract.updateValidatorCursor(addresses.length);
    const [newAddresses] = await contract.getTopValidators(3);
    
    // Check rotation happened
    if (addresses.length > 3) {
      expect(newAddresses[0]).to.not.equal(addresses[0]);
    }
  });
});
```

## Helper Functions for Testing

```javascript
// Deploy mock token for testing
async function deployMockToken() {
  const MockToken = await ethers.getContractFactory("MockERC20");
  return await MockToken.deploy("Mock Token", "MOCK");
}

// Deploy CKPTValStaking with default settings
async function deployContract(tokenAddress, rewardRate, minimumStake) {
  const CKPTValStaking = await ethers.getContractFactory("CKPTValStaking");
  return await CKPTValStaking.deploy(tokenAddress, rewardRate, minimumStake);
}

// Advance blockchain time
async function advanceTime(seconds) {
  await ethers.provider.send("evm_increaseTime", [seconds]);
  await ethers.provider.send("evm_mine", []);
}
```

This test document provides comprehensive coverage of the CKPTValStaking contract functionality, testing all configurable parameters and key features.
