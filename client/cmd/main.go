package main

import (
	"context"
	"log"

	"github.com/SKilliu/grpc-service/server/proto/protogo"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := protogo.NewCoordinatesSaverClient(conn)

	resp, err := client.SaveCoordinates(context.Background(), &protogo.SaveRequest{
		Location:  "Kharkiv",
		Longitude: 10,
		Latitude:  11,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.GetOperationResult())
}
