#!/bin/sh

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

docker build -t messaging.connection:test .

docker run messaging.connection:test