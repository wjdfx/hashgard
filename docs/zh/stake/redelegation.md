# hashgardcli stake redelegation

## 描述

基于委托者地址，原源验证者地址和目标验证者地址的重新委托记录查询 

## 用法

```
hashgardcli stake redelegation [flags]
```
打印帮助信息
```
hashgardcli stake redelegation --help
```

## 特有的flags

| 名称, 速记                  | 默认值                      | 描述                                                                | 必需     |
| -------------------------- | -------------------------- | ------------------------------------------------------------------- | -------- | 
| --address-delegator        |                            | [string] 委托者bech地址                                              | Yes      |
| --addr-validator-dest   |                            | [string] 目标验证者bech地址                                          | Yes      |
| --addr-validator-source |                            | [string] 源验证者bech地址                                            | Yes      |

## 示例

查询重新委托记录
```
hashgardcli stake redelegation --addr-validator-source=SourceValidatorAddress --addr-validator-dest=DestinationValidatorAddress --address-delegator=DelegatorAddress --trust-node
```

运行成功以后，返回的结果如下：

```txt
Redelegation
Delegator: gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet
Source Validator: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Destination Validator: gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq
Creation height: 1130
Min time to unbond (unix): 2018-11-16 07:22:48.740311064 +0000 UTC
Source shares: 0.1000000000
Destination shares: 0.1000000000
```
