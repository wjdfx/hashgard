# hashgardcli gov query-deposits

## 描述

查询保证金的充值明细

## 使用方式

```
  hashgardcli gov query-deposit [proposal-id] [depositer-address] [flags]

```
打印帮助信息:

```
hashgardcli gov query-deposit --help
```

## 例子

### 查询充值保证金

```shell
hashgardcli gov query-deposit 1 gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node

```

通过指定提议、指定存款人查询保证金充值详情，得到结果如下：

```txt
{
  "depositor": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
  "proposal_id": "1",
  "amount": [
    {
      "denom": "apple",
      "amount": "50"
    }
  ]
}
```
