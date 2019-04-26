# hashgardcli issue query

## 描述
查询指定issue-id值的发行的币的信息。
## 使用方式
```bash
hashgardcli issue query [issue-id] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 示例
### 查询发行信息
```bash
hashgardcli issue query coin174876e800
```
```json
{
 "type": "issue/CoinIssueInfo",
 "value": {
  "issue_id": "coin174876e800",
  "issuer": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye",
  "owner": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye",
  "issue_time": "2019-04-17T08:17:05.109247975Z",
  "name": "mycoin",
  "symbol": "MY",
  "total_supply": "9999999999999838889",
  "decimals": "18",
  "description": "",
  "burning_off": true,
  "burning_from_off": true,
  "burning_any_off": true,
  "minting_finished": false
 }
}
```

