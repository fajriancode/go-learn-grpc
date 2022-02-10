package main

import (
	"context"
	"log"

	"github.com/fajriancode/go-learn-grpc/proto/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect; %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server : %s", response.Body)

	broadcastMessage := chat.Message{
		Body: "This is broadcast message",
	}
	responseBroadcast, err := c.BroadcastMessage(context.Background(), &broadcastMessage)
	log.Printf("Response from server: %s", responseBroadcast.Body)
}
