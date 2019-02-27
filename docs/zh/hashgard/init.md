# hashgard init

## 描述

生成验证人和全节点所需要的配置文件和数据文件。

## 使用方式

```bash
hashgard init [flags]
```

## 标志

| 命令，缩写  | 默认值      | 描述                                                         | 是否必须 |
| ----------- | ----------- | ------------------------------------------------------------ | -------- |
| -h, --help  |             | init 模块帮助                                                | 否       |
| --chain-id  |             | 公链 ID，如果留空则将被随机创建                              | 否       |
| --moniker   |             | 设置节点的名称，将在浏览器的[验证人节点](https://www.gardplorer.io/validator)页面中显示 | 是       |
| --overwrite |             | 覆盖genesis.json文件                                         | 否       |
| --home      | ~/.hashgard | 配置和数据存放目录                                           | 否       |
| --trace     |             | 在出错时打印完整的调用栈                                     | 否       |

## 例子

`hashgard init --chain-id=testnet-1000 --moniker=hashgard`

