# Protocol Buffer Definitions

This directory contains the Protocol Buffer definitions for the Hetu Checkpoint project.

## Overview

The protocol buffer definitions define the data structures and service interfaces used for:

- Checkpoint data structures
- Hetu-chain common messages

## Key Files

- `checkpoint.proto`: Core checkpoint data structures and messages
- [common](./common/): Hetu chain common messages

## Building

The protobuf files are compiled using:

- protoc compiler
- gogo protobuf plugin
- grpc-gateway plugin

See the Makefile for the exact build commands.