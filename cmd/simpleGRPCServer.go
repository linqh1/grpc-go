package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"grpc-go/proto/simple"
	"log"
	"net"
)

type myserver struct {
}

func (myserver) SayHello(context context.Context, req *simple.HelloRequest) (resp *simple.HelloReply, err error) {
	p, ok := peer.FromContext(context)
	if ok {
		log.Printf("[grpc server] receive from: %v\n", p.Addr.String())
	}
	md, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("[grpc server] client metadata:\n%v", md)
	}
	log.Printf("[grpc server] request has been received successfully: %v\n", *req)
	resp = &simple.HelloReply{
		Message: req.Name + " has been received",
	}
	return
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.MaxRecvMsgSize(100*1024*1024), grpc.MaxSendMsgSize(100*1024*1024))
	simple.RegisterGreeterServer(s, new(myserver))
	s.Serve(lis)
}
