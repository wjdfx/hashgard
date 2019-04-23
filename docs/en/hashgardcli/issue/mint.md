# hashgardcli issue mint

## Description
Owner can mint addtional tokens issued by the owner under the conditional that the token mint function was not disabled. Token will be minted to owner's account or owner's designated account.
## Usage
```
 hashgardcli issue mint [issue-id] [amount] [flags]
```
| Name      |   Type   |    Required            |    Default             | Details                                                                                                                                                  |
| -----------------  | -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| --to                  | string | false | "" | Mint  to the specified account address                                                 |

## Global Flags

 ### [hashgardcli](../README.md)

## Example

### Mint to
```shell
hashgardcli issue mint coin174876e800 88888 --to=gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from=foo -o=json
```
Token minting will be completed after entering the correct passward. 
```txt
{
 "height": "3896",
 "txhash": "EFA5BC8F97DB4697037A7E85E5A85237A57F819860E5B6595D33AC412F25DEF6",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "1000000000",
 "gas_used": "18989223",
 "tags": [
  {
   "key": "action",
   "value": "issue_mint"
  },
  {
   "key": "recipient",
   "value": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```