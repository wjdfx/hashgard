# hashgardcli stake delegation

## 描述

基于委托者和验证者地址查询委托交易

## 用法

```
hashgardcli stake delegation [delegator-addr] [validator-addr] [flags]
```
打印帮助信息
```
hashgardcli stake delegation --help
```

## 示例

### 查询验证者

```shell
hashgardcli stake delegation gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu gardvaloper13nyheuxft7nylrmxmtzewdrs8ukh9r6eq4rya3 --trust-node
```

运行成功以后，返回的结果如下：

```txt
Delegation:
  Delegator: gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu
  Validator: gardvaloper13nyheuxft7nylrmxmtzewdrs8ukh9r6eq4rya3
  Shares:    100000000.000000000000000000
```
