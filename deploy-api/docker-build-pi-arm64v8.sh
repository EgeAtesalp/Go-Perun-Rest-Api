#!/bin/bash
VERSION="v0.6_2021-xx-xx"
docker build -f ubuntu_pi4_arm64v8/Dockerfile --tag ubuntu-rpi-arm64v8-api:latest .
docker ubuntu-rpi-arm64v8-api:latest ubuntu-rpi-arm64v8-api:$VERSION
