# hashgardcli issue

## Usage
Issue HRC10 token on Hashgard public chain. We provide additional issuance, burning, onwer transfer and information check services. 
```
Note: There is no requirement on GARD holdings under your account, as long as there is enough GARD to pay the gas fee for each process.
```

## Usage

```shell
hashgardcli issue [command]
```
Print help messages:
```
hashgardcli issue --help
```

## Available Commands

| 命令                                        | 描述                                 |
| ------------------------------------------- | ------------------------------------ |
| [create](create.md)                         | Issue a new token。                   |
| [describe](describe.md)                     | Add Usage to the token。                 |
| [transfer-ownership](transfer-ownership.md) | Transfer the ownership to new account。                         |
| [query](query.md)                           | Query token information。 |
| [mint](mint.md)                             | Mint                                 |
| [burn](burn.md)                             | Owner burn the token                        |
| [burn-from](burn-from.md)                   | Token holder burn the token                        |
| [burn-any](burn-any.md)                     | Token owner has the right to burn token           |
| [finish-minting](finish-minting.md)         | Finished minting                             |
| [burn-off](burn-off.md)                     | Disable the burning  by owner                  |
| [burn-from-off](burn-from-off.md)           | Disable the permission for token holder                  |
| [search](search.md)                         | Search                     |