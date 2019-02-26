# hashgard unsafe-reset-all

## 描述

```
重置区块链数据库，删除地址簿文件，并将priv_validator.json重置为genesis状态
```

## 使用方式

```
hashgard unsafe-reset-all [flags]
```

## 标志

| 命令，缩写 | 默认值          | 描述                                        | 是否必须 |
| ---------- | --------------- | ------------------------------------------- | -------- |
| -h, --help |                 | unsafe-reset-all模块帮助                    | 否       |
| --home     | /root/.hashgard | 配置和数据的目录（默认为“/root/.hashgard”） | 否       |

## 例子

`hashgard unsafe-reset-all --home=/root/.hashgard`

