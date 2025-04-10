const { ethers } = require("hardhat");
const { fromBech32 } = require('@cosmjs/encoding');
const fs = require('fs');
const path = require('path');

let dispatcherAddrs = [];
let validatorAddrs = [];

// Helper function to execute transactions with retry logic
async function executeWithRetry(txFunc, maxRetries = 3, message = "") {
    let retries = 0;
    while (retries < maxRetries) {
        try {
            const tx = await txFunc();
            const receipt = await tx.wait();
            console.log(`${message} - Success (TX: ${receipt.hash})`);
            return receipt;
        } catch (error) {
            retries++;
            console.log(`Attempt ${retries}/${maxRetries} failed for ${message}: ${error.message}`);
            if (retries >= maxRetries) {
                throw error;
            }
            // Exponential backoff: wait longer between retries
            await new Promise(r => setTimeout(r, 5000 * retries));
        }
    }
}

// Function to read and process JSON files
async function processJsonFiles(directoryPath, stakingTokenAddress) {
    try {
        // Read all files in the directory
        const files = fs.readdirSync(directoryPath);

        // Filter for JSON files
        const jsonFiles = files.filter(file => file.endsWith('.json'));

        // Process each JSON file
        for (const file of jsonFiles) {
            const filePath = path.join(directoryPath, file);
            const fileContent = fs.readFileSync(filePath, 'utf8');
            const jsonData = JSON.parse(fileContent);

            // Check if the file contains eth.address
            if (jsonData.eth && jsonData.eth.address) {
                const targetAddress = jsonData.eth.address;
                console.log(`Processing address from ${file}: ${targetAddress}`);

                // 根据文件名判断是validator还是dispatcher
                if (file.includes("validator")) {
                    validatorAddrs.push(jsonData.eth.address);
                } else {
                    dispatcherAddrs.push(jsonData.eth.address);
                }

                // Check ETH balance before transferring
                const [deployer] = await ethers.getSigners();
                const ethBalance = await deployer.provider.getBalance(targetAddress);
                const requiredEth = ethers.parseEther("1000");

                if (ethBalance < Number(requiredEth) / 2) {
                    console.log(`ETH balance (${ethers.formatEther(ethBalance)} ETH) is less than required 1000 ETH, transferring...`);
                    // Transfer native ETH with retry
                    await executeWithRetry(
                        () => deployer.sendTransaction({
                            to: targetAddress,
                            value: requiredEth,
                            gasLimit: 100000 // Add explicit gas limit
                        }),
                        3,
                        `Transferring 1000 ETH to ${targetAddress}`
                    );
                } else {
                    console.log(`ETH balance (${ethers.formatEther(ethBalance)} ETH) is sufficient, skipping transfer`);
                }

                // Check ERC20 balance before transferring
                const StakingToken = await ethers.getContractAt("ERC20MinterBurnerDecimals", stakingTokenAddress);
                const tokenBalance = await StakingToken.balanceOf(targetAddress);
                const requiredTokens = ethers.parseUnits("1000", 18);

                if (tokenBalance < Number(requiredTokens) / 2) {
                    console.log(`ERC20 balance (${ethers.formatUnits(tokenBalance, 18)} tokens) is less than required 1000 tokens, transferring...`);
                    // Transfer ERC20 tokens with retry
                    await executeWithRetry(
                        () => StakingToken.transfer(
                            targetAddress,
                            requiredTokens,
                            { gasLimit: 200000 } // Add explicit gas limit
                        ),
                        3,
                        `Transferring 1000 ERC20 tokens to ${targetAddress}`
                    );
                } else {
                    console.log(`ERC20 balance (${ethers.formatUnits(tokenBalance, 18)} tokens) is sufficient, skipping transfer`);
                }
            }
        }
    } catch (error) {
        console.error("Error processing JSON files:", error);
    }
}

async function main() {
    // Process JSON files keystores in the specified directory
    const jsonDirectory = "/Users/home/hetu-checkpoint/keys"; // Replace with your actual directory path

    const [deployer] = await ethers.getSigners();
    const privateKey = "C4B7A5B244AA818369FCBAC7A4D717066B794DD360666A49C7F1CD2FAC0BB9D4"; // replace with your private key
    const wallet = new ethers.Wallet(privateKey, deployer.provider);

    const valStakingAddress = "0xE7F6C60D813a17e2FA16869371351d8C800c7fC1"; // replace with ValStaking contract address
    const ValStaking = await ethers.getContractAt("CKPTValStaking", valStakingAddress, wallet);

    const stakingTokenAddress = "0x9E96fc769eFb3D2f681D4D4157aac04e97DC9953"; // replace with ERC20 staking token address
    const StakingToken = await ethers.getContractAt("ERC20MinterBurnerDecimals", stakingTokenAddress, wallet);

    const DECIMALS = 18;
    const amount = ethers.parseUnits("180", DECIMALS);

    // Check wallet balance
    const walletBalance = await StakingToken.balanceOf(wallet.address);
    console.log(`Wallet Balance: ${ethers.formatUnits(walletBalance, DECIMALS)} HET`);

    await processJsonFiles(jsonDirectory, stakingTokenAddress);

    // Set stake lock period to 0 for testing
    console.log("Set stake lock period to 0");
    await executeWithRetry(
        () => ValStaking.setStakeLockPeriod(0, { gasLimit: 200000 }),
        3,
        "Setting stake lock period to 0"
    );

    // Grant validator role to validator addresses
    console.log("Grant validator role to validator addresses");
    for (const validatorAddr of validatorAddrs) {
        await executeWithRetry(
            () => ValStaking.grantValidatorRole(validatorAddr, { gasLimit: 200000 }),
            3,
            `Granting validator role to ${validatorAddr}`
        );
    }

    // Grant dispatcher role to dispatcher addresses
    console.log("Grant dispatcher role to dispatcher addresses");
    for (const dispatcherAddr of dispatcherAddrs) {
        await executeWithRetry(
            () => ValStaking.grantDispatcherRole(dispatcherAddr, { gasLimit: 200000 }),
            3,
            `Granting dispatcher role to ${dispatcherAddr}`
        );

        // Grant distributer role to dispatcher addresses
        console.log("Grant distributer role to dispatcher addresses");
        await executeWithRetry(
            () => ValStaking.grantDistributerRole(dispatcherAddr, { gasLimit: 200000 }),
            3,
            `Granting distributer role to ${dispatcherAddr}`
        );
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });