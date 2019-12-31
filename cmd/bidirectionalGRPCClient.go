package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"myProtobuf/proto/bidirectional"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//waitc := make(chan struct{})
	client := bidirectional.NewChatServerClient(conn)
	for {
		err := chat(client)
		if err == io.EOF {
			log.Printf("[client] communication end with EOF\n")
		} else if err != nil {
			log.Printf("[client] communication error:%v\n", err)
		}
		time.Sleep(5 * time.Second)
	}
	//taskClient.CloseSend()
	//<-waitc
}

func chat(client bidirectional.ChatServerClient) error {
	log.Printf("[client] start chat with server\n")
	taskClient, err := client.GetTask(context.Background())
	if err != nil {
		log.Printf("[client] start chat with server error:%v\n", err)
		return err
	}
	e := make(chan error)
	go func() {
		fmt.Println("=====enter goroutine1======")
		for {
			in, err := taskClient.Recv()
			if err == io.EOF {
				log.Printf("[client] received EOF\n")
				//close(waitc)
				// 因为有其他的goroutine也在操作e，所以可能e<-io.EOF会阻塞。导致线程退出不了
				// 所以加上default,
				select {
				case e <- io.EOF:
				default:
				}
				fmt.Println("=====exit goroutine1======")
				return
			}
			if err != nil {
				log.Printf("[client] Failed to receive: %v", err)
				select {
				case e <- err:
				default:
				}
				fmt.Println("=====exit goroutine1======")
				return
			} else {
				log.Printf("[client] received:%v\n", *in)
			}
		}
	}()
	go func() {
		fmt.Println("=====enter goroutine2======")
		message := []string{"hello", "how do you do", "good bye"}
		i := 0
		for {
			log.Printf("[client] send %v:\n", message[i])
			err := taskClient.Send(&bidirectional.ChatInfo{Name: message[i]})
			if err == io.EOF {
				log.Printf("[client] send return EOF\n")
				select {
				case e <- io.EOF:
				default:
				}
				fmt.Println("=====exit goroutine2======")
				break
			}
			if err != nil {
				log.Printf("[client] send error:%v\n", err)
				select {
				case e <- err:
				default:
				}
				fmt.Println("=====exit goroutine2======")
				break
			}
			i = (i + 1) % len(message)
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return <-e
}
