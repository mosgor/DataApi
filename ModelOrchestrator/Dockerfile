FROM golang:latest

WORKDIR /ModelOrchestrator

COPY . .

EXPOSE 8082 11411

ENV CONFIG_PATH=./config/prod.yaml

RUN go mod tidy

ENTRYPOINT go run cmd/main.go