# hashgardcli issue create

## Description

Issue a new token

## Usage

```
hashgardcli issue create [name] [symbol] [total-supply] [flags] --from
```

## Flags

| Name          | type| Required  | Default| Description                               |
| ------------- | ---- | -------- | ------ | --------------------------------------- |
| --decimals    | int  | Yes    | 18     | Decimals of the token |
| --burn-owner  | Bool | No   | false  | Disable token owner burn the token |
| --burn-holder | bool | No  | false  | 关闭普通账号销毁该自己持有的代币功能    |
| --burn-from   | bool | No   | false  | 关闭Owner销毁非管理员账户持有的代币功能 |
| --minting     | bool | No   | false  | 是否不再增发功能                        |
| --freeze      | bool | No | false  | 关闭冻结用户转入转出功能                |



## Global Flags
**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### 发行一个新币

```shell
hashgardcli issue create issuename AAA 10000000000000 --from
```

输入正确的密码之后，你就完成发行了一个代币，需要注意的是要记下你的issue-id值，这是可以检索及操作你的代币的唯一要素。

```txt
{
   Height: 2967
  TxHash: 84B19F831958A6334C4806967E66E6C8640F0A2E7958A5E99A1DF3B6B6E6378C
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 43428
  Tags: 
    - action = issue
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```

查询自己的账号

```shell
hashgardcli bank account gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye
```

你将会看到你的持币列表里多了一个形如“币名（issue-id）”特殊名称的币。后续对该币的操作请使用issue-id的值来进行，包括进行转账操作，要转的币也请使用该issue-id。

```
{
 Account:
  Address:       gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Pubkey:        gardpub1addwnpepqfpd8mkl3jg43fw7y02fe99cgaxutf5npv9y9gx9dvrrcdwl36shv694apw
  Coins:         9999999990001issuename(coin174876e802)
  AccountNumber: 0
  Sequence:      16
}
```