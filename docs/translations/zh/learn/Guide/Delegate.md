# 委托代币增加验证者的股份
你可以通过对一个验证器去委托代币操作来增加其股份

## 委托代币
您可以通过运行以下命令来查看所有的验证器信息：
```bash
 hashgardcli stake validators --trust-node
```

接下来您可以选择一个验证器地址并进行委托：
```bash
hashgardcli stake delegate --validator=${validator_address} --amount=${amount}  --chain-id=${chain-id} --from=${wallet_name} 
```

进行委托之后，您可以通过查看该验证器的信息来确定您是否委托成功
```bash
hashgardcli stake validator ${validator_address}  --trust-node
```