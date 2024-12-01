FROM golang:latest

WORKDIR /SourceManager

COPY . .

EXPOSE 8081

ENV CONFIG_PATH=./config/prod.yaml

RUN go mod tidy

ENTRYPOINT go run cmd/main.go