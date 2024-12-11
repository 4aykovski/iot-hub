FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod ./ 
COPY go.sum ./ 

RUN go mod download

COPY . .

RUN go build -o ./app cmd/iot/main.go

FROM alpine 

WORKDIR /app

COPY --from=builder /app/app ./ 
COPY --from=builder /app/migrations ./migrations 
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/.env ./.env

CMD ["./app"]
