const { ethers } = require("hardhat");

async function main() {
    let MyToken, HGT_Token, deployer, addr;
    const NAME = "HETU";
    const SYMBOL = "HETU";
    const DECIMALS = 18;

    // deploy token contract
    MyToken = await ethers.getContractFactory("ERC20MinterBurnerDecimals");
    HGT_Token = await MyToken.deploy(NAME, SYMBOL, DECIMALS);
    addr = await HGT_Token.getAddress()

    console.log(`HGT token deployed to: ${addr}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
