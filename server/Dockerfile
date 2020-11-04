FROM golang:1.13-stretch

WORKDIR $GOPATH/src/github.com/SKilliu/grpc-service

COPY . .

RUN go build -o ./server -v ./cmd/main.go