const { expect } = require("chai");
const { describe, it, beforeEach } = require("mocha");
const { ethers } = require("hardhat");

describe("ERC20MinterBurnerDecimals", function () {
  let Token, HGT_Token, deployer, addr1, addr2;
  const NAME = "HETU";
  const SYMBOL = "HETU";
  const DECIMALS = 18;
  const INITIAL_SUPPLY = ethers.parseUnits("50000000", DECIMALS); // 50M
  const MAX_SUPPLY = ethers.parseUnits("4200000000", DECIMALS); // 4.2B

  beforeEach(async function () {
    // get test accounts
    [deployer, addr1, addr2] = await ethers.getSigners();

    // deploy token contract
    Token = await ethers.getContractFactory("ERC20MinterBurnerDecimals");
    HGT_Token = await Token.deploy(NAME, SYMBOL, DECIMALS);
    console.log(await HGT_Token.getAddress());
  });

  describe("Deployment", function () {
    it("Should allocate initial supply to deployer", async function () {
      const deployerBalance = await HGT_Token.balanceOf(deployer.address);
      expect(deployerBalance).to.equal(INITIAL_SUPPLY);
      expect(await HGT_Token.totalSupply()).to.equal(INITIAL_SUPPLY);
    });

    it("Should set correct name, symbol, and decimals", async function () {
      expect(await HGT_Token.name()).to.equal(NAME);
      expect(await HGT_Token.symbol()).to.equal(SYMBOL);
      expect(await HGT_Token.decimals()).to.equal(DECIMALS);
    });

    it("Should assign all roles to deployer", async function () {
      const DEFAULT_ADMIN_ROLE = await HGT_Token.DEFAULT_ADMIN_ROLE();
      const MINTER_ROLE = await HGT_Token.MINTER_ROLE();
      const PAUSER_ROLE = await HGT_Token.PAUSER_ROLE();
      const BURNER_ROLE = await HGT_Token.BURNER_ROLE();

      expect(await HGT_Token.hasRole(DEFAULT_ADMIN_ROLE, deployer.address)).to.be.true;
      expect(await HGT_Token.hasRole(MINTER_ROLE, deployer.address)).to.be.true;
      expect(await HGT_Token.hasRole(PAUSER_ROLE, deployer.address)).to.be.true;
      expect(await HGT_Token.hasRole(BURNER_ROLE, deployer.address)).to.be.true;

      const minterMembers = await HGT_Token.getRoleMembers(MINTER_ROLE);
      expect(minterMembers).to.include(deployer.address);
      expect(await HGT_Token.getRoleMemberCount(MINTER_ROLE)).to.equal(1);
    });
  });

  describe("Minting", function () {
    it("Should allow minter to mint tokens within max supply", async function () {
      const mintAmount = ethers.parseUnits("1000000", DECIMALS); // 1M
      await HGT_Token.mint(addr1.address, mintAmount);
      const have_mint = INITIAL_SUPPLY + mintAmount;
      expect(await HGT_Token.balanceOf(addr1.address)).to.equal(mintAmount);
      expect(await HGT_Token.totalSupply()).to.equal(have_mint);
    });

    it("Should fail to mint beyond max supply", async function () {
      const remainingSupply = MAX_SUPPLY - INITIAL_SUPPLY; // 4.15B
      const excessAmount = remainingSupply + ethers.parseUnits("1", DECIMALS);
      await expect(HGT_Token.mint(addr1.address, excessAmount)).to.be.revertedWith(
        "ERC20MinterBurnerDecimals: mint would exceed max supply"
      );
    });

    it("Should fail if non-minter tries to mint", async function () {
      const mintAmount = ethers.parseUnits("1000", DECIMALS);
      await expect(HGT_Token.connect(addr1).mint(addr2.address, mintAmount)).to.be.revertedWith(
        "ERC20MinterBurnerDecimals: must have minter role to mint"
      );
    });
  });

  describe("Role Enumeration", function () {
    it("Should add and track new minter role members", async function () {
      const MINTER_ROLE = await HGT_Token.MINTER_ROLE();
      await HGT_Token.grantRole(MINTER_ROLE, addr1.address);
      const members = await HGT_Token.getRoleMembers(MINTER_ROLE);
      expect(members).to.have.lengthOf(2);
      expect(members).to.include(addr1.address);
      expect(await HGT_Token.getRoleMemberCount(MINTER_ROLE)).to.equal(2);
    });

    it("Should remove minter role members", async function () {
      const MINTER_ROLE = await HGT_Token.MINTER_ROLE();
      await HGT_Token.grantRole(MINTER_ROLE, addr1.address);
      await HGT_Token.revokeRole(MINTER_ROLE, addr1.address);
      const members = await HGT_Token.getRoleMembers(MINTER_ROLE);
      expect(members).to.have.lengthOf(1);
      expect(members).to.not.include(addr1.address);
      expect(members).to.include(deployer.address);
    });
  });

  describe("Pausing", function () {
    it("Should allow pauser to pause and unpause", async function () {
      await HGT_Token.pause();
      expect(await HGT_Token.paused()).to.be.true;
      await expect(HGT_Token.transfer(addr1.address, ethers.parseUnits("1000", DECIMALS))).to.be.revertedWith(
        "ERC20Pausable: token transfer while paused"
      );

      await HGT_Token.unpause();
      expect(await HGT_Token.paused()).to.be.false;
      await HGT_Token.transfer(addr1.address, ethers.parseUnits("1000", DECIMALS));
      expect(await HGT_Token.balanceOf(addr1.address)).to.equal(ethers.parseUnits("1000", DECIMALS));
    });

    it("Should fail if non-pauser tries to pause", async function () {
      await expect(HGT_Token.connect(addr1).pause()).to.be.revertedWith(
        "ERC20MinterBurnerDecimals: must have pauser role to pause"
      );
    });
  });

  describe("Burning", function () {
    it("Should allow burner to burn tokens", async function () {
      const burnAmount = ethers.parseUnits("1000", DECIMALS);
      await HGT_Token.transfer(addr1.address, burnAmount);
      await HGT_Token.burnCoins(addr1.address, burnAmount);
      const have_burn = INITIAL_SUPPLY - burnAmount;
      expect(await HGT_Token.balanceOf(addr1.address)).to.equal(0);
      expect(await HGT_Token.totalSupply()).to.equal(have_burn);
    });

    it("Should fail if non-burner tries to burn", async function () {
      const burnAmount = ethers.parseUnits("1000", DECIMALS);
      await HGT_Token.transfer(addr1.address, burnAmount);
      await expect(HGT_Token.connect(addr2).burnCoins(addr1.address, burnAmount)).to.be.revertedWith(
        "ERC20MinterBurnerDecimals: must have burner role to burn"
      );
    });
  });
});