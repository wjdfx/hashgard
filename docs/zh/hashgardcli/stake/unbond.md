# hashgardcli stake unbond

## 介绍

从一个验证人解绑股票

## 用法

```
hashgardcli stake unbond [flags]
```

打印帮助信息

```
hashgardcli stake unbond --help
```

## 特有flags

| 名称                | 类型   | 是否必填 | 默认值   | 功能描述         |
| --------------------| -----  | -------- | -------- | ------------------------------------------------------------------- |
| --validator | string | true     | ""       | 验证人地址 |
| --shares-amount     | float  | false    | 0.0      | 解绑的share数量，正数 |
| --shares-fraction | float  | false    | 0.0      | 解绑的比率，0到1之间的正数 |

用户可以用`--shares-amount`或者`--shares-percent`指定解绑定的token数量，这两个参数不可同时使用。

## 示例

```
hashgardcli stake unbond --validator=gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --shares-fraction=0.1 --from=hashgard --chain-id=hashgard

```
