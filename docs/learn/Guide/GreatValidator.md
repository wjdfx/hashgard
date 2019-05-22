# 创建验证人节点

> 注：验证人者负责通过投票将新块提交给区块链。验证人的股权如果变得不可用，双重签署交易或不投票，则会被削减。

### 创建验证人

在此之前，请确保您已经创建了一个本地钱包，[创建钱包参考文档](https://github.com/hashgard/hashgard/blob/master/docs/zh/hashgardcli/keys/add.md)。

您也可以前往水龙头地址创建并领取，再在本机执行恢复命令。
- [水龙头帮助文档](https://github.com/hashgard/hashgard/tree/master/docs/cli/hashgardcli/faucet/send.md)
- [通过助记词恢复帮助文档](https://github.com/hashgard/hashgard/blob/master/docs/cli/hashgardcli/keys/add.md)

您可以通过运行以下命令找到验证器 pubkey：
```bash
hashgard tendermint show-validator
```

接下来，通过 hashgardcli stake create-validator 命令使节点升为验证人节点：
```bash
hashgardcli stake create-validator \
    --from=${your_wallet_name}  \
	--pubkey=gardvalconspub1zcjduepqk5cft3vyae5m45gf844x4nqv4t04r07wjkclg5rdgccx32t3r5js5chg6m \
	--commission-max-change-rate=0.01 \
	--commission-rate=0.1 \
	--commission-max-rate=0.2 \
	--amount=40gard \
	--moniker=${your_node_name}  \
	--min-self-delegation=10 \
	--chain-id=sif-4000 \
	--fees=2gard \
	--output=json \
	--indent
```

### 查看验证人描述

使用以下命令查看验证人的信息：
```bash
hashgardcli stake validator ${validator-address} --trust-node
```
