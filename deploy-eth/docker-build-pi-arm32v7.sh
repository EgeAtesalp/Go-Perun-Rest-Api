#!/bin/bash
VERSION="v1.10.4_go1.13.8_2021-11-17"
docker build -f ubuntu_pi4_arm32v7/Dockerfile --tag ubuntu-arm32v7-eth:latest .
docker tag ubuntu-arm32v7-eth:latest ubuntu-arm32v7-eth:$VERSION
