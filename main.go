// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io/ioutil"
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

	api.POST("/uploadfile", func(c *gin.Context) {
		//获取表单数据 参数为name值
		f, err := c.FormFile("f1")
		//错误处理
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		} else {
			file, _ := f.Open()
			fileConent, _ := ioutil.ReadAll(file)
			fmt.Println(string(fileConent))
			////将文件保存至本项目根目录中
			//c.SaveUploadedFile(f, f.Filename)
			//保存成功返回正确的Json数据
			//c.JSON(http.StatusOK, gin.H{
			//	"message": "OK",
			//	"messageBody": string(fileConent),
			//})
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Main website",
			})
		}

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

	api.POST("/send/:msg", func(c1 *gin.Context) {
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
