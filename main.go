// Package main implements a client for Greeter service.
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	pb "message/message"
	"net/http"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("static/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.Static("/htmls", "./htmls")

	// router.StaticFS("/data", http.Dir("./data"))

	api := r.Group("/api")

	api.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	api.GET("/send/:word", func(c1 *gin.Context) {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewMsgServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		word := c1.Param("word")
		resp := clientSendMsg(word, ctx, c)

		if word == "1" {

		}

		c1.JSON(200, gin.H{
			"result": "1",
			"msg":    resp,
		})
	})

	r.Run()
}

func clientSendMsg(stringPars string, ctx context.Context, c pb.MsgServiceClient) (response string) {
	r, err := c.SendMsg(ctx, &pb.MsgRequest{Username: stringPars})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	return r.GetMessage()
}
