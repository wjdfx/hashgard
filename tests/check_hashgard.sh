#!/bin/sh

# Note: Bucky, I know you want to kill bash tests.
# Please show me how to do an alternative to this.
# I would rather get code running before I leave than
# fight trying to invent some new test harness that
# no one else will understand.
#
# Thus, I leave this as an exercise to the reader to
# port into a non-bash version. And I don't do it proper...
# just automate my manual tests

# WARNING!!!
rm -rf ~/.hashgard ~/.hashgardcli
cd $GOPATH/src/github.com/hashagard/hashgard
make build

# init stuff
SEED=`./build/hashgard init | tail -1`
PASS='12345678'
(echo $PASS; echo $SEED) | ./build/hashgardcli keys add demo --recover
ADDR=`./build/hashgardcli keys show demo | cut -f3`
echo "Recovered seed for demo:" $ADDR

# start up server
./build/hashgard start > ~/.hashgard/hashgard.log 2>&1 &
sleep 5
PID_SERVER=$!

# query original state
TO='ABCAFE00DEADBEEF00CAFE00DEADBEEF00CAFE00'
echo; echo "My account:" $ADDR
./build/hashgardcli account $ADDR
echo; echo "Empty account:" $TO
./build/hashgardcli account $TO

# send some money
TX=`echo $PASS | ./build/hashgardcli send --to=$TO --amount=1000mycoin --from=demo --seq=0`
echo; echo "SendTx"; echo $TX
HASH=`echo $TX | cut -d' ' -f6`
echo "tx hash:" $HASH

# let some blocks come up....
./build/hashgardcli status | jq .latest_block_height
sleep 2
./build/hashgardcli status | jq .latest_block_height

# balances change
echo; echo "My account went down"
./build/hashgardcli account $ADDR
echo; echo "Empty account got some cash"
./build/hashgardcli account $TO

# query original tx
echo; echo "View tx"
./build/hashgardcli tx $HASH

# wait a bit then dump out some blockchain state
sleep 10
./build/hashgardcli status --trace
./build/hashgardcli validators --trace

# shutdown, but add a sleep if you want to manually run some cli scripts
# against this server before it goes away
# sleep 120
kill $PID_SERVER

