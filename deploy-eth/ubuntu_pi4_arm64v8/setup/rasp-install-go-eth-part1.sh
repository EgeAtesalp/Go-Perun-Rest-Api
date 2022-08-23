#!/bin/bash
# required on all ethereum nodes
# run startup-rasp.sh miner to setup additional miner nodes
apt-get update
apt-get upgrade -y
# see: https://kauri.io/running-an-ethereum-full-node-on-a-raspberrypi-4-m/9695fcca217f46feb355245275835fc0/a
apt-get install htop vim git sysstat make gcc -y
pwd
if [ "$2" == "" ]; then
  version="14.6"
else 
  version="$1"
fi

if [ "$1" == "" ]; then
  platform="linux-arm64"
else
  platform="$1"
fi

file="go1.$version.$platform.tar.gz"
echo "try to install from  https://golang.org/dl/$file"
wget https://golang.org/dl/go1.$version.$platform.tar.gz -O /start/$file
tar -C /usr/local -xvf /start/$file
# workaround, error gzip - but have been extracted to .1
# tar -C /usr/local -xf /start/go1.14.6.linux-arm64.tar.gz.1
chown root:root /usr/local/go
chmod 755 /usr/local/go
echo "vi /etc/profile : and export PATH=$PATH:/usr/local/go/bin"
echo "reboot the system and run rasp-install-go-eth-part2.sh"
