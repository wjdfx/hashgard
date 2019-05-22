# 提交在线治理
- 任何用户都可以存入一些令牌来发起提案。存款达到一定值min_deposit后，进入投票期，否则将保留存款期。其他人可以在存款期内存入提案。一旦存款总额达到min_deposit，输入投票期。但是，如果冻结时间超过存款期间的max_deposit_period，则提案将被关闭。
- 进入投票期的提案只能由验证人和委托人投票。未投票的代理人的投票将与其验证人的投票相同，并且投票的代理人的投票将保留。到达“voting_period”时，票数将被计算在内。

## 如何提交在线治理？
您可以通过使用一下命令来提交治理提议：
```bash
hashgardcli gov submit-proposal \
     --chain-id=${chain-id} \
     --title=${proposal_title} \
     --type=${typ} \
     --description=${desription} \
     --from=${wallet_name}
```
提交区块链治理提议以及发起提议所涉及的初始保证金，其中提议的类型包括Text/ParameterChange/SoftwareUpgrade这三种类型。


若提案没有到达投票期，您还需要通过[抵押存款](Deposit.md)来激活提案


提案到达投票期之后，您可以进行[提案投票](Vote.md)
