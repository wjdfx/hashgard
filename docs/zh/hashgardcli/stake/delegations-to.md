# hashgardcli stake delegations-to

## 描述

查询对一个验证人进行的所有委托

## 用法

```
hashgardcli stake delegations-to [validator-addr] [flags]
```

打印帮助信息
```
hashgardcli stake delegations-to --help
```

## 示例

查询验证人的所有委托
```
hashgardcli stake delegations-to gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --trust-node
```

运行成功以后，返回的结果如下：

```txt
[
  {
    "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "validator_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
    "shares": "99.0000000000"
  }
]
```
