# hashgardcli stake edit-validator

## 介绍

编辑已有验证人账户，包括佣金比率，验证人节点名称以及其他描述信息等

## 用法

```
hashgardcli stake edit-validator [flags] 
```

打印帮助信息
```
hashgardcli stake edit-validator --help
```

## 特有flags

| 名称                | 类型   | 是否必填 | 默认值   | 功能描述         |
| --------------------| -----  | -------- | ------------------ | ------------------------------------------------------------------- |
| --commission-rate   | float  | false    | 0.0                | 佣金比率 |
| --details           | string | false    | "[do-not-modify]"  | 验证人节点的详细信息 |
| --identity          | string | false    | "[do-not-modify]"  | 身份签名 |
| --moniker           | string | false    | "[do-not-modify]"  | 验证人名称 |
| --website           | string | false    | "[do-not-modify]"  | 网址 |

## 示例

```shell
hashgardcli stake edit-validator \
--from=hashgard \
--chain-id=hashgard \
--website=http://hashgard.com \
--details=hashgard_is_great
```
