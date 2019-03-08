# hashgard collect-gentxs

## 描述

从 config/gentx/ 目录收集验证人节点的 gentx-xxxx.json 文件，并添加到 genesis.json 文件中。

## 使用方式

```
hashgard collect-gentxs [flags]
```

## flags

| 命令，缩写  | 类型   | 默认值                    | 描述                                               | 是否必须 |
| ----------- | ------ | ------------------------- | -------------------------------------------------- | -------- |
| --gentx-dir | string | ~/.hashgard/config/gentx/ | 覆盖默认的“gentx”目录，从中收集并执行 genesis 事务 | 否       |
| -h, --help  |        |                           | collect-gentxs 模块帮助文档                        | 否       |
| --home      | string | ~/.hashgard               | 配置和数据的目录                                   | 否       |
| --trace     | bool   |                           | 在出错时打印完整的调用栈                           | 否       |

## 例子

`hashgard collect-gentxs`
