package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"grpc-go/proto/bidirectional"
	"io"
	"log"
	"math/rand"
	"net"
)

type doublesideserver struct{}

func (doublesideserver) GetTask(server bidirectional.ChatServer_GetTaskServer) error {
	p, ok := peer.FromContext(server.Context())
	if ok {
		log.Printf("[server] receive from: %v\n", p.Addr.String())
	}
	md, ok := metadata.FromIncomingContext(server.Context())
	if ok {
		log.Printf("[server] metadata:\n%v", md)
	}
	index := rand.Int31()
	log.Printf("[server] [index=%v]enter loop", index)
	for {
		in, err := server.Recv()
		if err == io.EOF {
			log.Printf("[server] receive EOF\n")
			log.Printf("[server] [index=%v]exit loop", index)
			return nil
		}
		if err != nil {
			log.Printf("[server] receive error:%v\n", err)
			log.Printf("[server] [index=%v]exit loop", index)
			return err
		}
		log.Printf("[server] receive:%v\n", *in)
		if err := server.Send(&bidirectional.ChatInfo{Name: in.Name}); err != nil {
			log.Printf("[server] send error:%v\n", err)
			log.Printf("[server] [index=%v]exit loop", index)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("[server] listen on port:%v\n", "8888")
	s := grpc.NewServer()
	bidirectional.RegisterChatServerServer(s, new(doublesideserver))
	s.Serve(lis)
}
