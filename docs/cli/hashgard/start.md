# hashgard start

## Description

Run the full node

## Usage

```shell
hashgard start [flags]
```

## Flags

| Name，shorthand                    | Default               | description              | Required  |
| ------------------------------- | --------------------- | ----------------------------------- | -------- |
| --abci                          | socket                | Specify abci transport (socket or grpc) | No      |
| --address                       | tcp://0.0.0.0:26658   | Listen address                                           | No  |
| --consensus.create_empty_blocks | true                  |  Set this to false to only produce blocks when there are txs or when the AppHash changes  | No  |
| --fast_sync                     | true                  | Fast blockchain syncing                            | No  |
| --minimum_fees                  |                       |  Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum                            | No  |
| --moniker                       | instance-c5m0fg87     | Node Name                                             | No  |
| --p2p.laddr                     | tcp://0.0.0.0:26656   |  Node listen address. (0.0.0.0:0 means any interface, any port)   | No  |
| --p2p.persistent_peers          |                       | Comma-delimited ID@host:port persistent peers              | No  |
| --p2p.pex                       | true                  | Enable/disable Peer-Exchange                            | No  |
| --p2p.private_peer_ids          |                       | Comma-delimited private peer IDs                      | No  |
| --p2p.seed_mode                 |                       | Enable/disable seed mode                              | No  |
| --p2p.seeds                     |                       | Comma-delimited ID@host:port seed nodes                 | No  |
| --p2p.upnp                      |                       | Enable/disable UPNP port forwarding                     | No  |
| --priv_validator_laddr          |                       | Socket address to listen on for connections from external priv_validator process| No  |
| --proxy_app                     | tcp://127.0.0.1:26658 | Proxy app address, or one of: 'kvstore', 'persistent_kvstore', 'counter', 'counter_serial' or 'noop' for local testing. | No  |
| --pruning                       | syncable              | Pruning strategy: syncable, nothing, everything     | No  |
| --replay                        |                       | Replay the last block                                  | No  |
| --rpc.grpc_laddr                |                       | GRPC listen address (BroadcastTx only). Port required | No  |
| --rpc.laddr                     | tcp://0.0.0.0:26657   |  RPC listen address. Port required            | No  |
| --rpc.unsafe                    |                       | Enabled unsafe rpc methods                              | No  |
| --trace-store                   |                       | Enable KVStore tracing to an output file                | No  |
| --with-tendermint               | true                  | Run abci app embedded in-process with tendermint | No  |
| -h, --help                      |                       | help for start                           | No  |

## Example

```shell
hashgard start --home=/root/.hashgard
```

