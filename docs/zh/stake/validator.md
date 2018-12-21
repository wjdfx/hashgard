# hashgardcli stake validator

## 描述

查询验证者

## 用法

```
hashgardcli stake validator [validator-address] [flags]
```
打印帮助信息
```
hashgardcli stake validator --help
```

## 示例

查询验证者
```
hashgardcli stake validator [validator-address] --trust-node
```

运行成功以后，返回的结果类似如下：

```txt
Validator 
Operator Address: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Validator Consensus Pubkey: gardvalconspub1zcjduepq7h0hv847a27ck3vmn4ednw5qrsjeykhdg7gnuj28ls5snsallt3svmlckm
Jailed: false
Status: Bonded
Tokens: 89.1000000000
Delegator Shares: 89.1000000000
Description: {instance-c5m0fg87  http://hgdev.com hashgard}
Bond Height: 0
Unbonding Height: 0
Minimum Unbonding Time: 1970-01-01 00:00:00 +0000 UTC
Commission: {rate: 0.1000000000, maxRate: 0.2000000000, maxChangeRate: 0.0100000000, updateTime: 0001-01-01 00:00:00 +0000 UTC}

```
