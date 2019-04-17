# hashgardcli exchange query-frozen

## 描述

查看指定地址的冻结资金

## 使用方式

```
hashgardcli exchange query-frozen [address] [flags]
```

## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子

### 查询地址的冻结资金

```shell
hashgardcli exchange query-frozen gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa --chain-id hashgard -o=json --indent
```

下面是地址gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa的冻结资金

```txt
[
 {
  "denom": "apple",
  "amount": "33"
 },
 {
  "denom": "gard",
  "amount": "100"
 }
]
```
