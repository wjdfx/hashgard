# hashgardcli issue list

## 描述

查询指定issue-id值的发行的币的信息。

## 使用方式

```
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

## 例子

### 返回列表

```shell
hashgardcli issue list --limit 1 -o=json
```

```txt
[
 {
  "issue_id": "coin155556750600",
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