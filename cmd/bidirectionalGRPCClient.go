package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"myProtobuf/proto/bidirectional"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//waitc := make(chan struct{})
	client := bidirectional.NewChatServerClient(conn)
	taskClient, err := client.GetTask(context.Background())
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			in, err := taskClient.Recv()
			if err == io.EOF {
				//close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("[client] Failed to receive: %v", err)
			} else {
				log.Printf("[client] received:%v\n", *in)
			}
		}
	}()
	for {
		for _, v := range []string{"hello", "how do you do", "good bye"} {
			err := taskClient.Send(&bidirectional.ChatInfo{Name: v})
			if err != nil {
				log.Printf("[client] send error:%v\n", err)
			}
		}
		time.Sleep(200 * time.Millisecond)
	}
	//taskClient.CloseSend()
	//<-waitc
}
