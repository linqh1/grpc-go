package main

import (
	"google.golang.org/grpc"
	"log"
	"myProtobuf/serversiderpc"
	"net"
	"time"
)

type serversideserver struct {
}

func (serversideserver) GetTask(req *serversiderpc.RequestInfo, server serversiderpc.TaskDispatcher_GetTaskServer) error {
	err := server.Send(&serversiderpc.TaskInfo{
		Message: "first message",
	})
	if err != nil {
		log.Printf("[grpc server-side stream server] send error:%v\n", err)
	}
	time.Sleep(5 * time.Second)
	err = server.Send(&serversiderpc.TaskInfo{
		Message: "second message",
	})
	if err != nil {
		log.Printf("[grpc server-side stream server] send error:%v\n", err)
	}
	go func() {
		time.Sleep(2 * time.Second)
		goerr := server.Send(&serversiderpc.TaskInfo{ // return 完后再写入会出错
			Message: "third message send in go routine",
		})
		if goerr != nil {
			log.Printf("[grpc server-side stream server] send error:%v\n", goerr)
		}
	}()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("[grpc server-side stream server] listen on port:%v\n", "8888")
	s := grpc.NewServer()
	serversiderpc.RegisterTaskDispatcherServer(s, new(serversideserver))
	s.Serve(lis)
}
