FROM multiarch/ubuntu-core:x86_64-focal
USER root

ARG CACHEBUST=1
RUN apt-get update
RUN apt-get install software-properties-common -y
RUN add-apt-repository -y ppa:ethereum/ethereum
RUN apt-get install apt-utils openjdk-8-jdk net-tools openssh-server ethereum geth vim telnet -y
ARG CACHEBUST=1
RUN apt-get upgrade -y

RUN mkdir /start/
RUN mkdir /start/java
RUN mkdir /start/keystore
ARG CACHEBUST=0
COPY startup.sh  /start/startup.sh
ARG CACHEBUST=0
COPY genesis.json /start/genesis.json
COPY .passwd /start/.passwd
COPY ./keystore/UTC--2020-06-29T14-29-21.261226215Z--1e66d21a5beadc6afeb9a74e9703eab207da27f2 /start/keystore/
ARG CACHEBUST=0
RUN chmod -R 755 /start
ARG CACHEBUST=0
RUN chmod +755 /start/startup.sh
RUN ls -ali

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu -p ubuntu
USER ubuntu
WORKDIR /home/ubuntu

EXPOSE 22
EXPOSE 8545
EXPOSE 8546
EXPOSE 6060
EXPOSE 30303
USER root
ENTRYPOINT /start/startup.sh
