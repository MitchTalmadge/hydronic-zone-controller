#!/bin/bash

echo "amd64..."
env GOOS=linux GOARCH=amd64 go build -o ./out/amd64/zonecontrol ./src/*.go
echo "arm..."
env GOOS=linux GOARCH=arm go build -o ./out/arm/zonecontrol ./src/*.go
echo "arm64..."
env GOOS=linux GOARCH=arm64 go build -o ./out/arm64/zonecontrol ./src/*.go