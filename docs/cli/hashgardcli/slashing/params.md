# hashgardcli slashing params

## Description

Query genesis parameters for the slashing module

## Usage

```
hashgardcli slashing params [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```
hashgardcli slashing params --trust-node
```

The result is as follows：

```txt
Slashing Params:
  MaxEvidenceAge:          2m0s
  SignedBlocksWindow:      100
  MinSignedPerWindow:      0.500000000000000000
  DowntimeJailDuration:    10m0s
  SlashFractionDoubleSign: 0.050000000000000000
  SlashFractionDowntime:   0.010000000000000000
```
