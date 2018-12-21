# hashgardcli rest-server  

## 介绍

启动LCD（light-client守护程序），一个本地REST服务

## 用法

```
hashgardcli rest-server [subcommand]
```

打印子命令和参数

```
hashgardcli rest-server --help
```

## 子命令

| 名称                            | 功能                                                   |
| --------------------------------| --------------------------------------------------------------|
|chain-id|Tendermint节点的链ID|
|cors|设置可以发出CORS请求的域（*表示所有）|
|help|rest-server帮助文档|
|indent|添加缩进到JSON响应|
|insecure|不设置SSL/TLS层|
|laddr|要侦听的服务器的地址（默认为“tcp：// localhost：1317”）|
|max-open|最大打开连接数（默认为1000）|
|node|要连接的节点的地址（默认为“tcp：// localhost：26657”）|
|ssl-certfile|SSL证书文件的路径。如果未提供，将生成自签名证书|
|ssl-hosts|以逗号分隔的主机名和IP，用于生成证书|
|ssl-keyfile|密钥文件的路径;如果未提供证书文件，则忽略|
|trust-node|信任连接的完整节点（不验证响应的证据）|

### 示例
```
hashgardcli rest-server --laddr=tcp://0.0.0.0:1317 --chain-id=hashgard

```

你会得到类似于如下的结果：
```
 Starting Gaia Lite REST service...           module=rest-server 
 SHA256   Fingerprint=41:8B:1B:EA:72:DB:20:C7:82:DA:EB:F4:C4:6A:C4:52:7E:83:43:67:61:E0:9C:81:6F:75:49:D9:55:9D:28:55 module=rest-server 
 Starting RPC HTTPS server on [::]:1317 (cert: "/tmp/cert_332804498", key: "/tmp/key_208100809") module=rest-server

```

```cert```和 ```key```在不指定的情况下，每次的产生都是随机的名称，存放在tmp文件夹下