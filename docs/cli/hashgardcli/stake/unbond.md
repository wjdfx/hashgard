# hashgardcli stake unbond

## Description 

Unbond an amount of bonded shares from a validator:

## Usage

```
hashgardcli stake unbond [validator-addr] [amount] [flags]
```

## Available Commands（Subcommands）

|     Name      | type  | Required| Default| description         |
| -------------- | ------ | -------- | ------ | ------------------- |
| validator-addr | string | `true`     |        | Bech address of the validator |
| amount         | int    | `true`     |        | Amount of source-shares to either unbond or redelegate as a positive integer or decimal|

## Flags

| Name   | type  | Required| Default| description          |
| ------ | ------ | -------- | ------ | -------------------- |
| --from | string | `true`     | ""     | Delegators account name or address|

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli stake unbond \
gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx \
5000 \
--from=hashgard \
--chain-id=hashgard
```
