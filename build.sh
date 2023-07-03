#!/bin/sh

version="$(date +%Y%m%d%H%M)"
docker build -t registry.cn-hangzhou.aliyuncs.com/haydenzhou/mdtube-api:${version} .
docker push registry.cn-hangzhou.aliyuncs.com/haydenzhou/mdtube-api:${version}


