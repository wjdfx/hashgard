#!/usr/bin/env bash

COMMAND=$1
CONFIG_PATH=$2
CHAIN_ID=sif-3000
NODE_PREFIX="testnode-"
DOCKER_NETWORK="hashgard-network"

function usage() {
    echo "Usage:"
    echo "  ./testnet.sh command config-path"
    echo ""
    echo "Command:"
    echo "  run      Create new container for each node. "
    echo "  start    Start exist containers. "
    echo "  stop     Stop exist containers. "
    echo "  rm       Remove exist containers. "
    echo ""
    echo "config-path"
    echo "  Path to parameter '--output-dir' of command 'hashgard testnet' "
    echo ""
}

function node_count() {
    echo $(ls -1 ${CONFIG_PATH} | grep node -c)
}

function run() {
    # Create a network for connections between nodes
    if [[ "" == "$(docker network ls | grep ${DOCKER_NETWORK})" ]]; then
        docker network create ${DOCKER_NETWORK}
    fi

    # Create new container for each node
    for (( i=0; i<$(node_count); i++ )); do
        NODE_ROOT=${CONFIG_PATH}/node${i}
        if [[ ! -d ${NODE_ROOT} ]]; then
            echo "Node${i}'s config DOSE NOT exist !"
            exit -1
        fi

        NODE_NAME=${NODE_PREFIX}${i}
        echo -n "Create ${NODE_NAME} ... "
        docker run -d \
            --name ${NODE_NAME} \
            --net ${DOCKER_NETWORK} \
            -v ${NODE_ROOT}/.hashgard:/root/.hashgard \
            -v ${NODE_ROOT}/.hashgardcli:/root/.hashgardcli \
            hashgard:${CHAIN_ID} \
            > /dev/null
        echo "Done !"
    done
}

function start() {
    for (( i=0; i<$(node_count); i++ )); do
        NODE_NAME=${NODE_PREFIX}${i}
        echo -n "Start ${NODE_NAME} ... "
        docker start ${NODE_NAME} > /dev/null
        echo "Done !"
    done
}

function stop() {
    for (( i=0; i<$(node_count); i++ )); do
        NODE_NAME=${NODE_PREFIX}${i}
        echo -n "Stop ${NODE_NAME} ... "
        docker stop ${NODE_NAME} > /dev/null
        echo "Done !"
    done
}

function rm() {
    for (( i=0; i<$(node_count); i++ )); do
        NODE_NAME=${NODE_PREFIX}${i}
        echo -n "Remove ${NODE_NAME} ... "
        docker rm -f ${NODE_NAME} > /dev/null
        echo "Done !"
    done
}

if [[ $# != 2 ]]; then
    usage
    exit -1
fi

if [[ ! -d ${CONFIG_PATH}/node0 ]]; then
    echo -e  "\nError: Can not find config in ${CONFIG_PATH} !\n"
    usage
    exit -2
fi

case "${COMMAND}" in
    "run")
        run
        ;;
    "start")
        start
        ;;
    "stop")
        stop
        ;;
    "rm")
        rm
        ;;
    *)
        usage
        exit -1
esac
