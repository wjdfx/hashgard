# hashgard collect-gentxs

## 描述

```
收集genesis txs并输出genesis.json文件
```

## 使用方式

```
hashgard collect-gentxs [flags]
```

## 标志

| 命令，缩写  | 默认值                | 描述                                                         | 是否必须 |
| ----------- | --------------------- | ------------------------------------------------------------ | -------- |
| --gentx-dir | ${home}/config/gentx/ | 覆盖默认的“gentx”目录，从中收集并执行genesis事务;默认[--home] / config / gentx / | 否       |
| -h, --help  |                       | collect-gentxs模块帮助文档                                   | 否       |
| --home      | /root/.hashgard       | 配置和数据的目录（默认为“/root/.hashgard”）                  | 否       |

## 例子

`hashgard collect-gentxs`

