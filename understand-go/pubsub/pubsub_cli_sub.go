package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"main/pubsub_grpc"
	"context"
)

func main() {
	conn, err := grpc.Dial("localhost:12349", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pubsub_grpc.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &pubsub_grpc.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}