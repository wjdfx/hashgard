# 提案投票
提案正式激活之后，您可以对此提案投票

您可以通过查看提案的详情，来确定是否是进入了投票期：
```bash
hashgardcli gov proposal ${proposal_id} --trust-node
```
得到的结果中，``` proposal_status ```为此提案的状态

若在voting状态时，可通过如下操作进行投票：

```shell
hashgardcli gov vote ${proposal_id} ${option} --chain-id=hashgard --from hashgard
```
需要指定投票的id```${proposal_id}```，和投票的选项```${option}```，可选项包括: Yes/No/NoWithVeto/Abstain

