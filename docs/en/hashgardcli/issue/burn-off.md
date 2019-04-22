# hashgardcli issue burn-off

## Description
Token owner disable the burning token by owner function.
```
Note: irreversible
```
## Usage
```
 hashgardcli issue burn-off [issue-id] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example
### Disable burning function
```shell
hashgardcli issue burn-off coin174876e800 --from foo -o=json
```
Burning token by onwer function will be disabled after you entering the correct password. 
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
   "value": "issue_burn_off"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
