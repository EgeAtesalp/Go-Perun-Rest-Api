FROM arm32v7/ubuntu

USER root

ARG CACHEBUST=0
RUN apt-get update
ARG CACHEBUST=0
RUN apt-get install software-properties-common -y --no-install-recommends
ARG CACHEBUST=0
RUN add-apt-repository -y ppa:ethereum/ethereum
ARG CACHEBUST=0
RUN apt-get install apt-utils openjdk-8-jdk net-tools openssh-server golang vim telnet git sudo -y --no-install-recommends
ARG CACHEBUST=0
RUN apt-get upgrade -y

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

ARG CACHEBUST=0
RUN chown -R ubuntu /start

ARG CACHEBUST=0
RUN chmod 755 /start/*.sh

ARG CACHEBUST=0
RUN ls -ali

ARG CACHEBUST=0
RUN ls -ali /start

RUN sh /start/rasp-install-go-eth-part1.sh linux-armv6l 1.15.6
RUN sh /start/rasp-install-go-eth-part2.sh

RUN mkdir -p /home/ubuntu/log
RUN mkdir -p /home/ubuntu/data

EXPOSE 8545
EXPOSE 8546
EXPOSE 6060
EXPOSE 30303
EXPOSE 30304
EXPOSE 30305

USER ubuntu
WORKDIR /home/ubuntu

ENTRYPOINT /start/startup-rasp.sh