# hashgardcli bank broadcast

## 描述
离线广播事务


## 使用方式

```
hashgardcli bank broadcast <file> [flags]
```

 
## 标志

| 命令，速记       | 类型   | 是否必须 | 默认值                | 描述                                                         |
| ---------------- | ------ | -------- | --------------------- | ------------------------------------------------------------ |
| -h, --help       |        | 否       |                       | 打印帮助                                                     |
| --chain-id       | String | 否       |                       | tendermint 节点网络ID                                        |
| --account-number | int    | 否       |                       | 账户数字用于签名通证发送                                     |
| --amount         | String | 是       |                       | 需要发送的通证数量，比如10apple                               |
| --async          |        | 否       | True                  | 异步广播传输信息                                             |
| --dry-run        |        | 否       |                       | 忽略--gas 标志 ，执行仿真传输，但不广播。                    |
| --fee            | String | 是       |                       | 设置传输需要的手续费                                         |
| --from           | String | 是       |                       | 用于签名的私钥名称                                                                     |
| --gas            | String | 否       | 20000                 | 每笔交易设定的gas限额; 设置为“simulate”以自动计算所需气体    |
| --gas-adjustment | Float  | 否       | 1                     | 调整因子乘以传输模拟返回的估计值; 如果手动设置气体限制，则忽略该标志 |
| --generate-only  |        | 否       |                       | 创建一个未签名的传输并写到标准输出中。                       |
| --indent         |        | 否       |                       | 在JSON响应中增加缩进                                         |
| --json           |        | 否       |                       | 以json格式返回输出                                           |
| --memo           | String | 否       |                       | 传输中的备注信息                                             |
| --print-response |        | 否       |                       | 返回传输响应 (仅仅当 async = false时有效)                    |
| --sequence       | Int    | 否       |                       | 等待签名传输的序列号。                                       |                   |
| --ledger         | String | 否       |                       | 使用一个联网的分账设备                                       |
| --node           | String | 否       | tcp://localhost:26657 | <主机>:<端口> 链上的tendermint rpc 接口。                    |
| --trust-node     | String | 否       | True                  | 不验证响应的证明                                             |





## 全局标志

| 命令，缩写            | 默认值         | 描述                                | 是否必须 |
| --------------------- | -------------- | ----------------------------------- | -------- |
| -e, --encoding string | hex            | 字符串二进制编码 (hex ,b64,btc ) | 否       |
| --home string         | $HOME/.hashgardcli | 配置和数据存储目录                  | 否       |
| -o, --output string   | text           | 输出格式 (text,json)              | 否       |
| --trace               |                | 出错时打印完整栈信息                | 否       |



## 例子

### 查询账户信息 

```
hashgardcli bank broadcast <file> [flags]

```

​    



​           
