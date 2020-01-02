package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"io"
	"log"
	"myProtobuf/proto/clientside"
	"net"
)

type clientsideserver struct {
}

func (clientsideserver) Upload(server clientside.FileBatchUpload_UploadServer) error {
	p, ok := peer.FromContext(server.Context())
	if ok {
		log.Printf("[grpc client-side stream server] receive from: %v\n", p.Addr.String())
	}
	for {
		info, err := server.Recv()
		if err == io.EOF {
			log.Printf("[grpc client-side stream server] receive end\n")
			break
		}
		if err != nil {
			return err
		}
		log.Printf("[grpc client-side stream server] receive %v\n", *info)
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
