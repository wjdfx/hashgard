# hashgard init

## 描述

初始化genesis config，priv-validator文件，p2p-node文件和应用程序配置文件

## 使用方式

```
 hashgard init [command]
```

## 标志

| 命令，缩写  | 默认值 | 描述                                          | 是否必须 |
| ----------- | ------ | --------------------------------------------- | -------- |
| -h, --help  |        | init模块帮助                                  | 否       |
| --chain-id  |        | genesis file chain-id，如果留空则将被随机创建 | 是       |
| --moniker   |        | 设置节点的名称                                | 是       |
| --overwrite |        | 覆盖genesis.json文件                          | 否       |

## 例子

`hashgard init --chain-id=${chain-id} --moniker=${moniker_name}`

