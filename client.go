package main

import "github.com/Scalingo/go-workers"

type Message struct {
    Name string
    Body string
    Time int64
}

func main() {
	workers.Configure(map[string]string{
		"process": "client1",
		"server":  "localhost:6379",
	})
	m := Message{"Alice", "Hello", 1294706395881547000}
	// b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	workers.Enqueue("myqueue", "MyGoWorker", m)
}
