FROM multiarch/ubuntu-core:x86_64-focal
USER root

RUN mkdir /start/
COPY ../scripts/startup.sh /start/
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
RUN chown -R ubuntu /start

USER ubuntu
WORKDIR /home/ubuntu

#TODO: git clone, build, deployment

EXPOSE 8080
ENTRYPOINT /start/startup.sh
