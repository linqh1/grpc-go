package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"grpc-go/proto/clientside"
	"io"
	"log"
	"net"
	"sync/atomic"
)

type clientsideserver struct {
}

var cnt int32

func (clientsideserver) Upload(server clientside.FileBatchUpload_UploadServer) error {
	index := atomic.AddInt32(&cnt, 1)
	p, ok := peer.FromContext(server.Context())
	if ok {
		log.Printf("[grpc client-side stream server-%v] receive from: %v\n", index, p.Addr.String())
	}
	for {
		info, err := server.Recv()
		if err == io.EOF {
			log.Printf("[grpc client-side stream server-%v] receive end\n", index)
			break
		}
		if err != nil {
			return err
		}
		log.Printf("[grpc client-side stream server-%v] receive %v\n", index, *info)
	}
	server.SendAndClose(&clientside.FileUploadResponse{Message: "receive success"})
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("[grpc client-side stream server] listen on port:%v\n", "8888")
	s := grpc.NewServer()
	clientside.RegisterFileBatchUploadServer(s, new(clientsideserver))
	s.Serve(lis)
}
