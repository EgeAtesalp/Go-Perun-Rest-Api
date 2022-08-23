FROM arm32v7/ubuntu
#armv7/armhf-ubuntu:latest
USER root

ARG CACHEBUST=0
RUN apt-get update
ARG CACHEBUST=0
RUN apt-get install software-properties-common -y
ARG CACHEBUST=0
RUN add-apt-repository -y ppa:ethereum/ethereum
ARG CACHEBUST=0
RUN apt-get install apt-utils openjdk-8-jdk net-tools openssh-server golang vim telnet git sudo -y
ARG CACHEBUST=0
RUN apt-get upgrade -y

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu -p ubuntu
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN mkdir /start
COPY ../scripts/startup.sh /start
RUN chown -R ubuntu /start
RUN chmod 755 /start/startup.sh

ARG CHACHEBUST=0
RUN ls -ali /start

USER ubuntu
WORKDIR /home/ubuntu

#TODO: git clone, build, deployment

EXPOSE 8080
ENTRYPOINT /start/startup.sh