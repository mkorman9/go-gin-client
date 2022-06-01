package main

import (
	"context"
	"github.com/mkorman9/go-gin-protocol/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial(
		"localhost:9090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("error while connecting to service: %v\n", err)
	}

	client := protocol.NewClientServiceClient(connection)

	clientsStream, err := client.GetClients(context.Background(), &protocol.ClientRequest{})
	if err != nil {
		log.Fatalf("error while calling GetClients: %v\n", err)
	}

	for {
		client, err := clientsStream.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Fatalf("error while receiving client data: %v\n", err)
		}

		log.Printf("%v\n", client)
	}
}
