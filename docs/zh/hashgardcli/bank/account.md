# hashgardcli bank account

## 描述

查询账户信息

## 使用方式

```
hashgardcli bank account [address] [flags] 
```

## 标志

| 命令，缩写   | 类型   | 是否必须 | 默认值                | 描述                                                         |
| ------------ | ------ | -------- | --------------------- | ------------------------------------------------------------ |
| --height     | int    | 否       |                       | 查询某个高度的区块链数据，如果是0，这返回最新的区块链数据。  |
| -h, --help   | string | 否       |                       | 打印帮助信息                                                 |
| --indent     | bool   | 否       | false                 | 格式化json字符串                                             |
| --ledger     | string | 否       |                       | 是否使用硬件钱包                                             |
| --node       | string | 否       | tcp://localhost:26657 | <主机>:<端口> tendermint节点的rpc地址。                      |
| --trust-node | string | 否       | True                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性 |

## 全局标志

| 命令，缩写            | 默认值         | 描述                                | 是否必须 |
| --------------------- | -------------- | ----------------------------------- | -------- |
| --chain-id | string | tendermint 节点网络ID | 是 |
| -e, --encoding string | hex            | 字符串二进制编码 (hex \|b64 \|btc ) | 否       |
| --home string         | $HOME/.hashgardcli | 配置和数据存储目录                  | 否       |
| -o, --output string   | text           | 输出格式 (text \|json)              | 否       |
| --trace               |                | 出错时打印完整栈信息                | 否       |

## 例子

### 查询账户信息 

```
 hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx -o json --trust-node --indent
```

执行完命令后，获得账户的详细信息如下：

```
{
 "type": "auth/Account",
 "value": {
  "address": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn",
  "coins": [
   {
    "denom": "gard",
    "amount": "1900000000"
   }
  ],
  "public_key": {
   "type": "tendermint/PubKeySecp256k1",
   "value": "AvM1uBBEml3ZtXP5GZD6vr7UIcit6GMjS0ZUdxuejShH"
  },
  "account_number": "0",
  "sequence": "1"
 }
}
```
如果你查询一个错误的地址，将会返回如下信息:
```
hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429zz
ERROR: decoding bech32 failed: checksum failed. Expected d429yx, got d429zz.
```
如果查询一个空地址，将会返回如下信息:
```
hashgardcli bank account gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7
ERROR: No account with address gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 was found in the state.
Are you sure there has been a transaction involving it?
```