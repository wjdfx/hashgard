# hashgard add-genesis-account

## 描述

```
将 genesis 帐户添加到 /path/to/.hashgard/config/genesis.json 文件中
```

## 使用方式

```
hashgard add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
```

## 标志

| 命令，缩写           | 默认值         | 描述                                   | 是否必须 |
| -------------------- | -------------- | -------------------------------------- | -------- |
| -h, --help           |                | add-genesis-account 模块帮助文档       | 否       |
| --home-client        | ~/.hashgardcli | 客户端根目录                           | 否       |
| --vesting-amount     |                | 冻结数量                               | 否       |
| --vesting-end-time   |                | 冻结结束时间 (unix epoch)              | 否       |
| --vesting-start-time |                | 冻结起始时间 (unix epoch)              | 否       |
| --home               | ~/.hashgard    | 配置和数据的目录（默认为“~.hashgard”） | 否       |
| --trace              |                | 在出错时打印完整的调用栈               | 否       |

## 例子

```bash
hashgardcli keys add root
hashgard add-genesis-account root 100gard
```





