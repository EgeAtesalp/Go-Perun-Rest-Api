#!/bin/bash

# run this script like 
# as bootnode:  nohup bash startup-rasp.sh bootnode  2>&1 > logfile-bootnode.log
# as miner:     nohup bash startup-rasp.sh miner  2>&1 > logfile-miner.log
# to execute in the background

if pwd = "/start"; then
  echo "correct start path /start"
else
  echo "start script from path /start"
  exit 42
fi

NODETYPE=$1
SETUP=$NODETYPE/installed
mkdir -p /start/$NODETYPE
mkdir -p /home/ubuntu

if test -f "$SETUP"; then
  echo "setup already done."
else
  # requires that the keystore, .passwd and genesis.json have been copied to /start
  echo "do setup."
  echo "done" >> $SETUP
  if [ "$NODETYPE" = "bootnode" ]; then
    geth init --datadir /home/ubuntu/data /start/genesis.json
    cp /start/keystore/UTC--2020-06-29T14-29-21.261226215Z--1e66d21a5beadc6afeb9a74e9703eab207da27f2 /home/ubuntu/data/keystore/
  elif [ "$NODETYPE" = "miner" ]; then
    geth init --datadir /home/ubuntu/data-miner /start/genesis.json
    cp /start/keystore/UTC--2020-06-29T14-29-21.261226215Z--1e66d21a5beadc6afeb9a74e9703eab207da27f2 /home/ubuntu/data-miner/keystore/
    cp /start/.passwd /home/ubuntu/.passwd
  elif [ "$NODETYPE" = "miner2" ]; then
    geth init --datadir /home/ubuntu/data-miner-2 /start/genesis.json
    cp /start/keystore/UTC--2020-06-29T14-29-21.261226215Z--1e66d21a5beadc6afeb9a74e9703eab207da27f2 /home/ubuntu/data-miner-2/keystore/
    cp /start/.passwd /home/ubuntu/.passwd
  fi
fi

cd /home/ubuntu
pwd
ls -ali

if [ "$NODETYPE" = "bootnode" ]; then
  echo "bootnode"
  geth --datadir data --networkid 2000420101 --nat extip:192.168.1.199 --netrestrict 192.168.1.0/24 --miner.gasprice 10 --ws --rpc --wsaddr 0.0.0.0 --rpcaddr 0.0.0.0
elif [ "$NODETYPE" = "miner" ]; then
  echo "miner node starting"
  # you need to run on the bootnode "geth attach data/geth.ipc --exec admin.nodeInfo.enr" while geth is running as bootnode @217 to get enode infos
  geth --datadir data-miner --networkid 2000420101 --netrestrict 192.168.1.0/24 --port 30304  --bootnodes enode://7004a26dc02eda9eaa73073fdec08916f0986f1c1ad7889619639113261d3e681f22ea846befba351748f4e60992321e04ce65a0b53d97b945ff4ed0f0f337fb@192.168.1.199:30303 --mine --miner.threads 4 --miner.gasprice 10 --unlock 1e66d21a5beadc6afeb9a74e9703eab207da27f2  --password .passwd
elif [ "$NODETYPE" = "miner2" ]; then
  echo "miner 2 node starting"
  # you need to run on the bootnode "geth attach data/geth.ipc --exec admin.nodeInfo.enr" while geth is running as bootnode @217 to get enode infos
  geth --datadir data-miner-2 --networkid 2000420101 --netrestrict 192.168.1.0/24 --port 30305  --bootnodes enode://7004a26dc02eda9eaa73073fdec08916f0986f1c1ad7889619639113261d3e681f22ea846befba351748f4e60992321e04ce65a0b53d97b945ff4ed0f0f337fb@192.168.1.199:30303 --mine --miner.threads 4 --miner.gasprice 10 --unlock 1e66d21a5beadc6afeb9a74e9703eab207da27f2  --password .passwd
elif [ "$NODETYPE" = "" ]; then
  echo "no parameter given"
else
  echo "node is sth else"
fi