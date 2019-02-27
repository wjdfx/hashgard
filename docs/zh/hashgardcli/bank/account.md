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
 hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx --trust-node
```

执行完命令后，获得账户的详细信息如下

```
{"type":"auth/Account","value":{"address":"gard1uhauythtet90ewtuy40v4hrymlqf5n45wcxcxc","coins":[{"denom":"gard","amount":"99890"},{"denom":"honr","amount":"10000"}],"public_key":{"type":"tendermint/PubKeySecp256k1","value":"A40U8BK0MgKiVx0kSFvwUe7y+OV32X0+4abdYP+58dp4"},"account_number":"1","sequence":"2"}}



```
如果你查询一个错误的地址，将会返回如下信息
```
 hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429zz
ERROR: decoding bech32 failed: checksum failed. Expected d429yx, got d429zz.
```
如果查询一个空地址，，将会返回如下信息。
```
hashgardcli bank account gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7
ERROR: No account with address gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 was found in the state.
Are you sure there has been a transaction involving it?
```


## 扩展描述

查询hashgard网络中的账户信息。

​    



​           
