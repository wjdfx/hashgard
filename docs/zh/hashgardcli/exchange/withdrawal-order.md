# hashgardcli exchange withdrawal-order

## 描述

某笔有效订单的seller可以撤销该笔订单，冻结的资金将返回其账户

## 使用方式

```bash
hashgardcli exchange withdrawal-order [order_id] [flags]
```

## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 示例

### 撤销订单

```bash
hashgardcli exchange withdrawal-order 2 \
    --from mykey \
    --chain-id hashgard
```

必须是订单的seller账户操作，输入正确的密码后，order_id为2的订单已经撤销。

```json
{
 "height": "7162",
 "txhash": "7FA99BF4C271E8145A8DB695B9B08883A58A46F1AA369D8A7F6002684FDBC06A",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "27608",
 "tags": [
  {
   "key": "action",
   "value": "withdrawal_order"
  },
  {
   "key": "order_id",
   "value": "2"
  },
  {
   "key": "seller",
   "value": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa"
  },
  {
   "key": "order_status",
   "value": "inactive"
  }
 ]
}
```