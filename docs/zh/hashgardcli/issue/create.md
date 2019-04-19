# hashgardcli issue create

## 描述

发行一个新的代币

## 使用方式

```
hashgardcli issue create [name] [symbol] [total-supply] [flags]
```

## Flags

| 名称               | 类型   | 是否必须 | 默认值 | 描述                                              |
| ------------------ | ------ | -------- | ------ | ------------------------------------------------- |
| --decimals         | int    | 否       | 18     | （可选）代币精度，默认18位，最大18位              |
| --burn-off         | string | 否       | false  | （可选）是否关闭Owner销毁该代币的功能             |
| --burn-from-off    | bool   | 否       | ""     | （可选）是否关闭普通账号销毁该代币的功能          |
| --burn-any-off     | bool   | 否       | false  | （可选）是否关闭Owner可销毁任意账号下该代币的功能 |
| --minting-finished | bool   | 否       | false  | （可选）是否不再增发                              |

## Global Flags

### 参考：[hashgardcli](../README.md)

## 例子

### 发行一个新币

```shell
hashgardcli issue create foocoin FOO 100000000 --from foo -o=json
```

输入正确的密码之后，你就完成发行了一个代币，需要注意的是要记下你的issue-id值，这是可以检索及操作你的代币的唯一要素。

```txt
{
 "height": "3394",
 "txhash": "81D4B2054F741E901BE5A540DDA37BF53D1DEA16C94BF9E4BBDB1D1CD548DFA1",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000000",
 "gas_used": "18994244",
 "tags": [
  {
   "key": "action",
   "value": "issue"
  },
  {
   "key": "recipient",
   "value": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```

查询自己的账号

```shell
hashgardcli bank account gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye
```

你将会看到你的持币列表里多了一个形如“币名（issue-id）”特殊名称的币。后续对该币的操作请使用issue-id的值来进行，包括进行转账操作，要转的币也请使用该issue-id。

```
{
 "type": "auth/Account",
 "value": {
  "address": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye",
  "coins": [
   {
    "denom": "foocoin(coin174876e800)",
    "amount": "100000000"
   },
   {
    "denom": "gard",
    "amount": "1010000000"
   }
  ],
  "public_key": {
   "type": "tendermint/PubKeySecp256k1",
   "value": "A/rSPb+egaljwS1XGSSFKpaFkfjFzLWJFmtUdAlaQpLl"
  },
  "account_number": "1",
  "sequence": "11"
 }
}
```