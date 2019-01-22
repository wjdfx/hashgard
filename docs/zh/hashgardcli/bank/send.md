# hashgardcli bank send

## 描述

发送通证到指定地址 

## 使用方式

```
hashgardcli bank send --to=<account address> --from <key name> --chain-id=<chain-id> --amount=10gard
```

 

## 标志

| 命令，速记       | 类型   | 是否必须 | 默认值                | 描述                                                         |
| ---------------- | ------ | -------- | --------------------- | ------------------------------------------------------------ |
| -h, --help       |        | 否       |                       | 打印帮助                                                     |
| --chain-id       | String | 否       |                       | tendermint 节点网络ID                                        |
| --account-number | int    | 否       |                       | 账户数字用于签名通证发送                                     |
| --amount         | String | 是       |                       | 需要发送的通证数量，比如10hashgard                               |
| --async          |        | 否       | True                  | 异步广播传输信息                                             |
| --dry-run        |        | 否       |                       | 忽略--gas 标志 ，执行仿真传输，但不广播。                    |
| --fee            | String | 是       |                       | 设置传输需要的手续费                                         |
| --from           | String | 是       |                       | 用于签名的私钥名称                                           |
| --from-addr      | string | 否       |                       | 在generate-only模式下指定的源地址                            |
| --gas            | String | 否       | 20000                 | 每笔交易设定的gas限额; 设置为“simulate”以自动计算所需气体    |
| --gas-adjustment | Float  | 否       | 1                     | 调整因子乘以传输模拟返回的估计值; 如果手动设置气体限制，则忽略该标志 |
| --generate-only  |        | 否       |                       | 创建一个未签名的传输并写到标准输出中。                       |
| --indent         |        | 否       |                       | 在JSON响应中增加缩进                                         |
| --json           |        | 否       |                       | 以json格式返回输出                                           |
| --memo           | String | 否       |                       | 传输中的备注信息                                             |
| --print-response |        | 否       |                       | 返回传输响应 (仅仅当 async = false时有效)                    |
| --sequence       | Int    | 否       |                       | 等待签名传输的序列号。                                       |
| --to             | String | 是       |                       | Bech32 编码的接收通证的地址。                                |
| --ledger         | String | 否       |                       | 使用一个联网的分账设备                                       |
| --node           | String | 否       | tcp://localhost:26657 | <主机>:<端口> 链上的tendermint rpc 接口。                    |
| --trust-node     | String | 否       | True                  | 不验证响应的证明                                             |



## 全局标志

| 命令，速记            | 默认值         | 描述                                | 是否必须 |
| --------------------- | -------------- | ----------------------------------- | -------- |
| -e, --encoding string | hex            | 字符串二进制编码 (hex \|b64 \|btc ) | 否       |
| --home string         | /root/.hashgardcli | 配置和数据存储目录                  | 否       |
| -o, --output string   | text           | 输出格式 (text \|json)              | 否       |
| --trace               |                | 出错时打印完整栈信息                | 否       |



## 

## 例子

### 发送通证到指定地址 

```
 hashgardcli bank send --to=gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet  --from=hashgard  --chain-id=hashgard--amount=10gard
```

命令执行完成后，返回执行的细节信息

```
Committed at block 137 (tx hash: CDDE07D9858A638B837F677D3147648B7560BD77C0225539C5AF8785599D1805, response: {Code:0 Data:[] Log:Msg 0:  Info: GasWanted:200000 GasUsed:22425 Tags:[{Key:[97 99 116 105 111 110] Value:[115 101 110 100] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[115 101 110 100 101 114] Value:[103 97 114 100 49 117 104 97 117 121 116 104 116 101 116 57 48 101 119 116 117 121 52 48 118 52 104 114 121 109 108 113 102 53 110 52 53 119 99 120 99 120 99] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[114 101 99 105 112 105 101 110 116] Value:[103 97 114 100 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 102 52 119 119 101 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}] Codespace: XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0})

```
