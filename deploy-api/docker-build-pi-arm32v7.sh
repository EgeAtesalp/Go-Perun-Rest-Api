#!/bin/bash
VERSION="v0.6_2021-xx-xx"
docker build -f ubuntu_pi4_arm32v7/Dockerfile --tag ubuntu-rpi-arm32v7-api:latest .
docker ubuntu-rpi-arm32v7-api:latest ubuntu-rpi-arm32v7-api:$VERSION
