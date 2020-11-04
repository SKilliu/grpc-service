FROM golang:1.13-stretch

WORKDIR $GOPATH/src/github.com/SKilliu/grpc-service/server

COPY . .

RUN go build -o server -v ./cmd/main.go