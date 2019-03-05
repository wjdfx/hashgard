#!/usr/bin/env bash

# ------------------------------------------------------------------------------
#
# Parameters which might be changed while upgrading.
#
# ------------------------------------------------------------------------------
CHAIN_ID=sif-3000
KEY_NAME=root
KEY_PASSWORD=12345678
COIN_GENESIS=10000000gard
COIN_DELEGATE=100000gard

# ------------------------------------------------------------------------------
#
# Init chain config fot test with only one node.
#
# ------------------------------------------------------------------------------
if [[ ! -d ~/.hashgard ]]; then
    # Config global client settings
    hashgardcli config chain-id ${CHAIN_ID}
    hashgardcli config trust-node true

    # Create private key for first delegation
    echo "${KEY_PASSWORD}" | hashgardcli keys add ${KEY_NAME}

    # Init hashgard chain
    hashgard init --moniker hashgard --chain-id ${CHAIN_ID}
    hashgard add-genesis-account ${KEY_NAME} ${COIN_GENESIS}
    echo "${KEY_PASSWORD}" | hashgard gentx --name ${KEY_NAME} --amount ${COIN_DELEGATE}
    hashgard collect-gentxs
fi

hashgard start
