package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Event struct {
	name string
}

type Subscriber struct {
	name string
	fn   func(wg *sync.WaitGroup, e Event)
}

type subscribers []Subscriber

// TODO: Add sqlite3
// TODO: Publish events
// TODO: Filter subscriber triggers
// TODO: Access events data in subscribers
// TODO: When you subscribe it should create an event that is possible to be
// trigger to be used to trigger that subscriber
func main() {
  e := Event{}
  subscribers := []Subscriber{}

	sub1 := Subscriber{"sub1", func(wg *sync.WaitGroup, e Event) {
		defer wg.Done()
		d, _ := time.ParseDuration("1s")
		time.Sleep(d)
		fmt.Printf("Inside sub 1\n")
	}}

	sub2 := Subscriber{"sub2", func(wg *sync.WaitGroup, e Event) {
		defer wg.Done()
		d, _ := time.ParseDuration("3s")
		time.Sleep(d)
		fmt.Printf("Inside sub 2\n")
	}}

	subscribers = append(subscribers, sub2)
	subscribers = append(subscribers, sub1)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", postPublish)

	publish(subscribers, e)

//err := http.ListenAndServe(":6969", mux)
//if err != nil {
//  fmt.Println("%s", err)
//}
}

func publish(s []Subscriber, e Event) {
	var wg sync.WaitGroup
	wg.Add(len(s))

	for _, sub := range s {
		fmt.Printf("%s\n", sub.name)
		go sub.fn(&wg, e)
	}

	wg.Wait()
}

func postPublish(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var e Event
  err := decoder.Decode(&e)
  if err != nil {
    fmt.Println("%s", err)
  }

	// publish(subscribers, e)

	io.WriteString(w, "This is my website!\n")
}
