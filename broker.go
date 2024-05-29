package main

import (
	"fmt"
	"sync"
)

type Event struct {
	topic string
	body  string
}

type Queue struct {
	items []interface{}
}

type Broker struct {
	subscribers map[string][]func(e *Event)
	queue       *Queue
}

func NewBroker() *Broker {
	subscribers := make(map[string][]func(e *Event))
	queue := &Queue{}
	broker := Broker{subscribers, queue}
	fmt.Printf("[INFO] Making Broker%s\n", broker)
	return &broker
}

func (b *Broker) Subscribe(topic string, fn func(e *Event)) {
	b.subscribers[topic] = append(b.subscribers[topic], fn)

	_, ok := b.subscribers[topic]
	if ok {
		fmt.Printf("[INFO] Added subscriber `%s` to topic of `%s`\n", fn, topic)
	}
}

func (b *Broker) QueuePublish(e *Event) {
	b.queue.items = append(b.queue.items, e)

	for k, v := range b.queue.items {
		fmt.Printf("[INFO] queue currently has `%s`, `%s`\n", k, v)
	}
}

func (b *Broker) DrainQueue() {
	var wg sync.WaitGroup
	queue := b.queue.items
	fmt.Printf("[INFO] Queue size `%s`\n", len(queue))

	for _, v := range queue {
		value, ok := v.(*Event)

		if !ok {
			fmt.Printf("[ERROR] value was not type of [2]string\n")
			return
		}

		topic := value.topic
		args := value.body

		fmt.Printf("[INFO] topic `%s`\n", topic)
		fmt.Printf("[INFO] args `%s`\n", args)

		fns, ok := b.subscribers[topic]
		if !ok {
			fmt.Printf("[ERROR] Tried to fetch subscribers for topic: `%s` but none were found.\n", topic)
		}

		wg.Add(len(fns))
		for _, fn := range fns {
			go func(f func(e *Event)) {
				defer wg.Done()
				f(value)
			}(fn)
		}
	}

	fmt.Printf("[INFO] Queue has been drained\n")
	wg.Wait()
}
