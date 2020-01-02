package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"
	"log"
	"myProtobuf/proto/simple"
	"net"
	"time"
)

type myserver struct {
}

func (myserver) SayHello(context context.Context, req *simple.HelloRequest) (resp *simple.HelloReply, err error) {
	p, ok := peer.FromContext(context)
	if ok {
		log.Printf("[grpc server] receive from: %v\n", p.Addr.String())
	}
	log.Printf("[grpc server] request has been received successfully: %v\n", *req)
	resp = &simple.HelloReply{
		Message: req.Name + " has been received",
	}
	return
}

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	simple.RegisterGreeterServer(s, new(myserver))
	s.Serve(lis)
}
