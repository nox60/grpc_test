package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "message/message"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement message.GreeterServer.
type server struct {
	pb.UnimplementedMsgServiceServer
}

// SayHello implements message.GreeterServer
func (s *server) SendMsg(ctx context.Context, in *pb.MsgRequest) (*pb.MsgResponse, error) {
	// log.Printf("Received: %v", in.GetName())
	// println(in.GetUsername())
	return &pb.MsgResponse{Message: "Hello " + in.GetUsername()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	pb.RegisterMsgServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	println("done...")
}
