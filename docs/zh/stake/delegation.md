# hashgardcli stake delegation

## 描述

基于委托者和验证者地址查询委托交易

## 用法

```
hashgardcli stake delegation [flags]
```
打印帮助信息
```
hashgardcli stake delegation --help
```
## 特有的flags

| 名称, 速记             | 默认值                      | 描述                                                                 | 必需     |
| --------------------- | -------------------------- | -------------------------------------------------------------------- | -------- |
| --address-delegator   |                            | [string] 委托者bech地址                                               | Yes      |
| --validator   |                            | [string] 验证者bech地址                                               | Yes      |

## 示例

### 查询验证者

```
hashgardcli stake delegation --address-delegator=gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --validator=gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx

```

运行成功以后，返回的结果如下：

```txt
Delegation 
Delegator: gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet
Validator: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Shares: 99.000000000
```
