package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"myProtobuf/serversiderpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := serversiderpc.NewTaskDispatcherClient(conn)
	taskClient, err := client.GetTask(context.Background(), &serversiderpc.RequestInfo{
		Name: "request1",
	})
	if err != nil {
		panic(err)
	}
	for {
		info, err := taskClient.Recv()
		if err == io.EOF {
			log.Printf("[grpc server-side stream client] receive end\n")
			break
		}
		if err != nil {
			panic(err)
		}
		log.Printf("[grpc server-side stream client] receive message:%v\n", *info)
	}
}
