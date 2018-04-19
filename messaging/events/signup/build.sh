#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o main .
mkdir -p ./templates && cp -R ../../../email/templates/* ./templates
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main . 

docker build -t cha.signup.welcome:0.1 . 

# docker run --rm cha.signup.welcome:0.1