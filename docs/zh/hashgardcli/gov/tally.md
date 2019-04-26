# hashgardcli gov tally

## 描述

查询指定提案的投票统计

## 使用方式

```bash
 hashgardcli gov tally [proposal-id] [flags]
```

## Global Flags

 ### 参考：[hashgardcli](../README.md)


## 示例

### 查询投票统计

```bash
hashgardcli gov tally 1 --trust-node
```

可以查询指定提案每个投票选项的投票统计。

```json
{
  "yes": "89.1000000000",
  "abstain": "0.0000000000",
  "no": "0.0000000000",
  "no_with_veto": "0.0000000000"
}
```

