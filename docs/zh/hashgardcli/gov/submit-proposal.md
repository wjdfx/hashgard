# hashgardcli gov submit-proposal

## 描述

提交区块链治理提议以及发起提议所涉及的初始保证金，其中提议的类型包括Text/ParameterChange/SoftwareUpgrade这三种类型。

## 使用方式

```
hashgardcli gov submit-proposal [flags]
```
打印帮助信息:

```
hashgardcli gov submit-proposal --help
```
## 标志

| 名称, 速记        | 默认值                      | 描述                                                                                                                                                 | 是否必须  |
| ---------------- | -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| --deposit        |                            | [string] 提议的保证金                                                                                                                         |          |
| --description    |                            | [string] 提议的描述                                                                                                           | Yes      |
| --key            |                            | 参数的键名称                                                                                                                        |          |
| --op             |                            | [string] 对参数的操作                                                                                                             |          |
| --param          |                            | [string] 提议参数,例如: [{key:key,value:value,op:update}]                                                                                 |          |
| --path           |                            | [string] param.json文件路径                                                                                                                      |          |
| --title          |                            | [string] 提议标题                                                                                                                           | Yes      |
| --type           |                            | [string] 提议类型,例如:Text/ParameterChange/SoftwareUpgrade                                                                            | Yes      |

## 例子

### 提交一个'Text'类型的提议

```shell
hashgardcli gov submit-proposal --chain-id=hashgard --title="notice proposal" --type=Text --description="a new text proposal" --from=hashgard

```

输入正确的密码之后，你就完成提交了一个提议，需要注意的是要记下你的提议ID，这是可以检索你的提议的唯一要素。

```txt
Committed at block 14932 (tx hash: 049477583479EB543F1EB48D02C3D705CFAF6A2DA0CA03EA67FCC98865D8EB25, response: {Code:0 Data:[1 1] Log:Msg 0:  Info: GasWanted:200000 GasUsed:46248 Tags:[{Key:[97 99 116 105 111 110] Value:[115 117 98 109 105 116 95 112 114 111 112 111 115 97 108] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[112 114 111 112 111 115 101 114] Value:[103 97 114 100 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 102 52 119 119 101 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[112 114 111 112 111 115 97 108 45 105 100] Value:[1 1] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}] Codespace: XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0})
```

### 提交一个'ParameterChange'类型的提议

```shell
hashgardcli gov submit-proposal --chain-id=hashgard --title="update MinDeposit proposal" --param='{"key":"Gov/gov/DepositProcedure","value":"{\"min_deposit\":[{\"denom\":\"gard\",\"amount\":\"1000\"}],\"max_deposit_period\":20}","op":"update"}' --type=ParameterChange --description="a new parameter change proposal" --from=hashgard
```

提交之后，您完成了提交新的“ParameterChange”提议。
更改参数的详细信息（通过查询参数获取参数，修改它，然后在“操作”上添加“更新”，使用方案中的更多详细信息）和其他类型的提议字段与文本提议类似。
注意：在这个例子中, --path 和 --param 不能同时为空。

### 提交一个'SoftwareUpgrade'类型的提议

```shell
hashgardcli gov submit-proposal --chain-id=hashgard --title="hashgard" --type=SoftwareUpgrade --description="a new software upgrade proposal" --from=hashgard
```

在这种场景下，提议的 --title、--type 和--description参数必不可少，另外你也应该保留好提议ID，这是检索所提交提议的唯一方法。


如何查询提议详情？

请点击下述链接：

[proposal](proposal.md)

[proposals](proposal.md)

