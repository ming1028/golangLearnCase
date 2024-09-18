#!/bin/sh
echo Building http-server:build
docker build -t http-server:build . -f Dockerfile.build
docker create --name builder http-server:build
docker cp builder:/go/src/github.com/golangLearnCase/docker/http-server ./http-server
docker rm -f builder
echo Building http-server:latest
docker build -t http-server:latest
rm -rf ./http-server