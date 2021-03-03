FROM golang:latest

WORKDIR /go/src/list-exporter
COPY . .

ENTRYPOINT ["go","run","./main.go","./collector.go"]
