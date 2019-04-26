# hashgardcli issue search

## 描述
根据代币符号来搜索发行的代币信息
## 使用方式
```bash
hashgardcli issue search [symbol] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 示例

### 搜索
```bash
hashgardcli issue search fo
```
```json
[
 {
  "issue_id": "coin174876e800",
  "issuer": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc",
  "owner": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc",
  "issue_time": "2019-04-18T06:05:01.378656183Z",
  "name": "foocoin",
  "symbol": "FOO",
  "total_supply": "99998224",
  "decimals": "18",
  "description": "",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 }
]

```

