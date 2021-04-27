package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "helloworld/helloworld"
	"io"
	"net"
)

//https://www.grpc.io/docs/languages/go/quickstart/

type HelloService struct {
}

func (hs *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("你好，%s", req.Username)}, nil
}

// ++++++++++++++++++++++++++++++++
func (hs *HelloService) Chat(conn pb.HelloService_ChatServer) error {
	for {
		stream, err := conn.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("receive from client:", stream.Stream)

		conn.Send(&pb.ServerStream{
			Stream: newBytes(1, 2, 3, 4, 5, 67),
		})
	}
	return nil
}

// ++++++++++++++++++++
func main() {
	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloService{})
	go func() {
		s.Serve(lis)
	}()
	fmt.Println(s.GetServiceInfo())
	select {}
}

func newBytes(a ...byte) []byte {
	return a
}
