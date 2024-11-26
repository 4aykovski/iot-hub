!#/bin/bash

BACKEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/backend"

cd $BACKEND_SERVER_PATH
go run ./cmd/app/main.go &
