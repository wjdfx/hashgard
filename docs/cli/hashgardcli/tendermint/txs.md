# hashgardcli tendermint txs

## Description

Search for transactions that match the exact given tags where results are paginated.

## Usage

```
hashgardcli tendermint txs [flags]
```

## Flags

| Name，shorthand | type  | Required|Default| Description   |
| ---------- | ------ | ---- | ------ | ------------------------- |
| --limit    | int    | No  | 32     | Query number of transactions results per page returned     |
| --page     | int    | No  | 1      |  Query a specific page of paginated results|
| --tags     | string | Yes  |        | tag:value list of tags that must match|

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
 hashgardcli tendermint txs --tags '<tag1>:<value1>&<tag2>:<value2>' \
 --page 1 --limit 30 --trust-node
```

