# 重新委托
对于您曾委托过的验证器，您可以通过再次委托，对委托给前者的代币重新委托给后者

您可以通过如下操作：
```bash
 hashgardcli stake redelegate \
     --chain-id=${chain-id} \
     --from=${wallet_name}  \
     --addr-validator-source=${validator_address}  \
     --addr-validator-dest=${validator_address} \
     --shares-fraction=0.1 \
     --gas=simulate
```

```addr-validator-source``` 是您需要转移委托的原验证器地址
``` addr-validator-dest``` 是您需要转移到的新的验证器地址
``` shares-amount``` 转移的share数量，正数
```shares-fraction```转移的比率，0到1之间的正数 

设置```gas=simulate```会自动计算gas

**您可以用`--shares-amount`或者`--shares-percent`指定装委托的token数量。记住，这两个参数不可同时使用。**

