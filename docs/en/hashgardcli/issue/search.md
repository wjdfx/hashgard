# hashgardcli issue search

## Description
Search token information according to token symbol.
## Usage
```
hashgardcli issue search [symbol] [flags]
```
## Global Flags

 ### [hashgardcli](../README.md)

## Example

###Search
```shell
hashgardcli issue search fo -o=json
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
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 }
]

```
