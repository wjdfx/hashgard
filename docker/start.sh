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
INITIALIZED_FLAG="/initialized.flag"

function config_global_client_settings() {
    hashgardcli config chain-id ${CHAIN_ID}
    hashgardcli config trust-node true
}

# ------------------------------------------------------------------------------
#
# Initial Hashgard block chain full node
#
#   1. Init server's local configuration
#   2. Download preset configuration from github
#   3. Add flag to avoid reinitialization
#
# ------------------------------------------------------------------------------
function init_full_node() {
    if [[ -z ${MONIKER} ]]; then
        echo "Environment MONIKER must be set !"
        exit -1
    fi

    config_global_client_settings

    hashgard init --moniker ${MONIKER} --chain-id ${CHAIN_ID}
    cd /root/.hashgard/config
    rm -f config.toml genesis.json
    wget https://github.com/hashgard/testnets/blob/master/sif/sif-3000/config/config.toml
    wget https://github.com/hashgard/testnets/blob/master/sif/sif-3000/config/genesis.json
}

# ------------------------------------------------------------------------------
#
# Initial Hashgard block chain private single node
#
#   1. Generate account
#   2. Init server's local configuration
#   3. Assign coins to the account
#   4. Add account to genesis
#
# ------------------------------------------------------------------------------
function init_private_single() {
    # Create private key for first delegation
    echo "${KEY_PASSWORD}" | hashgardcli keys add ${KEY_NAME}

    # Init hashgard chain
    hashgard init --moniker hashgard --chain-id ${CHAIN_ID}
    hashgard add-genesis-account ${KEY_NAME} ${COIN_GENESIS}
    echo "${KEY_PASSWORD}" | hashgard gentx --name ${KEY_NAME} --amount ${COIN_DELEGATE}
    hashgard collect-gentxs
}

# ------------------------------------------------------------------------------
#
# Initial Hashgard block chain private multiple nodes
#
#   Folder .hashgard and .hashgardcli had been created by 'hashgard testnet'
#   and mount to /root/.hashgard and /root/.hashgardcli. So nothing to do.
#
# ------------------------------------------------------------------------------
function init_private_testnet() {
    echo "nothing to do !" > /dev/null
}

# ------------------------------------------------------------------------------
#
# Restart container
#
# ------------------------------------------------------------------------------
if [[ -e ${INITIALIZED_FLAG} ]]; then
    hashgard start
fi

# ------------------------------------------------------------------------------
#
# Node type:
#
#     FULL_NODE       - Run Hashgard block chain full node which can connect to
#                       Hashgard testnet and change to validator.
#     PRIVATE_SINGLE  - Run Hashgard block chain with single node.(Default)
#     PRIVATE_TESTNET - Run Hashgard block chain with multiple nodes created by
#                       command 'hashgard testnet'
#
# ------------------------------------------------------------------------------
if [[ -z ${NODE_TYPE} ]]; then
    NODE_TYPE="PRIVATE_SINGL"
fi
case "${NODE_TYPE}" in
    "FULL_NODE")
        init_full_node
        ;;
    "PRIVATE_SINGLE")
        init_private_single
        ;;
    "PRIVATE_TESTNET")
        init_private_testnet
        ;;
    *)
        echo ""
        echo "Error: "
        echo "    Environment NODE_TYPE must be one of FULL_NODE/PRIVATE_SINGLE/PRIVATE_TESTNET !"
        echo ""
        exit -1
esac

# Mark initial successfully
touch ${INITIALIZED_FLAG}

# Start service
hashgard start
