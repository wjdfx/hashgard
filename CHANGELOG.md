# Changelog


--------------------------

## 0.3.0

### BREAKING CHANGES

+ Hashgard REST API (hashgardlcd)
	+ now default mode is insecure, use `--tls` flag to use https.
	+ `tx/sign` endpoint now expects `BaseReq` fields as nested object.
	+ all endpoints renamed from `/stake` -> `/staking`
	+ rename:
		+ `LooseTokens` -> `NotBondedTokens`
		+ `Validator.UnbondingMinTime` -> `Validator.UnbondingCompletionTime`
		+ `Delegation` -> `Value` in `MsgCreateValidator` and `MsgDelegate`
		+ `MsgBeginUnbonding` -> `MsgUndelegate`

+ Hashgard CLI (hashgardcli)
	+ Rename chain_id and trust_node to chain-id and trust-node respectively in config file.
	+ Remove unimplemented `init` command.
	+ `version` command output extra latest commit and build machine info.

+ Hashgard (hashgard)
	+ use Storekeys of each module instead of literals.
	+ the `--gas` flag now takes auto instead of simulate .
	+ `version` command output extra latest commit and build machine info.
	+ `tendermint`'s show-validator and `show-address` `--json` flags removed in favor of `--output-format=json`.

+ Tendermint
	+ upgrade tendermint from v0.27.3 to v0.29.0

+ Cosmos SDK
	+ upgrade cosmos-sdk from v0.27.3 to v0.29.0
	+ rename module `stake` -> `staking`.
	+ rename `LooseTokens` -> `NotBondedTokens`
	+ [staking] Validator power type from `Dec` -> `Int`.
	+ [gov] Remove redundant action tag
	+ Sanitize sdk.Coin denom. Coins denoms are now case insensitive, i.e. 100fooToken equals to 100FOOTOKEN.
	+ Fix infinite gas meter utilization during aborted ante handler executions.
	+ Increase decimal precision to 18


### FEATURES

+ Hashgard REST API (hashgardlcd)
	+ add support for fees on transactions.
	+ add a custom memo on transactions.
	+ implement `/gov/proposals/{proposalID}/proposer` to query for a proposal's proposer.

+ Hashgard CLI (hashgardcli)
	+ new `keys add --multisig` flag to store multisig keys locally.
	+ new `bank sign --multisig` flag to enable multisig mode.
	+ add new `bank multisign` command to generate multisig signatures for transactions generated offline
	+ add `distribution params`, `distribution outstanding-rewards`, `distribution commission`, `distribution slashes`, `distributionrewards` commands to query more info.
	+ add new `slashing params` command to query the current slashing parameters.
	+ add new `gov param`, `gov proposer` commands to query more relative info.

+ Hashgard (hashgard)
	+ added queriers for querying a single redelegation he governance, slashing, and staking modules.


### BUG FIXES

+ Hashgard CLI (hashgardcli)
	+ Fix the bug in GetAccount when `len(res) == 0` or `err == nil`

+ Hashgard (hashgard)
	+ Correctly reset total accum update height and jailed-validator bond height / unbonding height on export-for-zero-height
	+ Fix unset governance proposal queues when importing state from old chain
	+ Fix `hashgard export` by resetting each validator's slashing period.


--------------------------

## 0.2.1

BREAKING CHANGES
* Cosmos-sdk
  * [cosmos-sdk] Now using cosmos-sdk 0.29.1
* Change the name of stake coin


--------------------------

## 0.2.0

BREAKING CHANGES
* Cosmos-sdk
  * [cosmos-sdk] Now using cosmos-sdk 0.29.0
* Tendermint
  * [tendermint] Now using tendermint 0.27.3

FEATURES
* Hashgard REST API(hashgardlcd)
  * [hashgardlcd] Split the LCD service from hashgardcli
* Other
  * [logjack] Introduced the logjack tool for saving logs w/ rotation

BUG FIXES
* Fixed the bug that did not recognize the msg of the distribution command
