# hashgard validate-genesis

## 描述

验证 genesis 文件内容的有效性

## 使用方式

```
hashgard validate-genesis [file] [flags]
```

## 参数

| 命令，缩写 | 类型   | 默认值                          | 描述             | 是否必须 |
| ---------- | ------ | ------------------------------- | ---------------- | -------- |
| [file]     | string | ~/.hashgard/config/genesis.json | genesis 文件位置 | 否       |

## flags

| 命令，缩写 | 类型   | 默认值      | 描述                             | 是否必须 |
| ---------- | ------ | ----------- | -------------------------------- | -------- |
| -h, --help |        |             | add-genesis-account 模块帮助文档 | 否       |
| --home     | string | ~/.hashgard | 配置和数据的目录                 | 否       |
| --trace    | bool   |             | 在出错时打印完整的调用栈         | 否       |

## 例子

```bash
hashgard validate-genesis
```
