# hashgard export

## Description

hashgard can export blockchain state at any height and output json format string.

## Usage

```shell
hashgard export [flags]
```

## Flags

| Name，shorthand      | type  | Default| description                                 | Required  |
| ----------------- | ------ | ------ | ------------------------------------------- | -------- |
| -h, --help        |        |        | help for export                          | No   |
| --for-zero-height |        |        | Export state to start at height zero   | No   |
| --height          | int    | -1     | Export state from a particular height   | No  |
| --jail-whitelist  | string |        | List of validators to not jail state export| No  |

## Example

`hashgard export`
