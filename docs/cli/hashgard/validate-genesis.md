# hashgard validate-genesis

## Description

validates the genesis file at the default location or at the location passed as an arg

## Usage

```
hashgard validate-genesis [file] [flags]
```

## Available Commands

| Name, shorthand|type  | Default                         | Description        | Required  |
| ---------- | ------ | ------------------------------- | ---------------- | -------- |
| [file]     | string | ~/.hashgard/config/genesis.json | genesis 文件位置 | No  |

## Flags

| Name, shorthand|type  | Default     | Description                        | Required  |
| ---------- | ------ | ----------- | -------------------------------- | -------- |
| -h, --help |        |             | help for validate-genesis | No  |
| --home     | string | ~/.hashgard | directory for config and data                | No  |
| --trace    | bool   |             | print out full stack trace on errors         | No  |

## Example

```bash
hashgard validate-genesis
```
