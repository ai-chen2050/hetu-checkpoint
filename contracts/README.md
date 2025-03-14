# Contracts

This directory contains smart contract-related files and artifacts.

## Artifacts

The following artifacts are generated during contract compilation:

- **ABI** (Application Binary Interface): JSON files that define how to interact with the contracts
- **BIN** (Binary): Compiled bytecode of the contracts

## Smart Contracts

The following smart contracts are implemented:

- **CKPT_Val_Staking**: Validator staking contract for managing validator stakes and rewards
- **ERC20MinterBurnerDecimals**: Custom ERC20 token implementation with:
  - Minting capability
  - Burning functionality
  - Configurable decimals

## Development

To work with these contracts:

1. Compile the contracts to generate artifacts
2. Use the generated ABI files for contract interaction
3. Deploy using the compiled bytecode (BIN files)

For detailed implementation and usage, refer to the individual contract files.