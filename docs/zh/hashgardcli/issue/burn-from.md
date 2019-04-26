# hashgardcli issue burn-from

## 描述
某个代币的Owner在没有关闭持币者自己可以销毁自己持有该代币前提下，持币者对自己持有的该代币进行销毁。
## 使用方式
```bash
 hashgardcli issue burn-from [issue-id] [amount] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 示例
### 销毁代币
```bash
hashgardcli issue burn-from coin174876e800 88888 --from=foo 
```
输入正确的密码之后，你的该代币的便完成了销毁。
```json
{
 "height": "4246",
 "txhash": "85A5A71E957424B5F702807A799DD4E372F5043AAD26A89373867E1596D88D15",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "1000000",
 "gas_used": "30594",
 "tags": [
  {
   "key": "action",
   "value": "issue_burn_from"
  },
  {
   "key": "sender",
   "value": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
