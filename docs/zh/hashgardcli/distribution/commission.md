# hashgardcli distribution commission

## 描述

查询验证人和委托人的佣金奖励
## 用法

```bash
hashgardcli distribution commission [validator] [flags]
```

## flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 示例

查询参数信息

```bash
hashgardcli distribution commission gardvaloper1m0g2n0r7l6s44sac2knmx40hlsdyv4esgcwg8w \
    --trust-node


```

运行成功以后，返回的结果如下：

```json
[
 {
  "denom": "gard",
  "amount": "0.337966901187138531"
 }
]
```
