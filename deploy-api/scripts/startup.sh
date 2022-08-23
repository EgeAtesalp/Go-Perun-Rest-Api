#!/bin/bash
sudo chown -R ubuntu /home/ubuntu/log
cd /home/ubuntu/go-perun-rest-api/
./restapidemo --server --bc=$BLOCKCHAIN --db=$DATABASE --pub=$EXTERNAL_IP 2>&1 | tee /home/ubuntu/log/rest-api.log > /dev/null &

# ideling to keep docker running
while true; do
   sleep 30
#   echo "... still alive ..."
done