# Hetu Checkpoint

Hetu Checkpoint is a secure and efficient checkpointing system that leverages Byzantine Consistent Broadcast (BCB) and BLS signatures to prevent long-range attacks on blockchain networks.

## Overview

Hetu Checkpoint creates secure checkpoints for multiple blocks using a network of validators. These checkpoints are then anchored to Ethereum, providing an additional layer of security and finality.

## Key Features

- **Epoch-based Checkpointing**: Divides blockchain into epochs with fixed validator sets
- **BCB Integration**: Secure Byzantine Consistent Broadcast for validator coordination
- **BLS Signatures**: Efficient aggregated signatures for checkpoint creation
- **Ethereum Anchoring**: Checkpoint hashes are recorded on Ethereum
- **Long-range Attack Prevention**: Enhanced security through periodic checkpointing

## Architecture

### Epoch Structure
```ascii:/README.md
+--------+--------+--------+--------+
|Block 1 |Block 2 |  ...   |Block N | => Checkpoint Hash => BLS Signatures
+--------+--------+--------+--------+                             |
               |                                                  V
        Hetu Chain Epoch n                                        L1
```            
## System Components
### Validator Network

Distributed network of validators
Consensus on checkpoint generation
BLS signature aggregation
### BCB Algorithm

Efficient consistency Broadcast
Validator coordination
Threshold signature scheme
### Checkpoint Generation

Aggregates multiple blocks
Creates unified checkpoint hash
BLS signature verification
### Ethereum Bridge

Submits checkpoint hashes to Ethereum
Provides immutable checkpoint records
Enables cross-chain verification

## Technical Details
### Epoch Configuration
Fixed number of blocks per epoch
Consistent validator set within each epoch
Configurable epoch parameters
### Security Features
- BLS Signatures: Efficient signature aggregation Reduced communication overhead Strong cryptographic security
- **BCB Protocol**: Byzantine Consistent Broadcast ensures that all honest validators agree on the same checkpoint data, providing a distributed trust model with no single point of failure and threshold-based security.


## Contributing
We welcome contributions! Please see our Contributing Guidelines for details.

## Security

## License
See the LICENSE file for details.
