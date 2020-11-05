FROM golang:1.13-stretch

WORKDIR $GOPATH/src/github.com/SKilliu/grpc-service

COPY . .

RUN go build -o grpc-service -v ./cmd/main.go