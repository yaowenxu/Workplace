package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"main/pubsub_grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:12349", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pubsub_grpc.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &pubsub_grpc.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &pubsub_grpc.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
