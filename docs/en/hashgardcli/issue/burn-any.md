# hashgardcli issue burn-any

## Description
Token owner has the right to burn token owned by any holder before disable the Token Burn function.

## Usage
```
 hashgardcli issue burn-any [issue-id] [address] [amount] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example

### Burn token
```shell
hashgardcli issue burn-any coin174876e800 gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye 888 --from=foo -o=json
```
You will burn the token owned by desgniated holder after entering the correct password.
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
   "value": "issue_burn_any"
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
