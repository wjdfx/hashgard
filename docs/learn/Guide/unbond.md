# 解绑委托，代币收回
对于您委托过的验证器，您也可以通过解绑来收回您的代币

## 解绑操作
您可以选择一个您曾委托过的验证器地址并进行解绑：
```bash
hashgardcli stake unbond --validator=${validator_address} --shares-fraction=0.1 --from=${wallet_name} --chain-id=${chain-id}
```

进行委托之后，您可以通过查看该验证器的信息来确定您是否解绑成功
```bash
hashgardcli stake validator ${validator_address}  --trust-node
```

​	