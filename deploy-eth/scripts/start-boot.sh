#!/bin/bash
geth --datadir data/data-boot --networkid $NETWORKID --nat extip:$EXTIP_BOOTNODE --miner.gasprice $GASPRICE --ws --rpc --ws.addr 0.0.0.0 --rpcaddr 0.0.0.0 --log.json
