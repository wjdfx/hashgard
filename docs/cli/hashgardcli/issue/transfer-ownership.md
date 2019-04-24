# hashgardcli issue transfer-ownership

## 描述
Owner可以将代币的所有者转移到新的账户下，如要将代币也转移请使用 send 命令进行发送。
## 使用方式
```
 hashgardcli issue transfer-ownership [issue-id] [to_address] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 转移Owner
```shell
hashgardcli issue transfer-ownership coin174876e800 gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from=foo -o=json
```
输入正确的密码之后，你的该代币的Onwer就完成了转移。
```txt
{
 "height": "3598",
 "txhash": "FA9DB4CFD21E70E16CB75332458004E2A296012FABF0B32018FC7E2A1E02EEC0",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000",
 "gas_used": "9086563",
 "tags": [
  {
   "key": "action",
   "value": "issue_transfer_ownership"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
