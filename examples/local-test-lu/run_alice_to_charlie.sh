#!/bin/bash
echo "Running a channel from alice to charlie."

IP_Port=$1
REPEAT=$2

if [ "$IP_Port" = "" ]; then
    echo "Missing 1st parameter, add IP address as parameter for script call"
    echo "usage $0 x.x.x.x:port <repeat:=1000>"
    exit 1
fi

if [ "$REPEAT" = "" ]; then
    REPEAT=1000
fi

echo "Running a channel from alice to charlie."
# open both connections, together with json, the payment channel is generated automatically
curl -XPOST http://$IP_Port/v2/paymentchannel/open/charlie
curl -XPOST http://$IP_Port/v2/paymentchannel/connect/alice -d@channelToCharlie.json
# now exchange

for((j=$REPEAT; j>0; j--))
do
    echo "$j from $REPEAT"
    sleep 0.01   # waiting 1/1000 s
    curl -XPOST http://$IP_Port/v2/paymentchannel/send/alice -d@sendMoneyToCharlie.json
done

# close both
curl -XPOST http://$IP_Port/v2/paymentchannel/disconnect/alice -d@channelToCharlie.json
curl -XPOST http://$IP_Port/v2/paymentchannel/close/charlie   -d@channelToCharlie.json

exit 0