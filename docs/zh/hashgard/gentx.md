# hashgard gentx

## 描述

```
此命令是'hashgard tx create-validator'命令'的别名。
它创造了一个带有自我授权的创世纪
以下委托和佣金默认参数：
授权金额：100
佣金率：0.1
佣金最高税率：0.2
佣金最高变动率：0.01
```

## 使用方式

```
 hashgard gentx [flags]
```

## 标志

| 命令，缩写                   | 默认值             | 描述                                            | 是否必须 |
| ---------------------------- | ------------------ | ----------------------------------------------- | -------- |
| --amount                     |                    | 债券的金额                                      | 是       |
| --commission-max-change-rate |                    | 最高佣金变动率百分比（每天）                    | 否       |
| --commission-max-rate        |                    | 最高佣金率百分比                                | 否       |
| --commission-rate            |                    | 初始佣金率百分比                                | 否       |
| --home-client                | /root/.hashgardcli | 客户端的主目录（默认为“/root/.hashgardcli”）    | 否       |
| --name                       |                    | 用于签署gentx的私钥的名称                       | 否       |
| --output-document            |                    | 将genesis事务JSON文档写入给定文件而不是默认位置 | 否       |
| --pubkey                     |                    | 验证器的Bech32编码的PubKey                      | 否       |
| -h, --help                   |                    | gentx模块帮助文档                               | 否       |
| --home                       | /root/.hashgard    | 配置和数据的目录（默认为“/root/.hashgard”）     | 否       |

## 例子

`hashgard gentx --amount=${amount}${coin} --home-clinet=${HASHGARDCLI_HOME} --home=${HASHGARD_HOME}`

