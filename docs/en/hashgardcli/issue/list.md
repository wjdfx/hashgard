# hashgardcli issue list

## Description
Check token information acoording to issue ID.
## Usage
```
hashgardcli issue list [flags]
```
## Flags

|  Name     |  Type | Required | Default |  Description           |
| ---------------- | ------ | -------- | ------ | --------------------- |
| --address        | string | false       | ""     |  Owner address|
| --limit          | int    | false       | 30     |  Number of returns per time  |
| --start-issue-id | string | false       | ""     |  Returns the data after the issue-id  |

## Global Flags

### [hashgardcli](../README.md)

## Example
```shell
hashgardcli issue list --limit 1 -o=json
```
```txt
[
 {
  "issue_id": "coin174876e800",
  "issuer": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc",
  "owner": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc",
  "issue_time": "2019-04-18T06:05:01.378656183Z",
  "name": "foocoin",
  "symbol": "FOO",
  "total_supply": "99998224",
  "decimals": "18",
  "Usage": "",
  "burning_off": true,
  "burning_from_off": true,
  "burning_any_off": true,
  "minting_finished": true
 }
]
```