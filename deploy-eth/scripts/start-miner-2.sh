#!/bin/bash
geth --datadir data/data-miner-2 --networkid $NETWORKID --nat extip:$EXTIP_MINER --port $EXTPORT_MINER --bootnodes enode://$ENODE@$INTIP_BOOTNODE:30303 --mine --miner.threads 4 --miner.gasprice $GASPRICE --unlock $MINER_HASH  --password .passwd --log.json
