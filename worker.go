package main

import (
	"fmt"
	"github.com/Scalingo/go-workers"
	"encoding/json"
	// "reflect"
	// "github.com/bitly/go-simplejson"
)

func main() {
	workers.Configure(map[string]string{
		"process": "worker1",
		"server":  "localhost:6379",
	})

	workers.Process("myqueue", MyGoWorker, 10)
	workers.Run()
}
	
type Message struct {
    Name string
    Body string
    Time int64
}

func MyGoWorker(msg *workers.Msg) {
	args := msg.Args().ToJson()	

	var m Message
	err := json.Unmarshal([]byte(args), &m)
	if(err != nil) {
		fmt.Println("unmarshal error")
	}
	fmt.Println(m.Name)
	
}
