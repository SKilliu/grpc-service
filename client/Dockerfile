FROM golang:1.13-stretch

WORKDIR $GOPATH/src/grpc-service/client

COPY . .

RUN go build -o client -v ./cmd/main.go