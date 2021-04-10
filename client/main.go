package main

import (
	"context"
	"log"
	"time"

	pb "github.com/preimmortal/actions-test/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
	input   = "preimmortal"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewActionsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ActionEcho(ctx, &pb.EchoRequest{Input: input})
	if err != nil {
		log.Fatalf("could not contact: %v", err)
	}
	log.Printf("Return: %s", r.GetMessage())
}
