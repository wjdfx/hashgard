# hashgardcli stake unbonding-delegation

## 描述

基于委托人地址和验证人地址的unbonding-delegation记录查询

## 用法

```
hashgardcli stake unbonding-delegation [delegator-addr] [validator-addr] [flags]
```

打印帮助信息
```
hashgardcli stake unbonding-delegation --help
```

## 示例

查询委托人和验证人的 unbonding-delegation
```
hashgardcli stake unbonding-delegation gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --chain-id=hashgard
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
