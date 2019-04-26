# hashgardcli issue list

## 描述

查询指定issue-id值的发行的币的信息。

## 使用方式

```bash
hashgardcli issue list [flags]
```

## Flags

| 名称             | 类型   | 是否必须 | 默认值 | 描述                  |
| ---------------- | ------ | -------- | ------ | --------------------- |
| --address        | string | 否       | ""     | （可选）Owner账号地址 |
| --limit          | int    | 否       | 30     | （可选）每次返回条数  |
| --start-issue-id | string | 否       | ""     | （可选）起始issue-id  |

## Global Flags

### 参考：[hashgardcli](../README.md)

## 示例

### 返回列表

```bash
hashgardcli issue list
```
```json
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

### 返回分页列表
```bash
hashgardcli issue list --limit 1 --start-issue-id coin174876e801 
```
```json
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
### 返回某一地址的列表

```bash
hashgardcli issue list --address=gard1sepa9tuxt238xj3jmvf98k6uk5z7wuwmm4f4mx
```
```json
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