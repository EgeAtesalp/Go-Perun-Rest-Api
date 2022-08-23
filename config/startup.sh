#!/bin/bash
SETUP=./installed

if test -f "$SETUP"; then
  echo "setup already done."
else
  echo "do setup."
  echo "done" >> $SETUP
  geth init --datadir /home/ubuntu/data /start/genesis.json
  cp /start/keystore/UTC--2020-06-29T14-29-21.261226215Z--1e66d21a5beadc6afeb9a74e9703eab207da27f2 /home/ubuntu/data/keystore/
  cp /start/.passwd /home/ubuntu/.passwd
fi

cd /home/ubuntu
pwd
ls -ali

if [ "$NODETYPE" = "bootnode" ]; then
  echo "bootnode"
  geth --datadir data --networkid 2000420101 --nat extip:192.168.0.21 --netrestrict 172.21.0.0/16 --miner.gasprice 10 --ws --rpc --wsaddr 0.0.0.0 --rpcaddr 0.0.0.0
elif [ "$NODETYPE" = "miner" ]; then
  echo "miner node starting" 
  geth --datadir data --networkid 2000420101 --netrestrict 172.21.0.0/16  --bootnodes enode://73501dbcbba8e77269ef523784e0b77003b53ef1a09abc5339a12c147b1fe5c8f4e8c05544a7d2fc96608ad1a4b86fb9c573ed93f554350353d418c7de00b11c@172.21.0.2:30303 --mine --miner.threads 4 --miner.gasprice 10 --unlock 1e66d21a5beadc6afeb9a74e9703eab207da27f2  --password .passwd
elif [ "$NODETYPE" = "" ]; then
  echo "no parameter given"
else
  echo "node is sth else"
fi

#while true; do
#   sleep 3 #600
#   echo "... still alive ..."
#done


