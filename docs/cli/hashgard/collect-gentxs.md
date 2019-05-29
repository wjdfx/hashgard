# hashgard collect-gentxs


## Description
Collect genesis txs and output a genesis.json file

## Usage
```
hashgard collect-gentxs [flags]
```

## Flags
| Name，shorthand| type  | Default                   | description                   | Required |
| ----------- | ------ | ------------------------- | ------------------------------ | -------- |
| --gentx-dir | string | ~/.hashgard/config/gentx/ |  override default "gentx" directory from which collect and execute genesis transactions| No  |
| -h, --help  |        |                           |  help for collect-gentxs                    | No  |
| --home      | string | ~/.hashgard               |  directory for config and data              | No  |
| --trace     | bool   |                           | print out full stack trace on erro          | No  |

## Example
`hashgard collect-gentxs`