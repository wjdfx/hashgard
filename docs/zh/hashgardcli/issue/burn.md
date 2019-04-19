# hashgardcli issue burn

## 描述
在发行的代币可以销毁的前提下，Owner可以对自己发行的代币进行销毁。
## 使用方式
```
 hashgardcli issue burn [issue-id] [amount] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 销毁代币
```shell
hashgardcli issue burn coin174876e800 88888 --from=foo -o=json
```
输入正确的密码之后，你的该代币的便完成了销毁。
```txt
{
 "height": "4007",
 "txhash": "1972DC3A17E74FE8030CB9F551B0C14050D9397AB3ED3CD3F271A38BA7C831AB",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "10000000000000000",
 "gas_used": "18989013",
 "tags": [
  {
   "key": "action",
   "value": "issue_burn"
  },
  {
   "key": "sender",
   "value": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```