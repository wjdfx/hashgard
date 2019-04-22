# hashgardcli issue burn-from-off

## Description
Token owner disable the permission for token holder to burn the token.
```
Note: irreversible
```
## Usage
```
 hashgardcli issue burn-from-off [issue-id] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example

### Disable burning function
```shell
hashgardcli issue burn-from-off coin174876e800 --from foo -o=json
```
Burning by the owner function will be disabled after you entering the correct password.
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
