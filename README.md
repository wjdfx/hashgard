# Hashgard
Public blockchain, based on [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) development.

## Required
[Go 1.10+](https://golang.org/dl/)

## Install
Please make sure have already installed `Go` correctly, and set environment variables : `$GOPATH`, `$GOBIN`, `$PATH`.

Put the Hashgard project in the specific path，switch to `master` branch，download related dependencies, then make install:
```
mkdir -p $GOPATH/src/github.com/hashgard
cd $GOPATH/src/github.com/hashgard
git clone https://github.com/hashgard/hashgard
cd hashgard && git checkout master
make get_tools && make get_vendor_deps && make install
```

Check if the installation is successful:
```
$hashgard --help
$hashgardcli --help
```

## Explorer
[hashgard explorer](https://github.com/hashgard/gardplorer)

## Testnets
[hashgard testnets](https://github.com/hashgard/testnets)