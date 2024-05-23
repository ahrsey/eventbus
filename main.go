package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	topic string // An event can have a single topic
	body  string // [1]
}

// [1] An event body which will be added to the event when published and passed
// to the subscriber

type Subscriber struct {
	id     string            // Used to store subscribers in a map
	events chan *Event       // To listen for events
	topics map[string]string // topics that this subscriber is subbed to
}

type Broker struct {
	subcribers map[string]string // Map of subscribers by id
	topics     map[string]string // [2]
}

// [2] Map of topics that subscribers can subbed to, this will be used for
// message filtering

type subscribers []Subscriber

// TODO: Create structs: Event, Subscriber, Broker
// TODO: Create interfaces: Event, Subscriber, Broker
func main() {
	e := Event{}
	subscribers := []Subscriber{}

	sub1 := Subscriber{"sub1", func(wg *sync.WaitGroup, e Event) {
		defer wg.Done()
		d, _ := time.ParseDuration("3s")
		time.Sleep(d)
		fmt.Printf("Inside sub 1\n")
	}}
	subscribers = append(subscribers, sub1)

	sub2 := Subscriber{"sub2", func(wg *sync.WaitGroup, e Event) {
		defer wg.Done()
		d, _ := time.ParseDuration("1s")
		time.Sleep(d)
		fmt.Printf("Inside sub 2\n")
	}}
	subscribers = append(subscribers, sub2)

	publish(subscribers, e)
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
