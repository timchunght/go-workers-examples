package main

import (
	// "fmt"
	"github.com/jrallison/go-workers"
	// "reflect"
	// "time"
)

func main() {
	workers.Configure(map[string]string{
		"process": "client1",
		"server":  "localhost:6379",
	})
	workers.EnqueueIn("default", "WhiteboardCacheWorker", 2*60, []string{})

	// now := time.Now()
	// hoursTo9 := (time.Duration(9 - now.Hour())) * time.Hour
	// at := now.AddDate(0, 0, 1).Truncate(time.Hour).Add(hoursTo9).Unix()
	// fmt.Println(reflect.TypeOf(at))
	// at := time.Date(2016, 4, 12, hour, min, sec, nsec int, loc *Location) Time
	// workers.EnqueueAt("default", "WhiteboardCacheWorker", now, nil)
}
