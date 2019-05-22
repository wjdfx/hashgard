# hashgard testnet

## Description

Note, strict routability for addresses is turned off in the config file.

## Usage

```
hashgard testnet [flags]
```

## Flags

| Name，shorthand          | type  | Default      | Description                                            | Required  |
| --------------------- | ------ | ------------ | ---------------------------------------------------- | -------- |
| -h, --help            |        |              | help for testnet                                    | No  |
| --chain-id            | string |              | genesis file chain-id, if left blank will be randomly created| `Yes`     |
| --minimum-gas-prices  | string | 0.000006gard |  Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum                      | `Yes`     |
| --node-cli-home       | string | hashgardcli  | Home directory of the node's cli configuration            | No  |
| --node-daemon-home    | string | hashgard     | Home directory of the node's daemon configuration| No  |
| --node-dir-prefix     | string | node         | Prefix the directory name for each node with (node results in node0, node1, ...) | No  |
| --output-dir          | string | ./mytestnet  | Directory to store initialization data for the testnet| No  |
| --starting-ip-address | string | 192.168.0.1  | Starting IP address                                     | No  |
| --v                   | int    | 4            |  Number of validators to initialize the testnet with| No  |

## Example

```shell
hashgard testnet--chain-id=${chain-id}
```
