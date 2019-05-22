# hashgardcli keys add

## Description

Derive a new private key and encrypt to disk.

## Usage

```
hashgardcli keys add <name> [flags]
```

## Flags

| Name, shorthand      | type  | Required  | Default| Description                                                         |
| --------------- | --------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| --account       | int | No| 0 | Account number for HD derivation                      |
| --index         | int | No| 0 | Add indent to JSON response          |
| --interactive | string | No| "" | Interactively prompt user for BIP39 passphrase and mnemonic|
| --multisig | strings | No| "" |  Construct and store a multisig public key (implies --pubkey)|
| --multisig-threshold | int | No| 1 |K out of N required signatures. For use in conjunction with --multisig|
| --no-backup     | bool | No| false |Don't print out seed phrase (if others are watching the terminal) |
| --nosort | bool | No| false |Keys passed to --multisig are taken in the order they're supplied |
| --pubkey | string | No| "" | Parse a public key in bech32 format and save it to disk|
| --recover       | string | No| "" |  Provide seed phrase to recover existing key instead of creating            |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Create a new key

```shell
hashgardcli keys add MyKey
```

You'll be asked to enter a password for your key, note: password must be at least 8 characters.

```txt
Enter a passphrase for your key:
Repeat the passphrase:
```

After that, you're done with creating a new key, but remember to backup your seed phrase, it is the only way to recover your account if you ever forget your password or lose your key.
```txt
NAME:	TYPE:	ADDRESS:						PUBKEY:
MyKey	local	gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet	gardpub1addwnpepqvu549hgyhnxlveqmtdn2xywygxpgzcsqefxur47zkz4e0e9x67hvjr6r6p
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

oval green shrug term already arena pilot spirit jump gain useful symbol hover grid item concert kiss zero bleak farm capable peanut snack basket
```

The 24 words above is a seed phrase just for example, DO NOT use it in production.


### Recover an existing key

If you forget your password or lose your key, or you wanna use your key in another place, you can recover your key by your seed phrase.

```txt
hashgardcli keys add MyKey --recover
```

You'll be asked to enter a new password for your key, and enter the seed phrase. Then you get your key back.

```txt
Enter a passphrase for your key:
Repeat the passphrase:
Enter your recovery seed phrase:
```

