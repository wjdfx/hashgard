# hashgardcli keys show

## Description

Return public details of one local key.

#

## Usage

```
hashgardcli keys show [name] [flags]
```

## Flags

| Name, shorthand | type      | Required    |Default       | Description                                                  |
| -------------------- | ----------------- | -------------------------------------------------------------- | -------- | -------- |
| -a, --address | string | No| "" | output the address only                    |
| --bech               | string         | No           | acc               | The Bech32 prefix encoding for a key (acc/val/cons)|
| --multisig-threshold | int              | NO        | 1                 | [uint] K out of N required signatures                          |
| --pubkey             | bool | No | false  | 仅输出公钥                                                      |

**Global flags, query command flags** [hashgardcli](../README.md)

## # Examples

### Show a given key

```shell
hashgardcli keys show MyKey
```

You'll get the local public keys with 'address' and 'pubkey' element of a given key.

```txt
NAME:   TYPE:   ADDRESS:                                    PUBKEY:
MyKey   local   gardkkm4w5pvmcw0e3vjcxqtfxwqpm3k0zakl7lxn5  gardaddwnpepq0gsl90v9dgac3r9hzgz53ul5ml5ynq89ax9x8qs5jgv5z5vyssskww57lw
```