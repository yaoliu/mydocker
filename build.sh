#!/usr/bin/env bash
GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/docker  cmd/docker/*
scp -r bin/docker root@192.168.11.1:/root/