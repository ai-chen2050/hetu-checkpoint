const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("CKPTValStaking Contract", function () {
    let ValStaking, valStaking, MyToken, stakingToken, owner, addr1, addr2, addr3, valStakingAddress;
    const DECIMALS = 18;

    beforeEach(async function () {
        // Deploy a HGT ERC20 token for staking
        const NAME = "HETU";
        const SYMBOL = "HETU";
        const DECIMALS = 18;

        // deploy token contract
        MyToken = await ethers.getContractFactory("ERC20MinterBurnerDecimals");
        stakingToken = await MyToken.deploy(NAME, SYMBOL, DECIMALS);
        const stakingTokenAddress = await stakingToken.getAddress()
        console.log(stakingTokenAddress);

        // Get the contract factory and signers
        ValStaking = await ethers.getContractFactory("CKPTValStaking");
        [owner, addr1, addr2, addr3, _] = await ethers.getSigners();

        // Deploy the ValStaking contract
        valStaking = await ValStaking.deploy(stakingTokenAddress, ethers.parseUnits("1", DECIMALS), ethers.parseUnits("100", DECIMALS));
        valStakingAddress = await valStaking.getAddress();
        console.log(valStakingAddress);

        // Grant VALIDATOR_ROLE to test accounts
        await valStaking.grantValidatorRole(addr1.address);
        await valStaking.grantValidatorRole(addr2.address);
        await valStaking.grantValidatorRole(addr3.address);

        // Grant DISPATCHER_ROLE and DISTRIBUTER_ROLE to owner
        await valStaking.grantDispatcherRole(owner.address);
        await valStaking.grantDistributerRole(owner.address);

        // Set stake lock period to 0 for testing
        await valStaking.setStakeLockPeriod(0);
    });

    it("Should set the correct owner", async function () {
        expect(await valStaking.owner()).to.equal(owner.address);
    });

    it("Should set the minimum stake correctly", async function () {
        await valStaking.setMinimumStake(ethers.parseUnits("200", DECIMALS));
        expect(await valStaking.minimumStake()).to.equal(ethers.parseUnits("200", DECIMALS));
    });

    it("Should allow staking with the exact minimum amount", async function () {
        await stakingToken.transfer(addr1.address, ethers.parseUnits("100", DECIMALS));
        await stakingToken.connect(addr1).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));

        await expect(valStaking.connect(addr1).stake(ethers.parseUnits("100", DECIMALS)))
            .to.emit(valStaking, "Staked")
            .withArgs(addr1.address, ethers.parseUnits("100", DECIMALS));

        const validator = await valStaking.validators(addr1.address);
        expect(validator.stakedAmount).to.equal(ethers.parseUnits("100", DECIMALS));
        expect(validator.index).to.equal(0);
    });

    it("Should return top validators correctly", async function () {
        const valnum = 2;
        await stakingToken.transfer(addr1.address, ethers.parseUnits("100", DECIMALS));
        await stakingToken.transfer(addr2.address, ethers.parseUnits("100", DECIMALS));
        await stakingToken.transfer(addr3.address, ethers.parseUnits("100", DECIMALS));
        await stakingToken.connect(addr1).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));
        await stakingToken.connect(addr2).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));
        await stakingToken.connect(addr3).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));

        await valStaking.connect(addr1).stake(ethers.parseUnits("100", DECIMALS));
        await valStaking.connect(addr2).stake(ethers.parseUnits("100", DECIMALS));
        await valStaking.connect(addr3).stake(ethers.parseUnits("100", DECIMALS));

        await valStaking.connect(addr1).updateValidatorInfo("https://validator1.com", "blsPublicKey1");
        await valStaking.connect(addr2).updateValidatorInfo("https://validator2.com", "blsPublicKey2");
        await valStaking.connect(addr3).updateValidatorInfo("https://validator3.com", "blsPublicKey3");

        const [addresses, stakes, dispatcherURLs, blsPublicKeys] = await valStaking.getTopValidators(valnum);
        console.log("Addresses:", addresses);
        console.log("Stakes:", stakes);
        console.log("Dispatcher URLs:", dispatcherURLs);
        console.log("BLS Public Keys:", blsPublicKeys);

        // Update the validator cursor
        await valStaking.updateValidatorCursor(valnum);

        expect(addresses).to.include(addr1.address);
        expect(addresses).to.include(addr2.address);

        const [addresses2, stakes2, dispatcherURLs2, blsPublicKeys2] = await valStaking.getTopValidators(valnum);
        console.log("Addresses:", addresses2);
        console.log("Stakes:", stakes2);
        console.log("Dispatcher URLs:", dispatcherURLs2);
        console.log("BLS Public Keys:", blsPublicKeys2);
        expect(addresses2).to.include(addr3.address);
    });

    it("Should allow dispatcher to submit a checkpoint", async function () {
        const epochNum = 1;
        const blockHash = ethers.keccak256(ethers.toUtf8Bytes("blockHash"));
        const bitmap = "0x01";
        const blsMultiSig = "0x1234";
        const blsAggrPk = "0x5678";
        const powerSum = 1000;

        await expect(
            valStaking.submitCheckpoint(epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum)
        ).to.emit(valStaking, "CheckpointSubmitted")
            .withArgs(epochNum, blockHash, powerSum);
    });

    it("Should not allow non-dispatcher to submit a checkpoint", async function () {
        const epochNum = 1;
        const blockHash = ethers.keccak256(ethers.toUtf8Bytes("blockHash"));
        const bitmap = "0x01";
        const blsMultiSig = "0x1234";
        const blsAggrPk = "0x5678";
        const powerSum = 1000;

        await expect(
            valStaking.connect(addr1).submitCheckpoint(epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum)
        ).to.be.revertedWith("CKPTValStaking: must have dispatcher role to submit checkpoint");
    });

    it("Should allow distributer to distribute checkpoint rewards", async function () {
        const epochNum = 1;
        const blockHash = ethers.keccak256(ethers.toUtf8Bytes("blockHash"));
        const bitmap = "0x11";
        const blsMultiSig = "0x1234";
        const blsAggrPk = "0x5678";
        const powerSum = 100;

        // Submit a checkpoint
        await valStaking.submitCheckpoint(epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum);

        // Add a validator to ensure rewards can be distributed
        await stakingToken.transfer(addr1.address, ethers.parseUnits("100", DECIMALS));
        await stakingToken.connect(addr1).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));
        await valStaking.connect(addr1).stake(ethers.parseUnits("100", DECIMALS));

        let before = await valStaking.getValidator(addr1);
        console.log("Distr Before: ", before[1]);
        // Distribute rewards
        await expect(valStaking.distributeCheckpointRewards(epochNum))
            .to.not.be.reverted;

        let after = await valStaking.getValidator(addr1);
        console.log("Distr After:  ", after[1]);
        // Verify the epoch is marked as distributed
        expect(await valStaking.distributedEpochs(epochNum)).to.be.true;
    });

    it("Should not allow non-distributer to distribute checkpoint rewards", async function () {
        const epochNum = 1;
        const blockHash = ethers.keccak256(ethers.toUtf8Bytes("blockHash"));
        const bitmap = "0x01";
        const blsMultiSig = "0x1234";
        const blsAggrPk = "0x5678";
        const powerSum = 1000;

        // Submit a checkpoint
        await valStaking.submitCheckpoint(epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum);

        // Attempt to distribute rewards as a non-distributer
        await expect(
            valStaking.connect(addr1).distributeCheckpointRewards(epochNum)
        ).to.be.revertedWith("CKPTValStaking: must have distributer role to distribute rewards");
    });

    it("Should not allow distributing rewards for the same epoch twice", async function () {
        const epochNum = 1;
        const blockHash = ethers.keccak256(ethers.toUtf8Bytes("blockHash"));
        const bitmap = "0x01";
        const blsMultiSig = "0x1234";
        const blsAggrPk = "0x5678";
        const powerSum = 1000;

        // Submit a checkpoint
        await valStaking.submitCheckpoint(epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum);

        // Add a validator to ensure rewards can be distributed
        await stakingToken.transfer(addr1.address, ethers.parseUnits("100", DECIMALS));
        await stakingToken.connect(addr1).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));
        await valStaking.connect(addr1).stake(ethers.parseUnits("100", DECIMALS));

        // Distribute rewards
        await valStaking.distributeCheckpointRewards(epochNum);

        // Attempt to distribute rewards again
        await expect(valStaking.distributeCheckpointRewards(epochNum))
            .to.be.revertedWith("Rewards already distributed for this epoch");
    });

    it("Should distribute checkpoint rewards for 20 validators", async function () {
        // Get available signers
        const signers = await ethers.getSigners();
        // Use only the available signers (typically 20 in Hardhat)
        const validatorCount = signers.length;
        console.log(`Available validators: ${validatorCount}`);
        const validators = signers;

        // Grant validator roles and stake tokens for all validators
        for (let i = 0; i < validators.length; i++) {
            const validator = validators[i];

            // Grant validator role if not already assigned
            await valStaking.grantValidatorRole(validator.address);

            // Transfer and approve tokens
            await stakingToken.transfer(validator.address, ethers.parseUnits("100", DECIMALS));
            await stakingToken.connect(validator).approve(valStakingAddress, ethers.parseUnits("100", DECIMALS));

            // Stake tokens
            await valStaking.connect(validator).stake(ethers.parseUnits("100", DECIMALS));

            // Update validator info with unique values
            await valStaking.connect(validator).updateValidatorInfo(
                `https://validator${i}.com`,
                `blsPublicKey${i}`
            );
        }

        const stake19 = await valStaking.getValidator(validators[19]);
        console.log("Stake Validators 19: ", stake19[4]);

        // Set up checkpoint data
        const epochNum = 2;
        const blockHash = ethers.keccak256(ethers.toUtf8Bytes("blockHash512"));
        const bitmap = "0x" + "f".repeat(128); // Full bitmap with all validators
        const blsMultiSig = "0x1234";
        const blsAggrPk = "0x5678";
        const powerSum = validatorCount * 100;

        // Submit the checkpoint
        await valStaking.submitCheckpoint(epochNum, blockHash, bitmap, blsMultiSig, blsAggrPk, powerSum);
        console.log("Checkpoint submitted");

        // Sample a few validators to check rewards before distribution
        const sampleIndices = [0, 2, 8, 19];
        const beforeRewards = {};

        for (const idx of sampleIndices) {
            const validator = validators[idx];
            const validatorInfo = await valStaking.getValidator(validator);
            beforeRewards[idx] = validatorInfo[1]; // Store pending rewards
            console.log(`Validator ${idx} before: ${validatorInfo[1]}`);
        }

        // Distribute rewards
        await expect(valStaking.distributeCheckpointRewards(epochNum))
            .to.not.be.reverted;

        // Check rewards after distribution
        for (const idx of sampleIndices) {
            const validator = validators[idx];
            const validatorInfo = await valStaking.getValidator(validator);
            console.log(`Validator ${idx} after: ${validatorInfo[1]}`);
        }

        // Verify the epoch is marked as distributed
        expect(await valStaking.distributedEpochs(epochNum)).to.be.true;
    });
});