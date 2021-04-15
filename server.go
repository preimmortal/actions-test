package main

import (
	"context"
	"log"
	"net"

	pb "github.com/preimmortal/actions-test/proto"
	"google.golang.org/grpc"
)

// Server is used to implement actions-test.ActionsServer.
type Server struct {
	pb.ActionsServer
}

// ActionEcho implements actions-test.ActionsServer
func (s *Server) ActionEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Printf("Received: %v", in.GetInput())
	return &pb.EchoReply{Message: in.GetInput()}, nil
}

func serve(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting Server")
	log.Printf("Listening on Port %v", port)
	s := grpc.NewServer()

	log.Printf("Logging in to Postgres")

	pb.RegisterActionsServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
