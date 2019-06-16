@echo off

set GOOS=windows
set GOARCH=amd64
go build

set GOOS=linux
go build -o numberToWord-linux

set GOOS=darwin
go build -o numberToWord-mac