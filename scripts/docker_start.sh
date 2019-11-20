#!/bin/bash

docker build -t eferhatg/cas-server ./server
docker run -p 5000:5000 -t eferhatg/cas-server
docker build -t eferhatg/cas-cache ./cacheserver
docker run -p 18001:18001 -t eferhatg/cas-cache
docker build -t eferhatg/cas-client ./client
docker run -t eferhatg/cas-client
