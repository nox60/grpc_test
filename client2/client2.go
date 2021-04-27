// Package main implements a client for Greeter service.
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "helloworld/helloworld"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := pb.NewGreeterClient(conn)
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//r, err := c.SayHello(ctx, &pb.HelloRequest{Username: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	clientSendMsg("eee", ctx, c)
	clientSendMsg("asdfasfd", ctx, c)
	clientSendMsg("24234", ctx, c)
	clientSendMsg("asdf2", ctx, c)
	clientSendMsg("qer", ctx, c)
}

func clientSendMsg(stringPars string, ctx context.Context, c pb.HelloServiceClient) (response string) {
	r, err := c.SayHello(ctx, &pb.HelloRequest{Username: stringPars})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return ""
}
