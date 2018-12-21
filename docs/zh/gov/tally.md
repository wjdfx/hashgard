# hashgardcli gov tally

## 描述

查询指定提议的投票统计
 
## 使用方式

```
 hashgardcli gov tally [proposal-id] [flags]

```

打印帮助信息:

```
hashgardcli gov tally --help
```


## 例子

### 查询投票统计

```shell
hashgardcli gov tally 1 --trust-node

```

可以查询指定提议每个投票选项的投票统计。

```txt
{
  "yes": "89.1000000000",
  "abstain": "0.0000000000",
  "no": "0.0000000000",
  "no_with_veto": "0.0000000000"
}
```
