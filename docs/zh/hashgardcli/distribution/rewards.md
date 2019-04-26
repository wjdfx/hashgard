# hashgardcli distribution rewards

## 描述

查询委托人获得的所有奖励，可选择限制来自单个验证者的奖励

## 用法

```bash
hashgardcli distribution rewards [delegator-addr] [<validator-addr>] [flags]
```

## flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 示例

查询参数信息

```bash
hashgardcli distribution rewards gard1hr7vm7t7paeyg33ggd6efek2w58mu2hutvjrms \
    --trust-node

```

运行成功以后，返回的结果如下：

```json
[
 {
  "denom": "gard",
  "amount": "0.131833867963517125"
 }
]
```

