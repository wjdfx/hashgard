# hashgard start

## 描述

开始hashgard服务

## 使用方式

```
 hashgard start [command]
```

## 标志

| 命令，缩写                      | 默认值                | 描述                                                         | 是否必须 |
| ------------------------------- | --------------------- | ------------------------------------------------------------ | -------- |
| --abci                          | socket                | 指定abci的传输方式，socket或grpc，默认为socket               | 否       |
| --address                       | tcp://0.0.0.0:26658   | 监听地址（默认为“tcp：//0.0.0.0：26658”）                    | 否       |
| --consensus.create_empty_blocks | true                  | 设置为false以仅在有txs或AppHash更改时生成块（默认为true）    | 是       |
| --fast_sync                     | true                  | 快速区块链同步（默认为true）                                 | 否       |
| --minimum_fees                  |                       | 验证人接受交易的最低费用                                     |          |
| --moniker                       | instance-c5m0fg87     | 节点名称（默认为“instance-c5m0fg87”）                        |          |
| --p2p.laddr                     | tcp://0.0.0.0:26656   | 节点监听地址。 （0.0.0.0:0表示任何接口，任何端口）（默认为“tcp：//0.0.0.0：26656”） |          |
| --p2p.persistent_peers          |                       | 以逗号分隔的ID @ host：端口持久对等体                        |          |
| --p2p.pex                       | true                  | 启用/禁用Peer-Exchange（默认为true）                         |          |
| --p2p.private_peer_ids          |                       | 以逗号分隔的私有对等ID                                       |          |
| --p2p.seed_mode                 |                       | 启用/禁用种子模式                                            |          |
| --p2p.seeds                     |                       | 逗号分隔的ID @ host：端口种子节点                            |          |
| --p2p.upnp                      |                       | 启用/禁用UPNP端口转发                                        |          |
| --priv_validator_laddr          |                       | 侦听来自外部priv_validator进程的连接的套接字地址             |          |
| --proxy_app                     | tcp://127.0.0.1:26658 | 代理应用程序地址，或“nilapp”或“kvstore”用于本地测试。 （默认“tcp：//127.0.0.1：26658”） |          |
| --pruning                       | syncable              | 更改策略：可同步，没有，一切（默认“可同步”）                 |          |
| --replay                        |                       | 重播最后一个块                                               |          |
| --rpc.grpc_laddr                |                       | GRPC监听地址（仅限BroadcastTx）。需要端口                    |          |
| --rpc.laddr                     | tcp://0.0.0.0:26657   | RPC监听地址。需要端口（默认为“tcp：//0.0.0.0：26657”）       |          |
| --rpc.unsafe                    |                       | 启用不安全的rpc方法                                          |          |
| --trace-store                   |                       | 启用KVStore跟踪到输出文件                                    |          |
| --with-tendermint               | true                  | 用tendermint运行嵌入进程的abci app（默认为true）             |          |
| -h, --help                      |                       | start模块的帮助文档                                          | 否       |

## 例子

`hashgard start` 

