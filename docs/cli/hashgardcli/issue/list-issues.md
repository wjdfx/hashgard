# hashgardcli issue list-issues

## Description

查询用户在临界上发行的代币列表

## Usage

```
hashgardcli issue list-issues [flags]
```

## Flags

| 名称             | type  | Required  | Default| Description             |
| ---------------- | ------ | -------- | ------ | --------------------- |
| --address        | string | false    | ""     | （可选）Owner账号地址 |
| --limit          | int    | false    | 30     | （可选）每次返回条数  |
| --start-issue-id | string | false    | ""     | （可选）起始issue-id  |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### 返回列表

```shell
hashgardcli issue list-issues 
```
```txt
[
 {
  "issue_id": "coin174876e801",
  "issuer": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "owner": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "issue_time": "2019-04-19T06:23:00.748062914Z",
  "name": "joe234234",
  "symbol": "AAA",
  "total_supply": "1000000000000000",
  "decimals": "18",
  "description": "",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 },
 {
  "issue_id": "coin174876e800",
  "issuer": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "owner": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "issue_time": "2019-04-19T06:21:12.475597314Z",
  "name": "joe2342342344444",
  "symbol": "JOE",
  "total_supply": "1000000000000000",
  "decimals": "18",
  "description": "",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 }
]
```

```shell
hashgardcli issue list-issues --limit 1 --start-issue-id coin174876e801 
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
  "description": "",
  "burning_off": true,
  "burning_from_off": true,
  "burning_any_off": true,
  "minting_finished": true
 }
]
```

```shell
hashgardcli issue list-issues --address=gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx 
```
```txt
[
 {
  "issue_id": "coin174876e801",
  "issuer": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "owner": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "issue_time": "2019-04-19T06:23:00.748062914Z",
  "name": "joe234234",
  "symbol": "AAA",
  "total_supply": "1000000000000000",
  "decimals": "18",
  "description": "",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 },
 {
  "issue_id": "coin174876e800",
  "issuer": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "owner": "gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx",
  "issue_time": "2019-04-19T06:21:12.475597314Z",
  "name": "joe2342342344444",
  "symbol": "JOE",
  "total_supply": "1000000000000000",
  "decimals": "18",
  "description": "",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 }
]
```