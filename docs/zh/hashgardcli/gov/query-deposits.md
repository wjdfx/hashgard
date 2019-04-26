# hashgardcli gov query-deposits

## 描述

查询指定提案的保证金详细情况

## 使用方式

```bash
hashgardcli gov query-deposits [proposal-id] [flags]
```
## Global Flags

### 参考：[hashgardcli](../README.md)

## 示例

### 查询所有保证金

```bash
hashgardcli gov query-deposits 1 --trust-node

```

你可以查询到指定提案的所有保证金代币，包括每个存款人的充值详情。

```json
[
  {
    "depositor": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "proposal_id": "1",
    "amount": [
      {
        "denom": "gard",
        "amount": "50"
      }
    ]
  }
]

```
