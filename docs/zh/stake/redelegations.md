# hashgardcli stake redelegations

## 描述

基于委托者地址的所有重新委托记录查询

## 用法

```
hashgardcli stake redelegations [delegator-address] [flags]
```
打印帮助信息
```
hashgardcli stake redelegations --help
```

## 示例

基于委托者地址的所有重新委托记录查询
```
hashgardcli stake redelegations [delegator-address] --trust-node
```

运行成功以后，返回的结果如下：

```json
[
  {
    "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "validator_src_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
    "validator_dst_addr": "gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq",
    "creation_height": "24800",
    "min_time": "2018-12-21T02:49:44.731658304Z",
    "initial_balance": {
      "denom": "apple",
      "amount": "8"
    },
    "balance": {
      "denom": "apple",
      "amount": "8"
    },
    "shares_src": "8.9100000000",
    "shares_dst": "8.0000000000"
  }
]
```