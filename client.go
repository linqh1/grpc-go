package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"myProtobuf/myproto"
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
	message1 := getSendMessage()
	bytes1, _ := proto.Marshal(message1)
	msgLen1 := len(bytes1)
	message2 := getSendMessage()
	message2.TaskId = "taskid2"
	bytes2, _ := proto.Marshal(message2)
	msgLen2 := len(bytes2)
	fmt.Printf("[client] to send protobuf object:%#v\n.length:%v\n", message1, msgLen1)
	fmt.Printf("[client] to send protobuf object:%#v\n.length:%v\n", message2, msgLen2)
	for ; true; <-ticker.C {
		log.Println("[client] spilt protobuf data to 2 part")
		log.Println("[client] send first part")
		conn.Write(proto.EncodeVarint(uint64(msgLen1)))
		conn.Write(bytes1[:msgLen1/2]) // 拆包
		time.Sleep(time.Second)
		log.Println("[client] send second part. and append some extra bytes")
		conn.Write(append(bytes1[msgLen1/2:], append(proto.EncodeVarint(uint64(msgLen2)), bytes2[:msgLen2/2]...)...)) //粘包
		conn.Write(append(bytes2[msgLen2/2:], []byte{1, 3, 5, 7}...))
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

func getSendMessage() *myproto.DispatchTask {
	result := &myproto.DispatchTask{
		TaskId: "firstTaskId",
		Details: []*myproto.PurgeTask{
			{PurgeId: "1",
				Type:    "dir",
				Content: []string{"http://cp.quantil.com/a", "http://cp.quantil.com/b"},
				Headers: map[string]string{"type": "dir", "type1": "dir1"},
				Action:  "delete",
			},
			{PurgeId: "2",
				Type:    "file",
				Content: []string{"http://cp.quantil.com/icon.png", "http://cp.quantil.com/serverfile.txt"},
				Headers: map[string]string{"type": "file", "type1": "file1"},
				Action:  "prefetch",
			},
		},
	}
	return result
}
