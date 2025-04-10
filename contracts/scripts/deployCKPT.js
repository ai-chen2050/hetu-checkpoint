const { ethers } = require("hardhat");

async function main() {
    let CKPTStaking, ckpt_staking_contract, deployer, addr;
    const Token = "0x9E96fc769eFb3D2f681D4D4157aac04e97DC9953";
    const rate = 1;
    const limit = 18;

    // deploy token contract
    CKPTStaking = await ethers.getContractFactory("CKPTValStaking");
    ckpt_staking_contract = await CKPTStaking.deploy(Token, rate, limit);
    addr = await ckpt_staking_contract.getAddress()

    console.log(`CKPTStaking deployed to: ${addr}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
