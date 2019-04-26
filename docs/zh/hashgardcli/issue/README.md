# hashgardcli issue

## 描述

基于 Hashgard 公链发行自己的 HRC10 代币，我们提供增发、销毁、冻结、转移所有者，代币查询等链上原生服务。


>注：我们不限制你在公链账号里持有多少GARD币来发行自己的代币，只需足够支付每步操作的gas费用即可。


## 使用方式

```bash
hashgardcli issue [command]
```

打印子命令和参数

```bash
hashgardcli issue --help
```

## 相关命令

| 命令                                        | 描述                                 |
| ------------------------------------------- | ------------------------------------ |
| [create](create.md)                         | 发行一个新的代币。                   |
| [describe](describe.md)                     | 设置代币的描述信息。                 |
| [transfer-ownership](transfer-ownership.md) | 转移所有者。                         |
| [query](query.md)                           | 查询指定issue-id值的发行的币的信息。 |
| [mint](mint.md)                             | 增发                                 |
| [burn](burn.md)                             | Owner销毁代币                        |
| [burn-from](burn-from.md)                   | 持币者销毁代币                       |
| [burn-any](burn-any.md)                     | Owner销毁任意持币者的代币            |
| [finish-minting](finish-minting.md)         | 完成增发                             |
| [burn-off](burn-off.md)                     | 关闭Owner销毁功能                    |
| [burn-from-off](burn-from-off.md)           | 关闭持币者销毁功能                   |
| [search](search.md)                         | 根据代币符号搜索                     |