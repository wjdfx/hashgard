# hashgard unsafe-reset-all

## 描述

重置区块链数据库，删除地址簿文件，并将 priv_validator.json 重置为 genesis 状态

## 使用方式

```
hashgard unsafe-reset-all [flags]
```

## flags

| 命令，缩写 | 默认值      | 描述                      | 是否必须 |
| ---------- | ----------- | ------------------------- | -------- |
| -h, --help |             | unsafe-reset-all 模块帮助 | 否       |
| --home     | ~/.hashgard | 配置和数据的目录          | 否       |

## 例子

`hashgard unsafe-reset-all`
