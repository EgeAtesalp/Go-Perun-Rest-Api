#!/bin/bash
# required on all ethereum nodes
# run startup-rasp.sh miner to setup additional miner nodes
cd /start/
git clone https://github.com/ethereum/go-ethereum.git /start/go-ethereum
cd go-ethereum
make geth
mv /start/go-ethereum/build/bin/geth /usr/local/bin
geth version
echo "now finish setup run startup-rasp.sh"