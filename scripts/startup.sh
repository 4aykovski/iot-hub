!#/bin/bash

BACKEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/backend"
FRONTEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/frontend"

cd $BACKEND_SERVER_PATH
go run ./cmd/connector/main.go

cd $FRONTEND_SERVER_PATH
bun --bun run dev &

cd $BACKEND_SERVER_PATH
go run ./cmd/iot/main.go


