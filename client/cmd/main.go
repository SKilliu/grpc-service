package main

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"github.com/SKilliu/grpc-service/proto/protogo"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
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
		log.Fatal(errors.Wrap(err, "failed to send request"))
	}

	log.Println(resp.GetOperationResult())
}
