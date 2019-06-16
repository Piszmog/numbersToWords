#!/usr/bin/env bash

env GOOS=darwin GOARCH=amd64 go build -o numberToWord-mac
env GOOS=linux GOARCH=amd64 go build -o numberToWord-linux
env GOOS=windows GOARCH=amd64 go build