# hashgardcli issue finish-minting 

## Description
Token owner set the token issued by the owner as finished minting.
```
Note: irreversible
```
## Usage
```
 hashgardcli issue finish-minting  [issue-id] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example
### Finished minting
```shell
hashgardcli issue finish-minting  coin174876e800 --from foo -o=json
```
Your will not be able to mint the token after entering the correct password. 
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
   "value": "issue_finish_minting"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
