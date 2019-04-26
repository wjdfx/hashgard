# hashgardcli gov query-vote

## 描述

查询指定提案、指定投票者的投票情况

## 使用方式

```bash
hashgardcli gov query-vote [proposal-id] [voter-addr] [flags]
```

## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 示例

### 查询投票

```bash
hashgardcli gov query-vote 1 gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet \
    --trust-node
```

通过指定提案、指定投票者查询投票情况。

```json
{
  "voter": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
  "proposal_id": "1",
  "option": "Yes"
}

```

