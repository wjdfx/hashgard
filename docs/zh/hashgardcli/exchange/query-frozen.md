# hashgardcli exchange query-frozen

## 描述

查看指定地址的冻结资金

## 使用方式

```bash
hashgardcli exchange query-frozen [address] [flags]
```

## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子

### 查询地址的冻结资金

```bash
hashgardcli exchange query-frozen gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa \
    --chain-id hashgard 
```

下面是地址gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa的冻结资金

```json
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
