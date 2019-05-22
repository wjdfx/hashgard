# hashgard gentx

## Description

This command is an alias of the 'hashgard tx create-validator' command'.

It creates a genesis piece carrying a self delegation with the
following delegation and commission default parameters:

```
delegation amount:           100000000000000000000agard
	commission rate:             0.1
	commission max rate:         0.2
	commission max change rate:  0.01
	minimum self delegation:     1
```

## Usage

```shell
 hashgard gentx [flags]
```

## Flags

| Name，shorthand               | type  | Default        | description                                         | Required  |
| ---------------------------- | ------ | -------------- | ------------------------------------- | -------- |
| --amount                     | string |                |  Amount of coins to bond                | `Yes`     |
| --commission-max-change-rate | string |                | The maximum commission change rate percentage (per day) | No  |
| --commission-max-rate        | string |                | The maximum commission rate percentage           | No  |
| --commission-rate            | string |                | The initial commission rate percentage           | No  |
| --home-client                | string | ~/.hashgardcli | client's home directory                          | No  |
| --ip                         | string | localhost IP   | The node's public IP                             | No  |
| --min-self-delegation        | string |                | The minimum self delegation required on the validator  | `true`     |
| --name                       | string |                | name of private key with which to sign the gentx   | No  |
| --node-id                    | string |            | The node's NodeID                                  | No  |
| --output-document            | string |                | write the genesis transaction JSON document to the given file instead of the default location| No  |
| --pubkey                     | string |                | The Bech32 encoded PubKey of the validator                   | No  |
| -h, --help                   |        |                | help for gentx                               | No  |
| --home                       | string | ~/.hashgard    | directory for config and data                     | No  |
| --trace                      | bool   |                | print out full stack trace on error         | No  |

## Example

`hashgard gentx --name=root --amount=10000gard --ip=${validator_ip}`
