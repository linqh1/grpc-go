package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"myProtobuf/grpcservice"
	"time"
)

var list = []string{"shindou", "中文", "kafka", "otter", "golang"}

func main() {
	conn, err := grpc.Dial("hub.qtlcdn.com:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcclient := grpcservice.NewGreeterClient(conn)
	for _, v := range list {
		go runClient(grpcclient, v)
	}
	select {}
}

func runClient(grpcclient grpcservice.GreeterClient, name string) {
	for {
		reply, err := grpcclient.SayHello(context.TODO(), &grpcservice.HelloRequest{Name: name})
		if err != nil {
			log.Printf("[grpc client] grpc call error: %#v\n", err)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			continue
		}
		log.Printf("[grpc client]received grpc server reply: %#v\n", reply.Message)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
