# hashgard gentx

## 描述

```
此命令是 'hashgard tx create-validator' 命令的别名。
它创造了一个带有自我授权的创世纪
以下委托和佣金默认参数：
授权金额：10000000 GARD
佣金率：0.1
佣金最高税率：0.2
佣金最高变动率：0.01
最小自委托股权：100股 (1亿GARD)
```

## 使用方式

```
 hashgard gentx [flags]
```

## 标志

| 命令，缩写                   | 默认值         | 描述                                            | 是否必须 |
| ---------------------------- | -------------- | ----------------------------------------------- | -------- |
| --amount                     |                | 债券的金额                                      | 是       |
| --commission-max-change-rate |                | 最高佣金变动率百分比（每天）                    | 否       |
| --commission-max-rate        |                | 最高佣金率百分比                                | 否       |
| --commission-rate            |                | 初始佣金率百分比                                | 否       |
| --home-client                | ~/.hashgardcli | 客户端的主目录（默认为“~/.hashgardcli”）        | 否       |
| --ip                         | 本机 IP 地址   | 节点公网 IP 地址                                | 否       |
| --min-self-delegation        |                | 验证人最小委托股份                              | 否       |
| --name                       |                | 用于签署gentx的私钥的名称                       | 否       |
| --node-id                    |                | 节点 ID                                         | 否       |
| --output-document            |                | 将genesis事务JSON文档写入给定文件而不是默认位置 | 否       |
| --pubkey                     |                | 验证器的Bech32编码的PubKey                      | 否       |
| -h, --help                   |                | gentx模块帮助文档                               | 否       |
| --home                       | ~/.hashgard    | 配置和数据的目录（默认为“~/.hashgard”）         | 否       |
| --trace                      |                | 在出错时打印完整的调用栈                        | 否       |

## 例子

`hashgard gentx --name=root --amount=100000000gard --ip=${validator_ip}`

