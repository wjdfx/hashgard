# hashgardcli issue finish-minting 

## 描述
Owner对自己发行的某个代币设置为已经完成增发。
```
注：一旦完成便不可再增发。
```
## 使用方式
```
 hashgardcli issue finish-minting  [issue-id] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 完成增发
```shell
hashgardcli issue finish-minting  coin174876e800 --from foo -o=json
```
输入正确的密码之后，你的该代币便关闭了可由Owner销毁的功能。
```txt
{
 "height": "4952",
 "txhash": "4D0C00B78A7403B5151822B064D6AA4210E32A173A44EC93061CC0CB8FD6DA43",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000",
 "gas_used": "9086568",
 "tags": [
  {
   "key": "action",
   "value": "issueFinishMinting"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
