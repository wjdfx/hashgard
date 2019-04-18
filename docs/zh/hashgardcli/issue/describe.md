# hashgardcli issue describe

## 描述
Owner可以对自己代币进行补充描述，描述文件使用json格式。可以自定义各种属性，也可以使用官方推荐的模板。
## 使用方式
```
 hashgardcli issue describe [issue-id] [description-file] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 给代币设置描述
```shell
hashgardcli issue describe coin155556750600 path/description.json --from=foo -o=json
```
#### 模板
```
{
    "organization":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "description":"这是一个牛逼的项目" 
}
```
输入正确的密码之后，你的该代币的Onwer就完成了转移。
```txt
{
 "height": "3598",
 "txhash": "FA9DB4CFD21E70E16CB75332458004E2A296012FABF0B32018FC7E2A1E02EEC0",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000",
 "gas_used": "9086563",
 "tags": [
  {
   "key": "action",
   "value": "issueTransferOwnership"
  },
  {
   "key": "issue-id",
   "value": "coin155556750600"
  }
 ]
}
```
