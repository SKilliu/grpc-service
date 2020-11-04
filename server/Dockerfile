FROM golang:1.13-stretch

WORKDIR $GOPATH/src/grpc-service/grpc-server/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o grpc-server -v ./cmd/main.go