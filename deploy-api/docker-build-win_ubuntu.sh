#!/bin/bash#!/bin/bash
VERSION="v0.6_20211112"
docker build -f ubuntu_win_ubuntu/Dockerfile --tag ubuntu-win-x64-api:latest .
docker ubuntu-win-x64-api:latest ubuntu-win-x64-api:$VERSION