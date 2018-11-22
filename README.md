# Hashgard
hashgard 公链项目，基于 [cosmos-sdk](https://github.com/hashgard/cosmos-sdk] 开发


## Required
[Go 1.10+](https://golang.org/dl/)


## Install Hashgard
请先确保已经安装`Go`, 并且设置了 `$GOPATH`, `$GOBIN`, `$PATH` 这几个环境变量。

请将 Hashgard 项目放在指定目录，切换至 `master` 分支，进行安装：

```
mkdir -p $GOPATH/src/github.com/hashgard
cd $GOPATH/src/github.com/hashgard
git clone http://gitlab.hashgard.com/public-chain/hashgard.git
cd hashgard && git checkout master
make get tools && make get_vendor_deps && make install
```

NOTE: 如果无法正常下载依赖包，请设置合适的代理

然后检查安装是否成功:

```
$hashgard help
$hashgardcli help
```


## Run a Full Node


## Testnet
[hashgard testnet](http://gitlab.hashgard.com/public-chain/testnet)