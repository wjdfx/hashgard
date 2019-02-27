# hashgard collect-gentxs

## 描述

```
从 config/gentx/ 目录收集验证人节点的 gentx-xxxx.json 文件，并添加到 genesis.json 文件中。
```

## 使用方式

```
hashgard collect-gentxs [flags]
```

## 标志

| 命令，缩写  | 默认值                    | 描述                                                         | 是否必须 |
| ----------- | ------------------------- | ------------------------------------------------------------ | -------- |
| --gentx-dir | ~/.hashgard/config/gentx/ | 覆盖默认的“gentx”目录，从中收集并执行genesis事务；默认 [--home]/config/gentx/ | 否       |
| -h, --help  |                           | collect-gentxs 模块帮助文档                                  | 否       |
| --home      | ~/.hashgard               | 配置和数据的目录（默认为“~/.hashgard”）                      | 否       |

## 例子

`hashgard collect-gentxs`

