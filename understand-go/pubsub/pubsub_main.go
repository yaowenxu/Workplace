package main

import (
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"main/pubsub_grpc"
	"net"
	"strings"
	"time"
)

func main() {
	//rpc_pubsub()
	grpc_pubsub_server()
}

func grpc_pubsub_server() {
	grpcServer := grpc.NewServer()
	pubsub_grpc.RegisterPubsubServiceServer(grpcServer, pubsub_grpc.NewPubsubService())
	lis, err := net.Listen("tcp", ":12349")

	if err != nil {
		log.Fatal(err)
	}
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

func rpc_pubsub() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(1)

	go func() {
		fmt.Println("golang topic:", <-golang)
	}()
	go func() {
		fmt.Println("docker topic:", <-docker)
	}()

	for true {
		time.Sleep(time.Second)
		fmt.Println()
	}
}

