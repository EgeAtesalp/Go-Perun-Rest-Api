FROM arm64v8/ubuntu

USER root

RUN apt-get update
RUN apt-get install software-properties-common apt-utils openjdk-8-jdk net-tools openssh-server nano golang vim telnet git sudo gpg-agent wget -y --no-install-recommends
RUN add-apt-repository -y ppa:ethereum/ethereum
RUN apt-get upgrade -y --no-install-recommends

ARG CACHEBUST=0
RUN mkdir /start
RUN mkdir /start/keystore
COPY scripts/startup-rasp.sh /start/
COPY scripts/start-boot.sh /start/
COPY scripts/start-miner-1.sh /start/
COPY scripts/start-miner-2.sh /start/
COPY scripts/startup.sh /start/
COPY scripts/rasp-install-go-eth-part1.sh /start/
COPY scripts/rasp-install-go-eth-part2.sh /start/
COPY scripts/genesis.json /start/
COPY scripts/keystore/* /start/keystore/
COPY scripts/pwd/.passwd /start/

RUN groupadd ubuntu -g 1000
RUN useradd -rm -d /home/ubuntu -s /bin/bash -g ubuntu -G root,sudo -u 1000 ubuntu -p ubuntu
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN chown -R ubuntu /start
RUN chmod 755 /start/*.sh
RUN ls -ali
RUN ls -ali /start

RUN sh /start/rasp-install-go-eth-part1.sh linux-arm64 1.15.6
RUN sh /start/rasp-install-go-eth-part2.sh

RUN mkdir -p /home/ubuntu/log
RUN mkdir -p /home/ubuntu/data
RUN chown ubuntu:ubuntu /home/ubuntu
RUN chown ubuntu:ubuntu /home/ubuntu/data
RUN chown ubuntu:ubuntu /home/ubuntu/log

EXPOSE 8545
EXPOSE 8546
EXPOSE 6060
EXPOSE 30303
EXPOSE 30304
EXPOSE 30305

USER ubuntu
WORKDIR /home/ubuntu

ENTRYPOINT /start/startup-rasp.sh

RUN ls -alih
RUN ls -alih data
RUN ls -alih log
