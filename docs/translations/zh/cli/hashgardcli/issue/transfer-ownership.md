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
 hashgardcli issue transfer-ownership coin174876e802 gard1lgs73mwr56u2f4z4yz36w8mf7ym50e7myrqn65 --from
```
输入正确的密码之后，你的该代币的Onwer就完成了转移。
```txt
{
   Height: 3199
  TxHash: 3438C2C4F054730CD02FC30C408B3DA558CE9C5CC99810F83406DB1D41708CC9
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 26680
  Tags: 
    - action = issue_transfer_ownership
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```
