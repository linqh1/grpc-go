package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"myProtobuf/clientsiderpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	uploadClient, err := clientsiderpc.NewFileBatchUploadClient(conn).Upload(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range []string{"1.png", "2.css", "3.gif"} {
		log.Printf("[grpc client-side stream client] send:%v\n", v)
		err := uploadClient.Send(&clientsiderpc.FileInfo{
			Name: v,
		})
		if err == io.EOF {
			log.Printf("[grpc client-side stream client] receive end:%v\n", err)
			break
		}
		if err != nil {
			log.Printf("[grpc client-side stream client] send error:%v\n", err)
			break
		}
	}
	reply, err := uploadClient.CloseAndRecv()
	if err == io.EOF {
		log.Printf("[grpc client-side stream client] CloseAndRecv success\n")
	} else if err != nil {
		log.Fatalf("[grpc client-side stream client] %v.CloseAndRecv() got error %v, want %v", reply, err, nil)
	} else {
		log.Printf("[grpc client-side stream client] receive: %v\n", *reply)
	}
}
