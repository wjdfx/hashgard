# hashgardcli issue burn-from-off

## 描述
Owner关闭自己发行的某个代币的可由持币者销毁的功能。
```
注：一旦关闭便不可再打开。
```
## 使用方式
```
 hashgardcli issue burn-from-off [issue-id] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 关闭销毁功能
```shell
hashgardcli issue burn-from-off coin174876e800 --from foo -o=json
```
输入正确的密码之后，你的该代币便关闭了可由Owner销毁的功能。
```txt
{
 "height": "4880",
 "txhash": "6E18360856EF8101415DE5F92F6044BE812899EBC73B87A156344FFB59ACD193",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000",
 "gas_used": "9086433",
 "tags": [
  {
   "key": "action",
   "value": "issue_burn_from_off"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
