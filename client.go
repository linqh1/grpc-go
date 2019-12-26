package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"myProtobuf/score_server"
	"net"
	"os"
	"time"
)

var targetAddr = "127.0.0.1:8888"
var stringArr []string

func main() {
	log.Println("[client] start")
	StartClient()
}

func StartClient() {
	conn, err := net.DialTimeout("tcp4", targetAddr, time.Second)
	if err != nil {
		panic(err)
	}
	go readConn(conn)
	ticker := time.NewTicker(time.Hour)
	message := getSendMessage()
	fmt.Printf("[client] to send protobuf object:%#v\n", message)
	bytes, _ := proto.Marshal(message)
	msgLen := len(bytes)
	for ; true; <-ticker.C {
		log.Println("[client] spilt protobuf data to 2 part")
		log.Println("[client] send first part")
		conn.Write(proto.EncodeVarint(uint64(msgLen)))
		conn.Write(bytes[:msgLen/2])
		time.Sleep(time.Second)
		log.Println("[client] send second part. and append some extra bytes")
		conn.Write(append(bytes[msgLen/2:], []byte{1, 3, 4, 6, 7}...))
	}
}

func readConn(conn net.Conn) {
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil && err != io.EOF && !os.IsTimeout(err) {
			panic(err)
		}
		if err != io.EOF {
			log.Printf("[clinet] received message:%v\n", string(tmp[:n]))
		}
	}
}

func init() {
	for start := 'a'; start <= 'z'; start++ {
		stringArr = append(stringArr, string([]byte{byte(start)}))
	}
	stringArr = append(stringArr, ";", "-", "_")
	fmt.Println(stringArr)
}

func getSendMessage() *score_server.BaseScoreInfoT {
	score_info := &score_server.BaseScoreInfoT{}
	score_info.WinCount = new(int32)
	*score_info.WinCount = 1
	score_info.LoseCount = new(int32)
	*score_info.LoseCount = 2
	score_info.ExceptionCount = new(int32)
	*score_info.ExceptionCount = 3
	score_info.KillCount = new(int32)
	*score_info.KillCount = 4
	score_info.DeathCount = new(int32)
	*score_info.DeathCount = 5
	score_info.AssistCount = new(int32)
	*score_info.AssistCount = 6
	score_info.Rating = new(int64)
	*score_info.Rating = 1800
	return score_info
}
