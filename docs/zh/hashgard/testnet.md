# hashgard testnet

## 描述

```
初始化Hashgard testnet的文件
```

## 使用方式

```
hashgard testnet [flags]
```

## 标志

| 命令，缩写            | 默认值      | 描述                                                         | 是否必须 |
| --------------------- | ----------- | ------------------------------------------------------------ | -------- |
| -h, --help            |             | testnet模块帮助                                              | 否       |
| --chain-id            |             | genesis file chain-id，如果留空则将被随机创建                | 是       |
| --node-cli-home       | hashgardcli | 节点的cli配置的主目录（默认为“hashgardcli”）                 | 否       |
| --node-daemon-home    | hashgard    | 节点守护程序配置的主目录（默认为“hashgard”）                 | 否       |
| --node-dir-prefix     | node        | 使用（节点结果在node0，node1，...中）为每个节点添加前缀名称（默认为“node”） | 否       |
| --output-dir          | ./mytestnet | 用于存储testnet初始化数据的目录（默认为“./mytestnet”）       | 否       |
| --starting-ip-address | 192.168.0.1 | 起始IP地址（192.168.0.1导致持久对等列表ID0@192.168.0.1:26656，ID1 @ 192.168.0.2：26656，...）（默认为“192.168.0.1”） | 否       |
| --v                   |             | 初始化testnet的验证器数量                                    | 否       |

## 例子

`hashgard testnet--chain-id=${chain-id}`

