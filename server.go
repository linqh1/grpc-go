package main

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"log"
	"myProtobuf/myproto"
	"net"
	"os"
)

type ReceiveBytes struct {
	headerSet bool
	header    []byte
	data      []byte
}

var listenAdd = ":8888"

func (r *ReceiveBytes) Add(newData []byte) ([]*myproto.DispatchTask, error) {
	var result []*myproto.DispatchTask
	for len(newData) > 0 {
		if !r.headerSet {
			b := newData[0]
			highestByte := (b & 0x80) >> 7
			r.header = append(r.header, b)
			if len(r.header) > 5 {
				return nil, errors.New("header length large than 5")
			}
			if highestByte == 0 {
				log.Printf("[server] finished message header")
				r.headerSet = true
			}
			newData = newData[1:]
		} else {
			length, _ := proto.DecodeVarint(r.header)
			if uint64(len(newData)) >= length-uint64(len(r.data)) {
				protobufByte := append(r.data, newData[0:length-uint64(len(r.data))]...)
				var dispatchTask = &myproto.DispatchTask{}
				err := proto.Unmarshal(protobufByte, dispatchTask)
				if err != nil {
					return result, err
				} else {
					log.Printf("[server] parse success")
				}
				result = append(result, dispatchTask)
				newData = newData[length-uint64(len(r.data)):]
				r.Init()
			} else {
				r.data = append(r.data, newData...)
				break
			}
		}
	}
	return result, nil
}

func (r *ReceiveBytes) Init() {
	r.headerSet = false
	r.header = nil
	r.data = nil
}

func StartServer() {
	listener, err := net.Listen("tcp4", listenAdd)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[server] listener accept error:%v\n", err)
			continue
		}
		log.Printf("[server] accept connection from %v\n", conn)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	var receive ReceiveBytes
	tmp := make([]byte, 10)
	for {
		n, err := conn.Read(tmp)
		if err != nil && !os.IsTimeout(err) { // 出错但不是超时
			log.Printf("[server] read error:%v\n", err)
			conn.Close()
			return
		}
		strings, err := receive.Add(tmp[:n])
		if err != nil {
			log.Printf("[server] parse data error:%v\n", err)
			if len(strings) > 0 {
				log.Printf("[server] but parse some message success.\n")
				for i := range strings {
					log.Printf("[server] receive:%v\n", strings[i])
				}
			}
			conn.Close()
			return
		}
		for i := range strings {
			log.Printf("[server] receive:%v\n", strings[i])
		}
	}
}

func main() {
	log.Println("server start")
	StartServer()
}
