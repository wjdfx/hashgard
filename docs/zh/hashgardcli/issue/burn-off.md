# hashgardcli issue burn-off

## 描述
Owner关闭自己发行的某个代币的可由Owner销毁的功能。
```
注：一旦关闭便不可再打开。
```
## 使用方式
```
 hashgardcli issue burn-off [issue-id] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 关闭销毁功能
```shell
hashgardcli issue burn-off coin174876e800 --from foo -o=json
```
输入正确的密码之后，你的该代币便关闭了可由Owner销毁的功能。
```txt
{
 "height": "4844",
 "txhash": "2BB3FBF0D054C772CF668D948A2FE0B949E4192818C253828721C6F4EC8F7BEF",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000",
 "gas_used": "9086337",
 "tags": [
  {
   "key": "action",
   "value": "issueBurnOff"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
