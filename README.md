# Hashgard
hashgard 公链项目，基于 [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) 开发

## Required
[Go 1.10+](https://golang.org/dl/)

## Install
请先确保已经安装`Go`, 并且设置了 `$GOPATH`, `$GOBIN`, `$PATH` 这几个环境变量。

请将 Hashgard 项目放在指定目录，切换至 `master` 分支，进行安装：

```
mkdir -p $GOPATH/src/github.com/hashgard
cd $GOPATH/src/github.com/hashgard
git clone https://github.com/hashgard/hashgard
cd hashgard && git checkout master
make get_tools && make get_vendor_deps && make install
```

NOTE: 如果无法正常下载依赖包，请设置合适的代理

检查安装是否成功:

```
$hashgard --help
$hashgardcli --help
```

## Explorer
[hashgard explorer](https://github.com/hashgard/gardplorer)

## Testnets
[hashgard testnets](https://github.com/hashgard/testnets)