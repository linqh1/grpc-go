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
				log.Printf("[client] received EOF\n")
				//close(waitc)
				return
			}
			if err != nil {
				log.Printf("[client] Failed to receive: %v", err)
				return
			} else {
				log.Printf("[client] received:%v\n", *in)
			}
		}
	}()
	message := []string{"hello", "how do you do", "good bye"}
	i := 0
	for {
		log.Printf("[client] send %v:\n", message[i])
		err := taskClient.Send(&bidirectional.ChatInfo{Name: message[i]})
		if err == io.EOF {
			log.Printf("[client] send return EOF\n")
			break
		}
		if err != nil {
			log.Printf("[client] send error:%v\n", err)
			break
		}
		i = (i + 1) % len(message)
		time.Sleep(200 * time.Millisecond)
	}
	//taskClient.CloseSend()
	//<-waitc
}
