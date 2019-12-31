package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"myProtobuf/bidirectionalrpc"
	"net"
)

type doublesideserver struct{}

func (doublesideserver) GetTask(server bidirectionalrpc.ChatServer_GetTaskServer) error {
	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("[server] receive:%v\n", *in)
		for _, str := range []string{"hello", "what your name", "good bye"} {
			if err := server.Send(&bidirectionalrpc.ChatInfo{Name: str}); err != nil {
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("[grpc server-side stream server] listen on port:%v\n", "8888")
	s := grpc.NewServer()
	bidirectionalrpc.RegisterChatServerServer(s, new(doublesideserver))
	s.Serve(lis)
}
