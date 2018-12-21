# hashgardcli tendermint tx

## 描述

在所有提交的块上匹配此txhash

## 用法

```
hashgardcli tendermint tx [hash] [flags]
```

## 标志

| 名称, 速记 | 默认值                    | 描述                                                             | 必需      |
| --------------- | -------------------------- | --------------------------------------------------------- | -------- |
| --chain-id    | 无 | [string] tendermint节点的链ID   | 是       |
| --node string     |   tcp://localhost:26657                         | 要连接的节点  |                                     
| --help, -h      |           无| 	下载命令帮助|
| --trust-node    | true                       | 信任连接的完整节点，关闭响应结果校验                                            |          |

## 例子

### 

```shell
 hashgardcli tendermint tx 6E44164FDF456BAED405A8AA8C2F8CD7E9DA1C7BB751616C50614D1F4773B245 --trust-node
```
将会得到如下类似的结果：
```
{"hash":"6E44164FDF456BAED405A8AA8C2F8CD7E9DA1C7BB751616C50614D1F4773B245","height":"9946","tx":{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/BeginRedelegate","value":{"delegator_addr":"gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet","validator_src_addr":"gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx","validator_dst_addr":"gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq","shares_amount":"11.0000000000"}}],"fee":{"amount":[{"denom":"","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AzlKlugl5m+zINrbNRiOIgwUCxAGUm4OvhWFXL8lNr12"},"signature":"rplv/JsF35H/bqlyniUv940M6HS6S0IDY8ynoHjKSc8V0je2nyGqunA66Tt//DkniWLenYxvm1a1SKNhxgWJGQ=="}],"memo":""}},"result":{"data":"CwigquzgBRCX8N0B","log":"Msg 0: ","gas_wanted":"200000","gas_used":"185386","tags":[{"key":"YWN0aW9u","value":"YmVnaW5fcmVkZWxlZ2F0ZQ=="},{"key":"ZGVsZWdhdG9y","value":"Z2FyZDFtM200bDZnNTc3NHFlNWpqOGN3bHlhc3VlMjJ5aDMyamY0d3dldA=="},{"key":"c291cmNlLXZhbGlkYXRvcg==","value":"Z2FyZHZhbG9wZXIxbTNtNGw2ZzU3NzRxZTVqajhjd2x5YXN1ZTIyeWgzMmptaHJ4Zng="},{"key":"ZGVzdGluYXRpb24tdmFsaWRhdG9y","value":"Z2FyZHZhbG9wZXIxeG40a3ZxODY3cmFwOHZrcndmbnA1bjJlbnR2cHEyYXZ0ZDB5dHE="},{"key":"ZW5kLXRpbWU=","value":"CwigquzgBRCX8N0B"}]}}
```
