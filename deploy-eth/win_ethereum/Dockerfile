FROM multiarch/ubuntu-core:x86_64-focal
USER root

ARG CACHEBUST=0
RUN apt-get update
ARG CACHEBUST=0
RUN apt-get install software-properties-common -y
ARG CACHEBUST=0
RUN add-apt-repository -y ppa:ethereum/ethereum
ARG CACHEBUST=0
RUN apt-get install apt-utils openjdk-8-jdk net-tools openssh-server ethereum geth vim telnet -y
ARG CACHEBUST=0
RUN apt-get upgrade -y

RUN mkdir /start/
RUN mkdir /start/java
RUN mkdir /start/keystore
ARG CACHEBUST=0
COPY startup.sh  /start/startup.sh
ARG CACHEBUST=0
COPY genesis.json /start/genesis.json
COPY scripts/pwd/.passwd /start/.passwd
COPY scripts/keystore/* /start/keystore/
ARG CACHEBUST=0
ARG CACHEBUST=0
RUN chmod +755 /start/startup.sh

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu -p ubuntu
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN chown -R ubuntu /start
RUN chmod -R 755 /start

RUN ls -ali
RUN ls -ali /start

EXPOSE 22
EXPOSE 8545
EXPOSE 8546
EXPOSE 6060
EXPOSE 30303-30310
ENTRYPOINT /start/startup.sh