#!/usr/bin/bash

# BACKEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/backend"
# FRONTEND_SERVER_PATH="/home/chaykovski/apps/iot-hub/frontend"

BACKEND_SERVER_PATH=~/code/iot-hub/backend
FRONTEND_SERVER_PATH=~/code/iot-hub/frontend


export PATH=$PATH:/usr/local/go/bin
export CGO_ENABLED=0

if [ ! $GOPATH ]; then
  export GOPATH=$HOME/go
fi

if [ ! $GOROOT ]; then
  export GOROOT=/usr/local/go
fi

go version

echo "scanning network"

hostname -I > nets.txt
NETS_FILE_PATH=$(pwd)/nets.txt

cd $BACKEND_SERVER_PATH
go run ./cmd/net/main.go \
  -subnetFile $NETS_FILE_PATH -port 19050 -path /data -output $BACKEND_SERVER_PATH/configs/.env.device 

echo "starting connector app"

cd $BACKEND_SERVER_PATH
go run ./cmd/connector/main.go

sleep 15

echo "scanning network"

hostname -I > nets.txt
NETS_FILE_PATH=$(pwd)/nets.txt

cd $BACKEND_SERVER_PATH
go run ./cmd/net/main.go \
  -subnetFile $NETS_FILE_PATH -port 19050 -path /data -output $BACKEND_SERVER_PATH/configs/.env.device 

echo "starting backend"

cd $BACKEND_SERVER_PATH
sudo chmod -R 777 ./postgres_data
docker-compose up --build -d

echo "starting frontend"

cd $FRONTEND_SERVER_PATH
bun --bun run dev 


