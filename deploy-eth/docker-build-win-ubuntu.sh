#!/bin/bash
VERSION="2021-11-17"
docker build -f win_ethereum/Dockerfile --tag eth-local:latest .
docker tag eth-local:latest eth-local:$VERSION
