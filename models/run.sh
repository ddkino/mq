#!/bin/sh

rm -f *.go
# protoc --go_out=plugins=protoc-gen-go:. *.proto
protoc --go_out=plugins=protoc-gen-go:./error error.proto
protoc --go_out=plugins=protoc-gen-go:./user user.proto