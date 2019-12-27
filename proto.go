package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"myProtobuf/myproto"
)

func main() {
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
	bytes, err := proto.Marshal(result)
	log.Println(bytes, len(bytes))
	if err != nil {
		log.Printf("Marshal error:%v\n", err)
		panic(err)
	}
	recoveredTask := &myproto.DispatchTask{}
	err = proto.Unmarshal(bytes, recoveredTask)
	if err != nil {
		log.Printf("Unmarshal error:%v\n", err)
		panic(err)
	}
	log.Printf(recoveredTask.TaskId)
	for i, v := range recoveredTask.Details {
		log.Printf("detail[%v]:%#v\n", i, *v)
	}
}
