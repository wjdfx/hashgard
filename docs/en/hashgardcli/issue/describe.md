# hashgardcli issue describe

## Description
Owner can add Usage of the token issued by owner, and the Usage need to be in json format. You can customize preferences or use recommended templates.
## Usage
```
 hashgardcli issue describe [issue-id] [description-file] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example

### Add Usage to the token
```shell
hashgardcli issue describe coin174876e800 path/Usage.json --from=foo -o=json
```
#### Templates
```
{
    "organization":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "Usage":"New Generation Digital Finance Public Chain" 
}
```
Your Usage of the token will be updated after entering the correct password.
```
{
 "type": "issue/CoinIssueInfo",
 "value": {
  "issue_id": "coin155547350023",
  "issuer": "gard1avx50wdu54rw6fh75hsvuzm8uy0ue6myxts029",
  "owner": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc",
  "issue_time": "2019-04-17T05:11:20.912597175Z",
  "name": "foocoin",
  "symbol": "qu8wh5",
  "total_supply": "100000000",
  "decimals": "18",
  "Usage": "{\"organization\":\"Hashgard\",\"website\":\"https://www.hashgard.com\",\"logo\":\"https://cdn.hashgard.com/static/logo.2d949f3d.png\",\"Usage\":\"New Generation Digital Finance Public Chain\"}",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 }
}
```
