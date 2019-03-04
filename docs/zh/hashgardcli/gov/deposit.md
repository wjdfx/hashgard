# hashgardcli gov deposit

## 描述

充值保证金以激活提案

## 使用方式

```
hashgardcli gov deposit [proposal-id] [depositer-addr] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)
 
## 例子

### 充值保证金

```shell
 hashgardcli gov deposit  1 50gard --from=hashgard --chain-id=hashgard
```

输入正确的密码后，你就充值了50个gard用以激活提案的投票状态。

```txt
Committed at block 15016 (tx hash: 51526EB59CC44C6FBEE1EFD8BFB7A24780944A0A66618A86B4C535920BC69A11, response: {Code:0 Data:[] Log:Msg 0:  Info: GasWanted:200000 GasUsed:42850 Tags:[{Key:[97 99 116 105 111 110] Value:[100 101 112 111 115 105 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[100 101 112 111 115 105 116 111 114] Value:[103 97 114 100 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 102 52 119 119 101 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[112 114 111 112 111 115 97 108 45 105 100] Value:[1] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[118 111 116 105 110 103 45 112 101 114 105 111 100 45 115 116 97 114 116] Value:[1] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}] Codespace: XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0})
```

如何查询保证金充值明细？

请点击下述链接：

[query-deposit](query-deposit.md)

[deposits](deposits.md)
