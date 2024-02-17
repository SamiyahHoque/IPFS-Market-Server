package main

import (
	"context"
	"log"
	"net"

	pb "example.com/fileserverproject"
	"google.golang.org/grpc"
)

// Adjust the import path
type server struct {
	pb.UnimplementedFileServer
}

// the (s *server) binds the functino to the server
func (s *server) SayHello (ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received %v", in.GetName())
	return &pb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to set up listening port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServer(s, &server{})
	log.Printf("Server Listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
}

