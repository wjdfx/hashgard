 # hashgardcli issue burn-from

## Description
Token holder could burn one's own token under the condition of token owner did not disable this function. 
## Usage
```
 hashgardcli issue burn-from [issue-id] [amount] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example
### Burn token
```shell
hashgardcli issue burn-from coin174876e800 88888 --from=foo -o=json
```
You token will be burned after entering correct password.
```txt
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
