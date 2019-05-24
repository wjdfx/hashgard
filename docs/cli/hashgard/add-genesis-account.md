# hashgard add-genesis-account

## Description
add genesis account to /path/to/.hashgard/config/genesis.json 


## Usage
```
hashgard add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
```


## Subcommands
| command          | type  | Default| description                | Required |
| --------------------- | ------ | ------ | ------------------- | -------- |
| [address_or_key_name] | string |        | Added account name or address    | Yes     |
| [coin]                | string |        | coin type and amount | Yes     |


## Flags
| Name，shorthand         | type  | Default        | Description                      | Required |
| -------------------- | ------ | -------------- | -------------------------------- | -------- |
| -h, --help           |        |                | help for add-genesis-account  | No  |
| --home-client        | string | ~/.hashgardcli | client's home directory       | No   |
| --vesting-amount     | string |                | amount of coins for vesting accounts  | No    |
| --vesting-end-time   | int    |                | schedule end time (unix epoch) for vesting accounts| No    |
| --vesting-start-time | int    |                | schedule start time (unix epoch) for vesting accounts| No    |
| --home               | string | ~/.hashgard    | directory for config and data| No    |
| --trace              | bool   |                | print out full stack trace on errors| No   |


## Example
```bash
hashgardcli keys add root
hashgard add-genesis-account root 100000000gard
```

