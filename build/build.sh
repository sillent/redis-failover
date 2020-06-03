#!/bin/sh

export CGO_ENABLED=0
export GOOS=linux

#go get -u ./cmd/failover/
go build -ldflags '-extldflags "-static"' -a -v -o bin/linux/failover ./cmd/failover/
