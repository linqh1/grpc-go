package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"myProtobuf/proto/simple"
	"time"
)

var list = []string{"shindou", "中文", "kafka", "otter", "golang"}

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	log.Printf("[grpc client] dial to server success.sleep 5 senconds and restart server")
	time.Sleep(5 * time.Second)
	grpcclient := simple.NewGreeterClient(conn)
	for _, v := range list {
		go runClient(grpcclient, v)
	}
	select {}
}

func runClient(grpcclient simple.GreeterClient, name string) {
	for {
		header := metadata.New(map[string]string{"name": "client", "type": "simple"})
		reply, err := grpcclient.SayHello(metadata.NewOutgoingContext(context.Background(), header), &simple.HelloRequest{Name: name})
		if err != nil {
			log.Printf("[grpc client] grpc call error: %#v\n", err)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			continue
		}
		log.Printf("[grpc client] received grpc server reply: %#v\n", reply.Message)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
