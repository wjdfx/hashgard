# hashgard start

## 描述

开始hashgard服务

## 使用方式

```
 hashgard start [command]
```

## 标志

| 命令，缩写                      | 默认值                | 描述 | 是否必须 |
| ------------------------------- | --------------------- | ---- | -------- |
| --abci                          | socket                |      | 否       |
| --address                       | tcp://0.0.0.0:26658   |      | f否      |
| --consensus.create_empty_blocks | true                  |      | 是       |
| --fast_sync                     | true                  |      | 否       |
| --minimum_fees                  |                       |      |          |
| --moniker                       | instance-c5m0fg87     |      |          |
| --p2p.laddr                     | tcp://0.0.0.0:26656   |      |          |
| --p2p.persistent_peers          |                       |      |          |
| --p2p.pex                       | true                  |      |          |
| --p2p.private_peer_ids          |                       |      |          |
| --p2p.seed_mode                 |                       |      |          |
| --p2p.seeds                     |                       |      |          |
| --p2p.upnp                      |                       |      |          |
| --priv_validator_laddr          |                       |      |          |
| --proxy_app                     | tcp://127.0.0.1:26658 |      |          |
| --pruning                       | syncable              |      |          |
| --replay                        |                       |      |          |
| --rpc.grpc_laddr                |                       |      |          |
| --rpc.laddr                     | tcp://0.0.0.0:26657   |      |          |
| --rpc.unsafe                    |                       |      |          |
| --trace-store                   |                       |      |          |
| --with-tendermint               | true                  |      |          |
| -h, --help                      |                       |      | 否       |

## 例子

`hashgard start` 

