#!/usr/bin/bash

BACKEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/backend"
FRONTEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/frontend"

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export CGO_ENABLED=0

echo "starting net app"

go version

cd $BACKEND_SERVER_PATH
go run ./cmd/net/main.go \
  -subnet 0.0.0 -port 19050 -path /data -output $BACKEND_SERVER_PATH/configs/.env.device 

echo "starting connector app"

cd $BACKEND_SERVER_PATH
go run ./cmd/connector/main.go

echo "starting frontend"

cd $FRONTEND_SERVER_PATH
bun --bun run dev &

echo "starting backend"

cd $BACKEND_SERVER_PATH
go run ./cmd/iot/main.go


