#!/bin/bash
VERSION="v1.10.4_go1.13.8_2022-06-08"
docker build -f ubuntu_pi4_arm64v8/Dockerfile --tag ubuntu-arm64v8-eth:latest .
docker tag ubuntu-arm64v8-eth:latest ubuntu-arm64v8-eth:$VERSION
docker tag ubuntu-arm64v8-eth:latest git.tu-berlin.de:5000/ods-blockchain/go-perun-rest-api/rpi/ubuntu-arm64v8-eth:latest
docker tag ubuntu-arm64v8-eth:latest git.tu-berlin.de:5000/ods-blockchain/go-perun-rest-api/rpi/ubuntu-arm64v8-eth:$VERSION
