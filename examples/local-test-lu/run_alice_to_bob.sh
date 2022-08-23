#!/bin/bash
echo "Running a channel from alice to bob."

IP_Port=$1
REPEAT=$2

if [ "$IP_Port" = "" ]; then
    echo "Missing 1st parameter, add IP address as parameter for script call"
    echo "usage $0 x.x.x.x:port <repeat:=10000>"
    exit 1
fi

if [ "$REPEAT" = "" ]; then
    REPEAT=10000
fi

echo "validate / deploy smart contract"
curl -XPOST http://$IP_Port/v2/paymentchannel/validateordeploy/alice

echo "Running a channel from alice to bob."
# open both connections, together with json, the payment channel is generated automatically
curl -XPOST http://$IP_Port/v2/paymentchannel/open/bob
curl -XPOST http://$IP_Port/v2/paymentchannel/open/alice -d@channelToBob.json
# now exchange

for((i=$REPEAT; i>0; i--))
do
    echo "$i from $REPEAT"
    sleep 0.01 # waiting 1/100 s
    curl -XPOST http://$IP_Port/v2/paymentchannel/send/alice -d@sendMoneyToBob.json
done

# close both
curl -XPOST http://$IP_Port/v2/paymentchannel/close/alice -d@channelToBob.json
curl -XPOST http://$IP_Port/v2/paymentchannel/close/bob   -d@channelToBob.json

exit 0