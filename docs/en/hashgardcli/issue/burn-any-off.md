# hashgardcli issue burn-any-off

## Description
Owner disabled the Burn Token function 
```
Note: non-reversible
```
## Usage
```
 hashgardcli issue burn-any-off [issue-id] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Examples

### Disable burning function
```shell
hashgardcli issue burn-any-off coin174876e800 --from foo -o=json
```
Burning function by token owner will be disabled once you entered the correct password.
```txt
{
 "height": "4917",
 "txhash": "B9F97B17BD986E9FA7CF41EF2FDF844E8D9582987D5183FB160A0FBDD6A7B045",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000",
 "gas_used": "9086502",
 "tags": [
  {
   "key": "action",
   "value": "issue_burn_any_off"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
