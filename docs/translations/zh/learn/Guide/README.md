# Hashgard 测试网指南

该文档目的是介绍如何运行一个 Hashgard 节点并加入当前的测试网。

> 当前为 sif-5001 测试网络，请务必指定 chain-id = sif-5001

> \${}代表变量，需要你手动替换成你自己设定的值

## 运行 Hashgard 节点

### 步骤 1：在您的服务器上安装 Hashgard

请参照 [hashgard 安装文档](https://github.com/hashgard/hashgard/blob/develop/docs/translations/zh/learn/installation.md) 安装 Hashgard。

### 步骤 2：创建钱包

如果尚未创建钱包，请参照 [钱包创建文档](https://github.com/hashgard/hashgard/blob/develop/docs/translations/zh/cli/hashgardcli/keys/add.md) 创建钱包。

### 步骤 3：设置默认参数

设置 hashgardcli 命令行客户端的默认连接节点：

```
hashgardcli config chain-id sif-5001
hashgardcli config trust-node true
```

格式化输出：

```
hashgardcli config indent true
hashgardcli config output json
```

### 步骤 4：初始化节点并获取配置文件

#### 4.1：初始化节点

```bash
hashgard init --chain-id=sif-5001 --moniker=${your_node_name}
```

如果您想在 genisis 阶段成为验证人，您可以按照文档 [参与 genesis](genesis.md) 来生成一个 json 文件，并提交给我们。
或者，您可以随后再发送相关交易，同样可以升级为验证人节点。

#### 4.2：下载 genesis 和 config 文件

进入 hashgard 的 config 文件路径，使用 sif-5001 测试网络的 genesis 和 config 文件替换掉 hashgard 初始化生成的对应文件。

```bash
cd ~/.hashgard/config/
rm genesis.json
rm config.toml
curl https://raw.githubusercontent.com/hashgard/testnets/master/sif/sif-5001/config/config.toml > config.toml
curl https://raw.githubusercontent.com/hashgard/testnets/master/sif/sif-5001/config/genesis.json > genesis.json
```

#### 4.3: 修改配置文件

编辑 `~/.hashgard/config/config.toml` 文件，
您需要修改其中 `moniker` 和 `external_address` 字段，
`moniker` 是您的节点名称， `external_address` 是您的服务器外网 ip 地址：

```toml
# A custom human readable name for this node
moniker = "${your_node_name}"
external_address = "${your_public_ip}:26656"
```

### 步骤 5：运行完整节点

使用以下命令启动节点：

```bash
hashgard start > hashgard.log &
```

检查节点状态是否正常：

```bash
hashgardcli status
```

如果节点成功运行，您将看到以下输出:

```json
{
  "node_info": {
    "protocol_version": {
      "p2p": "5",
      "block": "8",
      "app": "0"
    },
    "id": "b783ac2b41da096ddc3a26c2a39e3b0c1ea49d9e",
    "listen_addr": "127.0.0.190:26656",
    "network": "hashgard",
    "version": "0.27.0",
    "channels": "4020212223303800",
    "moniker": "hashgard_test",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "sync_info": {
    "latest_block_hash": "AFB6261A76DCC6176FF5248B3B5BEDEEBD38B6DA3D18AD21ADD4054AEDEED016",
    "latest_app_hash": "1DEAF3D71AD735F4E375439DAFD96C8934E944D8D32F6179F55C5470E219D132",
    "latest_block_height": "77280",
    "latest_block_time": "2018-12-24T07:13:53.787154732Z",
    "catching_up": false
  },
  "validator_info": {
    "address": "5959DF3D28F601407A98D0B038174E288ED45BD7",
    "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "wYrxKp7cw14eQiqzfGBggEV474ZA2lc35AieJM5SM6Y="
    },
    "voting_power": "950"
  }
}
```

当您看到 `catching_up` 是 `false`，表示节点的区块数据与 testnet 已经同步完成，否则表示它仍在同步。

您现在已经成功运行了一个 Hashgard 完整节点并接入了测试网。

您已经完成了测试网第一个激励任务，接下来您可以根据 [激励任务说明](<https://github.com/hashgard/testnets/tree/master/sif_CN/>) 参与后续步骤完成其它激励任务。

## 后续步骤

您现在拥有一个活动的完整节点。下一步是什么？

### 步骤 5：升级到验证人( Validator )节点

如果您参与了 genesis 文件生成过程，那么一旦完全同步，您的节点就会成为验证人中的一员。

如果您错过了 genesis 文件生成过程，您仍然可以将您的节点升级成为 Hashgard 验证人节点，继续进入 [创建验证人节点](./GreatValidator.md)。

您也可以[委托](Delegate.md)，[解绑](unbond.md)，[再委托](Redelegate.md)

### 步骤 6：链式治理

去中心化的链式，您可以在链上发起提案，也可参与到投票过程中

如何发起提案？进入[发起治理提案](SubmitProposal.md)

提案时需要进入激活状态才能发起投票，在此之前，您可以进行[抵押存款](Deposit.md)

对于被正式被激活的提案，您可以对其进行[投票](Vote.md)
