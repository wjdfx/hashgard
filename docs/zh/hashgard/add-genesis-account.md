# hashgard add-genesis-account

## 描述

```
将genesis帐户添加到genesis.json
```

## 使用方式

```
hashgard add-genesis-account [address] [coin][,[coin]] [flags]
```

## 标志

| 命令，缩写 | 默认值          | 描述                                        | 是否必须 |
| ---------- | --------------- | ------------------------------------------- | -------- |
| -h, --help |                 | add-genesis-account模块帮助文档             | 否       |
| --home     | /root/.hashgard | 配置和数据的目录（默认为“/root/.hashgard”） | 否       |

## 例子

`hashgard add-genesis-account ${wallet_address} ${amount}${coin}`

