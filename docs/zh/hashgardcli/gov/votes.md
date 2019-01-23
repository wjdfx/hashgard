# hashgardcli gov votes

## 描述

查询指定提议的投票情况

## 使用方式

```
  hashgardcli gov votes [proposal-id] [flags]

```
打印帮助信息:

```
hashgardcli gov votes --help
```

## 例子

### Query votes

```shell
hashgardcli gov votes 1 --trust-node
```

通过指定的提议查询该提议所有投票者的投票详情。
 
```txt
[
  {
    "voter": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "proposal_id": "1",
    "option": "Yes"
  }
]
```
