# hashgardcli stake unbonding-delegation

## 描述

基于委托者地址和验证者地址的unbonding-delegation记录查询

## 用法

```
hashgardcli stake unbonding-delegation [flags]
```
打印帮助信息
```
hashgardcli stake unbonding-delegation --help
```

## 特有的flags

| 名称, 速记           | 默认值                     | 描述                                                                 | 必需     |
| ------------------- | -------------------------- | ------------------------------------------------------------------- | -------- |
| --address-delegator |                            | [string] 委托者bech地址                                              | Yes      |
| validator |                            | [string] 验证者bech地址                                             | Yes      |

## 示例

查询unbonding-delegation
```
hashgardcli stake unbonding-delegation --address-delegator=gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --validator=gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --chain-id=hashgard
```

运行成功以后，返回的结果如下：

```txt
Unbonding Delegation 
Delegator: gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet
Validator: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Creation height: 12610
Min time to unbond (unix): 2018-12-20 08:07:17.286706585 +0000 UTC
Expected balance: 9gard

```
