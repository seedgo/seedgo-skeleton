#!/bin/sh

version="$(date +%Y%m%d%H%M)"

# you should replace the {registry} with your container registry repository
docker build -t {registry}:${version} .
docker push {registry}:${version}


