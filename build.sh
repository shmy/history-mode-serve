#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -v -a -ldflags "-s -w -extldflags -static" .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -v -a -ldflags "-s -w -extldflags -static" .

upx history-mode-serve
upx history-mode-serve.exe
