# hashgard init

## Description

Initialize validators's and node's configuration files.

## Usage

```bash
hashgard init [flags]
```

## Flags

| Name，shorthand| type  | Default     | description                                                  | Required  |
| ----------- | ------ | ----------- | ------------------------------------------------------------ | -------- |
| -h, --help  |        |             | help for init                                           | No  |
| --chain-id  | string |             | genesis file chain-id, if left blank will be randomly created    | N0  |
| --moniker   | string |             | set the validator's moniker | Yes     |
| --overwrite | bool   |             | overwrite the genesis.json file                                      | No   |
| --home      | string | ~/.hashgard | directory for config and data                                          | No   |
| --trace     | bool   |             |  print out full stack trace on errors                                   | No  |

## Example

`hashgard init --chain-id=testnet-1000 --moniker=hashgard`
