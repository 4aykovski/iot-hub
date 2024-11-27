!#/bin/bash

BACKEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/backend"
FRONTEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/frontend"

export CGO_ENABLED=0

cd $BACKEND_SERVER_PATH
go run ./cmd/net/main.go \
  -subnet 0.0.0 -port 19050 -path /data -output $BACKEND_SERVER_PATH/configs/.env.device 

cd $BACKEND_SERVER_PATH
go run ./cmd/connector/main.go

cd $FRONTEND_SERVER_PATH
bun --bun run dev &

cd $BACKEND_SERVER_PATH
go run ./cmd/iot/main.go


