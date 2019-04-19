# hashgardcli issue burn-any

## 描述
某个代币的Owner在没有关闭销毁任意持币者持有的该代币前提下，Owner可以销毁任意持币者的该代币。
## 使用方式
```
 hashgardcli issue burn-any [issue-id] [address] [amount] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 销毁代币
```shell
hashgardcli issue burn-any coin174876e800 gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye 888 --from=foo -o=json
```
输入正确的密码之后，你便对指定的持币者的该代币完成了销毁。
```txt
{
 "height": "4320",
 "txhash": "2EEEBF75D230ED8F31B6A57FAAFC935E42FF055BBEABF141E125AC8D0A958D16",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "1000000000",
 "gas_used": "9097685",
 "tags": [
  {
   "key": "action",
   "value": "issueBurnAny"
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
