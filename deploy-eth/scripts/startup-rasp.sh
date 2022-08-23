#!/bin/bash

# run this script like
# as bootnode:  nohup bash startup-rasp.sh bootnode  2>&1 > logfile-bootnode.log
# as miner:     nohup bash startup-rasp.sh miner  2>&1 > logfile-miner.log
# to execute in the background

cd /start

if pwd = "/start"; then
  echo "correct start path /start"
else
  echo "start script from path /start"
  exit 42
fi

SETUP=$NODETYPE/installed
mkdir -p /home/ubuntu/data/$NODETYPE

echo "change owner to ubuntu for data and log folder"
sudo chown -R ubuntu:ubuntu /home/ubuntu/data
sudo chown -R ubuntu:ubuntu /home/ubuntu/log

if test -f "$SETUP"; then
  echo "setup already done."
else
  # requires that the keystore, .passwd and genesis.json have been copied to /start
  echo "do setup."
  echo "done" >> /home/ubuntu/data/$SETUP
  if [ "$NODETYPE" = "bootnode" ]; then
    geth init --datadir /home/ubuntu/data/data-boot /start/genesis.json
    cp /start/keystore/* /home/ubuntu/data/data-boot/keystore/
  elif [ "$NODETYPE" = "miner" ]; then
    geth init --datadir /home/ubuntu/data/data-miner-1 /start/genesis.json
    cp /start/keystore/*/home/ubuntu/data/data-miner-1/keystore/
    cp /start/.passwd /home/ubuntu/.passwd
  elif [ "$NODETYPE" = "miner2" ]; then
    geth init --datadir /home/ubuntu/data/data-miner-2 /start/genesis.json
    cp /start/keystore/* /home/ubuntu/data/data-miner-2/keystore/
    cp /start/.passwd /home/ubuntu/.passwd
  fi
fi

cd /home/ubuntu
pwd
ls -ali

function cleanup()
{
  if [ "$BOOT" != "" ]; then
    echo "try to stop process $BOOT"
    kill -s STOP ${BOOT}
  fi
  if [ "$MINPROC1" != "" ]; then
    echo "try to stop process $MINPROC1"
    kill -s STOP ${MINPROC1}
  fi
  if [ "$MINPROC2" != "" ]; then
    echo "try to stop process $MINPROC2"
    kill -s STOP ${MINPROC2}
  fi
  exit 1
}

trap cleanup EXIT HUP INT QUIT PIPE TERM

whoami
echo " tries " 

if [ "$NODETYPE" = "bootnode" ]; then
  echo "bootnode"
  nohup bash /start/start-boot.sh 2>&1 | tee /home/ubuntu/log/bootnode.log > /dev/null &
  BOOT=$!
  echo "process started with pid $BASHPID"
  rm /home/ubuntu/log/boot.pid
  echo $BASHPID >> /home/ubuntu/log/boot.pid
elif [ "$NODETYPE" = "miner" ]; then
  echo "miner node starting"
  # you need to run on the bootnode "geth attach data/geth.ipc --exec admin.nodeInfo.enr" while geth is running as bootnode @217 to get enode infos
  nohup bash /start/start-miner-1.sh 2>&1 | tee /home/ubuntu/log/miner-1.log > /dev/null &
  MINPROC1=$!
  echo "process started with pid $BASHPID"
  rm /home/ubuntu/log/miner1.pid
  echo $BASHPID >> /home/ubuntu/log/miner1.pid
elif [ "$NODETYPE" = "miner2" ]; then
  echo "miner 2 node starting"
  # you need to run on the bootnode "geth attach data/geth.ipc --exec admin.nodeInfo.enr" while geth is running as bootnode @217 to get enode infos
  nohup bash /start/start-miner-2.sh 2>&1 | tee /home/ubuntu/log/miner-2.log > /dev/null &
  MINPROC2=$!
  echo "process started with pid $BASHPID"
  rm /home/ubuntu/log/miner2.pid
  echo $BASHPID >> /home/ubuntu/log/miner2.pid
elif [ "$NODETYPE" = "" ]; then
  echo "no parameter given"
else
  echo "node is sth else"
fi

# TODO: implement function for start and stop (pgrep -P pid)
echo "start idleing"
while true; do
  #echo "... still alive ..."
  sleep 60
done
