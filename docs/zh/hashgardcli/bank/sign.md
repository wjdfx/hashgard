# hashgardcli bank sign

## 描述

签名生成的离线传输文件。该文件由 --generate-only 标志生成。

## 使用方式

```
hashgardcli bank sign <file> [flags]
```

 

## 标志

| 命令，速记       | 类型    | 是否必须 | 默认值                | 描述                                                         |
| ---------------- | ------- | -------- | --------------------- | ------------------------------------------------------------ |
| --account-number | int     | 否       |                       | 发起交易的账户的编号                                         |
| --append | bool | 否 | true | 将签名附加到现有签名。如果禁用，旧签名将被覆盖。如果--multisig打开则忽略（默认为true） |
| --async          | bool | 否       | false                 | 是否异步广播交易                                             |
| --dry-run        | bool | 否       | false                 | 模拟执行交易，并返回消耗的`gas`。`--gas`指定的值会被忽略     |
| --fees           | string  | 是       |                       | 交易费，例如： 10stake,1atom                                 |
| --from           | string  | 是       |                       | 发送交易的账户名称                                           |
| --gas            | string  | 否       | 2000000               | 交易的gas上限; 设置为"auto"将自动计算相应的阈值              |
| --gas-adjustment | float   | 否       | 1                     | gas调整因子，这个值降乘以模拟执行消耗的`gas`，计算的结果返回给用户; 如果`--gas`的值不是`atuo`，这个标志将被忽略 |
| --gas-prices     | string  | 否       |                       | 交易费用单价，(例如： 0.00001stake)                          |
| --generate-only  | bool | 否       | false                 | 是否仅仅构建一个未签名的交易便返回。                         |
| -h, --help       |         | 否       |                       | 打印帮助                                                     |
| --indent         | bool | 否       | false                 | 格式化json字符串                                             |
| --ledger         | bool | 否       | false                 | 是否使用硬件钱包                                             |
| --memo           | string  | 否       |                       | 备注信息                                                     |
| --multisig | string | 否 | | 代表交易签署的multisig帐户的地址 |
| --name | string | 否 | | 与之签名的私钥的名称 |
| --node           | string  | 否       | tcp://localhost:26657 | <主机>:<端口> tendermint节点的rpc地址。                      |
| --offline | bool | 否 | false | 链下模式，不查询全节点 |
| --output-document | string |  |  | 该文档将写入给定文件而不是STDOUT |
| --print-response | bool | 否       | true | 是否打印交易返回结果，仅在`async`为false的情况下有效  |
| --sequence       | int     | 否       |                       | 发起交易的账户的sequence                                     |
| --signature-only | bool | 否 | | 仅打印生成的签名，然后退出 |
| --trust-node     | bool | 否       | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性 |
| --validate-signatures | bool | 否 | false | 打印必须签署交易的地址，已签名的地址，并确保签名的顺序正确 |


## 全局标志

| 命令，速记            | 默认值         | 描述                                | 是否必须 | 类型   |
| --------------------- | -------------- | ----------------------------------- | -------- | ------ |
| | --chain-id | string | tendermint 节点网络ID | 是 |
| -e, --encoding string | hex            | 字符串二进制编码 (hex \|b64 \|btc ) | False    | String |
| --home string         | /root/.hashgardcli | 配置和数据存储目录                  | False    | String |
| -o, --output string   | text           | 输出格式(text \|json)               | False    | String |
| --trace               |                | 出错时打印完整栈信息                | False    |        |

## 例子

### 对一个离线发送文件签名

首先你必须使用 **hashgardcli bank send**  命令和标志 **--generate-only** 来生成一个发送记录，如下

```  
hashgardcli bank send --to=gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx  --from=test --chain-id=hashgard --amount=10gard --generate-only

{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/Send","value":{"inputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}],"outputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}]}}],"fee":{"amount":[{"denom":"gard","amount":"4000000000000000"}],"gas":"200000"},"signatures":null,"memo":""}}
```



保存输出到文件中，如  /root/output/output/node0/test_send_10hashgard.txt.

接着来签名这个离线文件.

```
hashgardcli bank sign /root/output/output/node0/test_send_10hashgard.txt --name=test  --offline=false --print-sigs=false --append=true
```

随后得到签名详细信息，如下输出中你会看到签名信息。 
**ci+5QuYUVcsARBQWyPGDgmTKYu/SRj6TpCGvrC7AE3REMVdqFGFK3hzlgIphzOocGmOIa/wicXGlMK2G89tPJg==**

```
hashgardcli bank sign /root/output/output/node0/test_send_10hashgard.txt --name=test  --offline=false --print-sigs=false --append=true
Password to sign with 'test':
{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/Send","value":{"inputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}],"outputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}]}}],"fee":{"amount":[{"denom":"gard","amount":"4000000000000000"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AzlCwiA5Tvxwi7lMB/Hihfp2qnaks5Wrrgkg/Jy7sEkF"},"signature":"ci+5QuYUVcsARBQWyPGDgmTKYu/SRj6TpCGvrC7AE3REMVdqFGFK3hzlgIphzOocGmOIa/wicXGlMK2G89tPJg==","account_number":"0","sequence":"2"}],"memo":""}}
```

