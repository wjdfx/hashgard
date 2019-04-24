# hashgardcli issue describe

## 描述
Owner可以对自己代币进行补充描述，描述文件使用不超过1MB的json格式。可以自定义各种属性，也可以使用官方推荐的模板。
## 使用方式
```
 hashgardcli issue describe [issue-id] [description-file] [flags]
```
## Global Flags

 ### 参考：[hashgardcli](../README.md)

## 例子
### 给代币设置描述
```shell
hashgardcli issue describe coin174876e800 path/description.json --from=foo -o=json
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
输入正确的密码之后，你的该代币的描述就设置成功了。
```txt
{
 "height": "17941",
 "txhash": "196C1FC96A604D34B7B7815C2425458BFAC1512D9255D5845A540F50D614F6F0",
 "data": "ERBjb2luMTU1NTQ3MzUwMDIz",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "1000000000000",
 "gas_used": "9093272",
 "tags": [
  {
   "key": "action",
   "value": "issue_description"
  },
  {
   "key": "issue-id",
   "value": "coin155547350023"
  }
 ]
}
```
### 查询发行信息
```shell
hashgardcli issue query coin155547350023 -o=json
```
最新的描述信息就生效了
```
{
 "type": "issue/CoinIssueInfo",
 "value": {
  "issue_id": "coin155547350023",
  "issuer": "gard1avx50wdu54rw6fh75hsvuzm8uy0ue6myxts029",
  "owner": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc",
  "issue_time": "2019-04-17T05:11:20.912597175Z",
  "name": "foocoin",
  "symbol": "qu8wh5",
  "total_supply": "100000000",
  "decimals": "18",
  "description": "{\"organization\":\"Hashgard\",\"website\":\"https://www.hashgard.com\",\"logo\":\"https://cdn.hashgard.com/static/logo.2d949f3d.png\",\"description\":\"这是一个牛逼的项目\"}",
  "burning_off": false,
  "burning_from_off": false,
  "burning_any_off": false,
  "minting_finished": false
 }
}
```