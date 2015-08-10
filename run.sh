#!/usr/bin/env bash
docker rm -f restcahce-go
docker run --name restcahce-go -d -p 8080:8080 restcache-go