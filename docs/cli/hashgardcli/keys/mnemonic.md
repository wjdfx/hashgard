# hashgardcli keys mnemonic

## 描述

通过读取系统熵来创建24个单词组成的bip39助记词，有时称为种子短语。如果需要传递自定义的熵，请使用--unsafe-entropy参数。

## 使用方式

```
hashgardcli keys mnemonic [flags]
```

## Flags

| 名称, 速记        | 默认值     | 描述                                                                          | 是否必须  |
| ---------------- | --------- | ----------------------------------------------------------------------------- | -------- |
| --unsafe-entropy |           | 提示用户提供自己的熵，而不是依赖于系统生成                                          |          |

## 例子

### 创建指定密钥的助记词

```shell
hashgardcli keys mnemonic MyKey
```

执行命令就可以得到24个单词组成的助记词。为了安全考虑，请注意保存，比如将单词手抄纸并将纸张妥善保存。

```txt
police possible oval milk network indicate usual blossom spring wasp taste canal announce purpose rib mind river pet brown web response sting remain airport
```