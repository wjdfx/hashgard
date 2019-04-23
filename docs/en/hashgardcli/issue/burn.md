# hashgardcli issue burn

## Description
Token owner could burn the token issued by the owner under the condition that the burning function were enabled.
## Usage
```
 hashgardcli issue burn [issue-id] [amount] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example

### Burn token
```shell
hashgardcli issue burn coin174876e800 88888 --from=foo -o=json
```
Your token will be burned after entering the correct password. 
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
