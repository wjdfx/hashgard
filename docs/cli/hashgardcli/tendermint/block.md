# hashgardcli tendermint block

## Description

Get verified data for a the block at given height 

## Usage

```
  hashgardcli tendermint block [height] [flags]
```

## Flags

| Name, shorthand  | Default               | Description            | Required               |
| ------------ | --------------------- | -------------------------- | ---------------------- |
| --chain-id   | No                    | Chain ID of Tendermint node                     | No|
| --node       | tcp://localhost:26657 | Node to connect to                      | No        |
| --trust-node | true        | Trust connected full node (don't verify proofs for responses)| No |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli tendermint block 114263  --trust-node
```
