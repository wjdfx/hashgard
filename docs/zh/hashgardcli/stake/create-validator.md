# hashgardcli stake create-validator

## 介绍

发送交易申请成为验证人，并在在此验证人上委托一定数额的 token

## 用法

```
hashgardcli stake create-validator [flags]
```

## flags

| 名称                         | 类型   | 是否必填 | 默认值 | 功能描述                                                |
| ---------------------------- | ------ | -------- | ------ | ------------------------------------------------------- |
| --address-delegator          | string | false    | ""     | 委托人地址                                              |
| --amount                     | string | true     | ""     | 委托 token 的数量                                       |
| --commission-max-change-rate | float  | false    | 0.0    | 佣金比率每天变化的最大数量                              |
| --commission-max-rate        | float  | false    | 0.0    | 最大佣金比例                                            |
| --commission-rate            | float  | false    | 0.0    | 初始佣金比例                                            |
| --details                    | string | false    | ""     | 验证人节点的详细信息                                    |
| --genesis-format             | bool   | false    | false  | 是否已 genesis transaction 的方式倒出                   |
| --identity                   | string | false    | ""     | 身份信息的签名                                          |
| --ip                         | string | false    | ""     | 验证人节点的 IP, 与`--generate-only`flag 同时使用时生效 |
| --min-self-delegation        | string | true     | ""     | 验证人节点要求的自委托最小股权数量                      |
| --moniker                    | string | true     | ""     | 验证人节点名称                                          |
| --node-id                    | string | false    | ""     | 节点 ID                                                 |
| --pubkey                     | string | true     | ""     | Amino 编码的验证人公钥                                  |
| --website                    | string | false    | ""     | 验证人节点的网址                                        |

**全局 flags、发送交易命令 flags** 参考：[hashgardcli](../README.md)

## 示例

```shell
hashgardcli stake create-validator \
--chain-id=hashgard \
--from=gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu \
--pubkey=gardpub1addwnpepqdds3r4g45dgp2va978tkj5gzucnt7z7mhgqwsv6u3cnw6utwe3y2m7vdxe \
--commission-max-change-rate=0.01 \
--commission-max-rate=0.2 \
--commission-rate=0.1 \
--amount=100gard \
--moniker=testing
```
