# hashgardcli issue query

## Description
Query token information according to issue-id.
## Usage
```
hashgardcli issue query [issue-id] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example
### Check token information
```shell
hashgardcli issue query coin155548903200 -o=json
```
```txt
{
 "type": "issue/CoinIssueInfo",
 "value": {
  "issue_id": "coin155548903200",
  "issuer": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye",
  "owner": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye",
  "issue_time": "2019-04-17T08:17:05.109247975Z",
  "name": "mycoin",
  "symbol": "MY",
  "total_supply": "9999999999999838889",
  "decimals": "18",
  "Usage": "",
  "burning_off": true,
  "burning_from_off": true,
  "burning_any_off": true,
  "minting_finished": false
 }
}
```
