# hashgard tendermint

## Description

Tendermint subcommands

## Usage

```shell
hashgard tendermint [subcommand] [flags]
```

## Subcommands

| Commands           | Description                            |
| ---------------- | ------------------------------------ |
| --show-node-id   | Show this node's ID                   |
| --show-validator | Show this node's tendermint validator info |
| --show-address   | Shows this node's tendermint validator consensus address |

## Flags

| Name, shorthand|type  | Default     | description              | Required  |
| ---------- | ------ | ----------- | ------------------------ | -------- |
| -h, --help |        |             | help for tendermint       | No  |
| --home     | string | ~/.hashgard | directory for config and data         | No  |
| --trace    | bool   |             | print out full stack trace on errors | No  |

## Example

```shell
hashgard tendermint show-node-id
hashgard tendermint show-validator
hashgard tendermint show-address
```
