#!/bin/bash
echo "Running a channel from alice to dori."

IP_Port=$1
REPEAT=$2

if [ "$IP_Port" = "" ]; then
    echo "Missing 1st parameter, add IP address as parameter for script call"
    echo "usage $0 x.x.x.x:port <repeat:=500>"
    exit 1
fi

if [ "$REPEAT" = "" ]; then
    REPEAT=500
fi

# open both connections, together with json, the payment channel is generated automatically
curl -XPOST http://$IP_Port/v2/paymentchannel/open/dori
curl -XPOST http://$IP_Port/v2/paymentchannel/connect/alice -d@channelToDori.json
# now exchange

for((k=$REPEAT; k>0; k--))
do
    echo "$k from $REPEAT"
    sleep 0.001   # waiting 1/1000 s 
    curl -XPOST http://$IP_Port/v2/paymentchannel/send/alice -d@sendMoneyToDori.json
done

# close both
curl -XPOST http://$IP_Port/v2/paymentchannel/disconnect/alice -d@channelToDori.json
curl -XPOST http://$IP_Port/v2/paymentchannel/close/dori   -d@channelToDori.json

exit 0