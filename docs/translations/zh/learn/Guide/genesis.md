# 如何参与Genesis过程
## 第1步：创建自己的帐户
首先，如果您没有帐户或忘记密码，则需要为自己创建一个帐户作为相应的验证程序操作员。

请按照[钱包创建文档](https://github.com/hashgard/hashgard/blob/master/docs/translations/zh/cli/hashgardcli/keys/add.md)创建



## 第2步：初始化您的节点

```bash
hashgard init --chain-id=sif-4000 --moniker=${your_node_name}
```

> 注：已初始化过的不用执行此步骤

## 第3步：向 genesis.json 中添加账户信息

```bash
hashgard add-genesis-account ${your_wallet_name} 100000000gard
```
> 注意：name仅支持ASCII字符。使用Unicode字符将使您的节点无法访问。

## 第4步：向 genesis.json 中添加账户信息

```bash
hashgard gentx --name=${your_wallet_name} --amount=100000000gard --ip=${validator_ip}
```

> 注意：name仅支持ASCII字符。使用Unicode字符将使您的节点无法访问。
>

将在以下目录中生成事务：~/.hashgard/config/gentx 创建 CreateValidator 事务并通过刚刚创建的验证器操作员帐户对事务进行签名默认佣金数据为：

- delegation amount: 100000000 gard
- commission rate: 0.1
- commission max rate: 0.2
- commission max change rate: 0.01



## 第5步：提交你的 gentx.json
通过创建拉取请求提交您的 gentx-node-ID.json 。
在团队收集了所有 gen-tx 事务后，我们将发布```genesis```文件。
然后，您可以下载最终的```genesis```文件并启动一个节点。
