# docker environments
- please use docker registry (e.g. docker login git.tu-berlin.de:5000)
- **Note:** buildung a complete new bootnode requires to:
    - reconfiguring all docker-compose files for new enode
    - checking all ip adresses and ports beeing used
    - resetup contracts by rest api, what does mean the contract hashes will change and needs also in the rest api to reconfigured

## ubuntu restapidemo [docker image](https://git.tu-berlin.de/ods-blockchain/go-perun-rest-api/container_registry/601)
- tested with ubuntu and windows docker environment
- please find [docker-compose.yaml](https://git.tu-berlin.de/ods-blockchain/go-perun-rest-api/-/tree/develop/deploy-api/win_ubuntu) as documentation for usage 

## local blockchain (1 boot with 3 miner nodes) [docker image](https://git.tu-berlin.de/ods-blockchain/go-perun-rest-api/container_registry/600)
- tested with ubuntu and windows docker environment
- please find [docker-compose.yaml](https://git.tu-berlin.de/ods-blockchain/go-perun-rest-api/-/tree/develop/deploy-api/win_ethereum) as documentation for usage 

## organize docker images
- delete images are not tagged
```
docker rmi -f $(docker images --filter "dangling=true" -q --no-trunc)
```
- delete images does not belong to at least one existing container
```
docker image prune -a
```

# set up environment on raspberrypi first time
- connect to the docker container registry in this git repository
- pull the docker container for the used hardware 
- tag the images that they fit the tags used in docker-compose files
- take care docker-compose is installed
    - example is for raspberry with PI4 8Gb - 64-bit OS and arm v8, for PI3 use arm32 tags instead
    - use tag ubuntu-arm64v8-eth:latest for ethereum nodes
    - use tag ubuntu-rpi-arm64v8-api:latest for api node 
- examples (TBD: do by ansible):
```
git clone https://git.tu-berlin.de/ods-blockchain/go-perun-rest-api.git

cd go-perun-rest-api
git switch feature/3-make-smart-contract-configurable

docker login git.tu-berlin.de:5000 -u <username> -p <token>
docker pull git.tu-berlin.de:5000/ods-blockchain/go-perun-rest-api/rpi/ubuntu-arm64v8-eth
docker tag git.tu-berlin.de:5000/ods-blockchain/go-perun-rest-api/rpi/ubuntu-arm64v8-eth:latest ubuntu-arm64v8-eth:latest

cd deploy-eth/ubuntu_pi4_arm64v8/

# as ubuntu user - for boot node
# modify ip addr (is done by /etc/profile.d/get-external-ip.sh)
# and enode
mkdir -p data-boot/data
mkdir -p data-boot/log
docker-compose up -f docker-compose-pi4-cluster.yml -d ubuntu-arm-64-boot

## miner docker-compose anpassen

# as ubuntu user - for miner-1-node
# modify ip addr in .env if mandatory
# EXTERNAL_BOOT_IP=192.168.1.171
mkdir -p data-miner-1/data
mkdir -p data-miner-1/log
docker-compose -f docker-compose-pi4-cluster.yml up -d ubuntu-arm-64-miner


# as ubuntu user - for miner-2-node
# modify ip addr in .env if mandatory
# EXTERNAL_BOOT_IP=192.168.1.171
mkdir -p data-miner-2/data
mkdir -p data-miner-2/log
docker-compose -f docker-compose-pi4-cluster.yml up -d ubuntu-arm-64-miner2

TBD: split/rotate logs
```

## [first] start up ethereum boot node
- update the docker-compose.yaml for the boot node EXTIP_BOOTNODE and INTIP_BOOTNODE to the public IP Adress
- start the ethereum boot node first time with time
```
docker-compose up -d ubuntu-arm-64-boot
```
- get the new enode for the boot node, check first line of log file (docker logs ...)

## [first] start up ethereum miner nodes
- update the docker-compose.yaml for the miner node ENODE, EXTIP_BOOTNODE and INTIP_BOOTNODE to the public IP Adress of boot node
- update the docker-compose.yaml for the miner node EXTPORT_MINER and EXTIP_MINER, by running 1 miner on a harware, just use default port 30303
    - if running a boot and a miner node take care that each port 30303 ... 30305 is used once on one hardware
- hint: start the miners each other step by step, check the logs for error/success
```
docker-compose up -d ubuntu-arm-64-miner
```
- get the new enode for the boot node, check first line of log file (docker logs ...)

## [first] start up api node
- update the docker-compose.yaml for the api node, change IP address in BLOCKCHAIN entry
- for a newly created blockchain it is required to deploy the smart contracts, once
    - this can be done with the alias alice, since that alias is configured for test and deploy
    - check the ports in docker-compose.yaml
    - call one time curl -XPOST http://`<ipaddress>:<port>`/v2/paymentchannel/open/alice 
    - check the log file (docker logs ...) for the new created hashes
    - update all <alias>.yaml in the root folder of go-perun-rest-api project folder to the new contract hashes
        - that modification can be made persistent, while commit the changes to a new docker image
        - note: tag the image to latest
    - inside docker the api is beeing started like: ./restapidemo --server --bc=ws://130.149.141.18:8545 --db=130.149.141.18 --pub=localhost , check that the docker-compose file sets the environment variables --bc=$BLOCKCHAIN --db=$DATABASE --pub=$EXTERNAL_IP correctly
```
docker-compose up -d ubuntu-arm-64-api
```

## stopping and restarting ethereum nodes
- use stop for stopping the boot and miner nodes
    - important: first the api and all miners, boot node at the last one
```
docker-compose stop ubuntu-arm-64-boot 
docker-compose stop ubuntu-arm-64-api
```

- use start for starting again the boot and miner nodes
    - important: first the boot node, than the miners step by step, at last the api
```
docker-compose start ubuntu-arm-64-boot
docker-compose start ubuntu-arm-64-api
```



# reminder for setup - from scratch

## geth (ethereum)
follow these tutorials with amd or x64 
- https://geth.ethereum.org/docs/install-and-build/installing-geth
- https://geth.ethereum.org/docs/interface/private-network

follow this for arm
- https://kauri.io/running-an-ethereum-full-node-on-a-raspberrypi-4-m/9695fcca217f46feb355245275835fc0/a

### swagger api
- docker pull swaggerapi/swagger-editor
- docker run -d -p 80:8080 swaggerapi/swagger-editor

> use and update projects swagger.yaml
> update swagger.json to enable/update API reference on api server


### config 
see config-docker-env.zip

create startup script based on startup.sh
- <del>192.168.1.199 will offer the boot node, 2 miners and rest api </del>
- <del>192.168.1.108 will offer just a miner node</dev>
- see https://git.tu-berlin.de/ods-blockchain/infrastructure-pi-cluster for infrastructure

copy the setup.zip and secrets.zip to the folder /start/
- run rasp-install-go-eth-part1.sh <os> <version>
    - note different raspberry might need a different ethereum version
    - Raspberry Pi 4 Model B Rev 1.2 with Ubuntu 18.04.5 LTS require linux-armv6l
    - Raspberry Pi 4 Model B Rev 1.4 with Ubuntu 20.04.1 LTS require linux-arm64
- modify vi /etc/profile : and export PATH=$PATH:/usr/local/go/bin
- reboot system
- run rasp-install-eth-part2.sh
- run startup-rasp.sh with option bootnade | miner | nothing just to init ethereum

run this script like
- as bootnode:  nohup bash startup-rasp.sh bootnode  2>&1 > logfile-bootnode.log &
- as miner:     nohup bash startup-rasp.sh miner  2>&1 > logfile-miner.log &
- as 2nd miner: nohup bash startup-rasp.sh miner2  2>&1 > logfile-miner-2.log &

to execute in the background 

## go-lang rest server
run server nohup ./restapidemo server 2>&1 > logfile.log &

### trouble shooting
- go server requires to find the package mux you can install with "go get github.com/gorilla/mux"
- usually missing modules are installed by using go build
- calling <del> http://130.149.141.17:8080/api </del> in ODS VPN (pi cloud) with a browser supports a full description and test environment for the rest api, the api calls will work e.g. with curl connecting <del> 192.168.1.199 </del>

### open issues TODO:
- finalize all returned data into json and put the important (log) data as result
- implement multi peer session mgmt
- prepare docker images for api and block chain

### deployed contracts and users
- ETH RPC URL: ws://192.168.1.199:8546
- ETHAssetHolder: 0x306D8b1CaFdD74b5E8c4C35C16B6Bc0Eab6401CA
- Adjudicator: 0xe0e6A97883959DA07250f878d9c7064bc63e1A9F
    - note: after setting up an new boot node, the api needs to deploy the contracts again
    - the user must have in the configuration testandeploy
    - the new transaction hashes have to bee replaced in all configuration files <alias>.yaml 
- alice   0x1e66D21a5BEADC6AfeB9a74E9703EaB207Da27f2   127.0.0.1:5751
- bob     0x30494CB6B8e89AbE164Ddf33FB30f69150DE67cA   127.0.0.1:5750
 
### curl examples (note: ip addresses have beeing changed)
- curl -XPOST http://130.149.141.17:8080/v2/paymentchannel/open/bob
- curl -XPOST http://130.149.141.17:8080/v2/paymentchannel/open/alice -d@channelToBob1000.json
- curl -XPOST http://130.149.141.17:8080/v2/paymentchannel/send/alice -d@sendMoneyToBob.json
- curl -XGET http://130.149.141.17:8080/v2/paymentchannel/info/alice
- curl -XPOST http://130.149.141.17:8080/v2/paymentchannel/close/alice -d@channelToBob1000.json

## channelToBob1000.json
```
{
  "target": "bob",
  "ownBalance": 1000,
  "theirsBalance": 0
}
```

## sendMoneyToBob.json
```
{
  "target": "bob",
  "balance": 10
}
```

# logging boot and minernodes

## load logs into postgres
```
CREATE TABLE "bootnode_log" ( "data" JSONB NULL );
\copy bootnode_log from '/mnt/f/TU Berlin/Data/log/bootnode.log';
CREATE INDEX datagin ON bootnode_log USING gin (data);
```

## request log data example
```
SELECT b.data->>'t' AS timestamps, b.data->>'msg' AS messages, b.data->>'lvl' AS log_level, 
b.data->>'txs' AS txs, b.data->>'hash' AS hash, b.data->>'mgas' AS mgas, b.data->>'dirty' AS dirty, 
b.data->>'mgasps' AS mgasps, b.data->>'blocks' AS blocks, b.data->>'number' AS NUMBER, b.data->>'elapsed' AS elapsed
from bootnode_log b 
ORDER BY timestamps ASC;
```

# logging and measurement network packets
```
tcpdump -i any -w pi4_1g_1.tcpdump
tcpdump -i any -G 3600 -w "%Y%m%d-%H%M_pi4-1g-1.tcpdump"
sudo tcpdump -i any -G 3600 -w $EXTERNAL_IP"_%Y%m%d-%H%M.tcpdump" &

tshark -r pi4_1g_1.tcpdump -T fields -e frame.number -e frame.time -e ip.src -e tcp.srcport -e ip.dst -e tcp.dstport -e frame.len -e frame.protocols -E header=y -E quote=n -E occurrence=f

tshark -r pi4_1g_1.tcpdump -T fields -e frame.number -e frame.time -e ip.src -e tcp.srcport -e ip.dst -e tcp.dstport -e frame.len -e frame.protocols -E header=y -E quote=n -E occurrence=f > pi4_1g_1.tcpdump.csv

tshark -r pi4_1g_1.tcpdump -T json > pi4_1g_1.tcpdump.json
```

# logging systems properties
```
$ virtualenv measure
$ source measure/bin/activate
$ cd measurement
# in case this library cannot installed because PEP 517
# check that by using ubuntu: libhdf5-dev have been installed
# sudo apt install libhdf5-dev
$ pip install -r requirements.txt

$ python /home/ubuntu/go-perun-rest-api/measurement/meas.py -f <file> -t json 
$ python meas.py -t json -f /mnt/c/Daten/test-log.json --duration 10 --interval 1000
$ python meas.py -t json -f /home/ubuntu/go-perun-rest-api/deploy-eth/ubuntu_pi4_arm64v8/data-boot/log/system-log.json --duration 36000 --interval 1000 2>&1 > /dev/null &
```

# copy all files from remote to local host, clean shutdown ...
```
pkill --signal=TERM geth
docker exec -ti c2c85e7f2371 pkill --signal=TERM startup
tail -f go-perun-rest-api/deploy-eth/ubuntu_pi4_arm64v8/log/miner-1.log
scp ubuntu@192.168.1.175:/home/ubuntu/go-perun-rest-api/deploy-eth/ubuntu_pi4_arm64v8/log/* log/175/
```

# solving git clone issue 'server certificate verification failed. CAfile: none CRLfile: none'
```
sudo apt-get install --reinstall ca-certificates
sudo mkdir /usr/local/share/ca-certificates/cacert.org
sudo wget -P /usr/local/share/ca-certificates/cacert.org http://www.cacert.org/certs/root.crt http://www.cacert.org/certs/class3.crt
sudo update-ca-certificates
git config --global http.sslCAinfo /etc/ssl/certs/ca-certificates.crt
sudo shutdown -r now
```
