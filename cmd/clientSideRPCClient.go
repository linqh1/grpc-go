package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"myProtobuf/proto/clientside"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := clientside.NewFileBatchUploadClient(conn)
	for i := 0; i < 10; i++ {
		go run1(i, client)
	}
	select {}
}

func run1(goroutineIndex int, client clientside.FileBatchUploadClient) {
	uploadClient, err := client.Upload(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range []string{"1.png", "2.css", "3.gif"} {
		log.Printf("[grpc client-side stream client-%v] send:%v\n", goroutineIndex, v)
		err := uploadClient.Send(&clientside.FileInfo{
			Name: v,
		})
		if err == io.EOF {
			log.Printf("[grpc client-side stream client-%v] receive end:%v\n", goroutineIndex, err)
			break
		}
		if err != nil {
			log.Printf("[grpc client-side stream client-%v] send error:%v\n", goroutineIndex, err)
			break
		}
	}
	reply, err := uploadClient.CloseAndRecv()
	if err == io.EOF {
		log.Printf("[grpc client-side stream client-%v] CloseAndRecv success\n", goroutineIndex)
	} else if err != nil {
		log.Fatalf("[grpc client-side stream client-%v] %v.CloseAndRecv() got error %v, want %v", goroutineIndex, reply, err, nil)
	} else {
		log.Printf("[grpc client-side stream client-%v] receive: %v\n", goroutineIndex, *reply)
	}
}
