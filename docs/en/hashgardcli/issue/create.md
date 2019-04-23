# hashgardcli issue create

## Description
Issue a new token
## Usage
```
hashgardcli issue create [name] [symbol] [total-supply] [flags]
```

## Flags

| Name               | Type   | Required | Default | Description                                              |
| ------------------ | ------ | -------- | ------ | ------------------------------------------------- |
| --decimals         | int    | false       | 18     | Token precision, default 18 digits, maximum 18 digits           |
| --burn-off         | string | false       | false  | Whether to close the Owner's ability to destroy the token             |
| --burn-from-off    | bool   | false       | ""     | Whether to close the ordinary account to destroy the function of the token         |
| --burn-any-off     | bool   | false       | false  | Whether to close the Owner can destroy the function of the token under any account |
| --minting-finished | bool   | false       | false  | Minting-finished                              |

## Global Flags

### [hashgardcli](../README.md)

## Example
### Issue a new token
```shell
hashgardcli issue create foocoin FOO 100000000 --from foo -o=json
```
You will complete the token issuance process after entering the correct password. Please write down your issue-id, this is the only way to check and operate your token. 
```txt
{
 "height": "3394",
 "txhash": "81D4B2054F741E901BE5A540DDA37BF53D1DEA16C94BF9E4BBDB1D1CD548DFA1",
 "data": "ERBjb2luMTU1NTU2NzUwNjAw",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "100000000000",
 "gas_used": "18994244",
 "tags": [
  {
   "key": "action",
   "value": "issue"
  },
  {
   "key": "recipient",
   "value": "gard1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc"
  },
  {
   "key": "issue-id",
   "value": "coin174876e800"
  }
 ]
}
```
```shell
hashgardcli bank account gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye
```

You will see a coin with the special name "(name)issue-id" in your currency list. For subsequent operations on the currency, please use the value of issue-id, including the transfer operation. Please also use the issue-id for the currency to be transferred.
```
{
 "type": "auth/Account",
 "value": {
  "address": "gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye",
  "coins": [
   {
    "denom": "mycoin(coin155548903200)",
    "amount": "9999999998999858889"
   },
   {
    "denom": "foocoin(coin174876e800)",
    "amount": "100000000"
   },
   {
    "denom": "gard",
    "amount": "1010000000"
   }
  ],
  "public_key": {
   "type": "tendermint/PubKeySecp256k1",
   "value": "A/rSPb+egaljwS1XGSSFKpaFkfjFzLWJFmtUdAlaQpLl"
  },
  "account_number": "1",
  "sequence": "11"
 }
}
```