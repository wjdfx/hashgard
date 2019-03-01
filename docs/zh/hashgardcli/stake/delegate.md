# hashgardcli stake delegate

## 介绍

发送委托交易

## 用法

```
hashgardcli stake delegate [validator-addr] [amount] [flags]
```

打印帮助信息
```
hashgardcli stake delegate --help
```

## 示例

```shell
hashgardcli stake delegate \
gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx \
10000gard \
--chain-id=hashgard --from=hashgard 
```

然后你将会得到如下消息：
```
Committed at block 9806 (tx hash: DAF8B140A281BB84444DE0A82AA9FC6FFCDEB8CDBB7D1BCFE5A04F1548AC96CA, response: {Code:0 Data:[] Log:Msg 0:  Info: GasWanted:200000 GasUsed:109083 Tags:[{Key:[97 99 116 105 111 110] Value:[100 101 108 101 103 97 116 101] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[100 101 108 101 103 97 116 111 114] Value:[103 97 114 100 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 102 52 119 119 101 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[100 101 115 116 105 110 97 116 105 111 110 45 118 97 108 105 100 97 116 111 114] Value:[103 97 114 100 118 97 108 111 112 101 114 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 109 104 114 120 102 120] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}] Codespace: XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0})
```