#!/bin/bash

cd /application/StreamMediaDevelopment/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api main.go

cd /application/StreamMediaDevelopment/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler main.go

cd /application/StreamMediaDevelopment/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver main.go

cd /application/StreamMediaDevelopment/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web main.go