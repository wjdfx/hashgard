# hashgardcli issue transfer-ownership

## Description
Token owner could transfer the ownership to new account, use "send" order to transfer tokens. 
## Usage
```
 hashgardcli issue transfer-ownership [issue-id] [to_address] [flags]
```
## Global Flags
 ### [hashgardcli](../README.md)
## Example
### Transfer owner
```shell
hashgardcli issue transfer-ownership coin174876e800 gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from=foo -o=json
```

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
