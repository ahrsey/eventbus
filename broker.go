package eventbus

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []interface{}
}

type Broker struct {
	subscribers map[string][]func(str string)
	queue       *Queue
}

func MakeBroker() *Broker {
	subscribers := make(map[string][]func(str string))
	queue := &Queue{}
	broker := Broker{subscribers, queue}
	fmt.Printf("[INFO] Making Broker%s\n", broker)
	return &broker
}

func (b Broker) Subscribe(topic string, fn func(str string)) {
	b.subscribers[topic] = append(b.subscribers[topic], fn)

	_, ok := b.subscribers[topic]
	if ok {
		fmt.Printf("[INFO] Added subscriber %s to topic of `%s`\n", fn, topic)
	}
}

func (b Broker) QueuePublish(topic, args string) {
	insert := [2]string{topic, args}

	b.queue.items = append(b.queue.items, insert)

	for k, v := range b.queue.items {
		fmt.Printf("[INFO] queue currently has %s, %s\n", k, v)
	}
}

func (b Broker) DrainQueue() {
	var wg sync.WaitGroup
	queue := b.queue.items
	fmt.Printf("[INFO] Queue size %s\n", len(queue))

	for _, v := range queue {
		value, ok := v.([2]string)

		if !ok {
			fmt.Printf("[ERROR] value was not type of [2]string\n")
			return
		}
		topic := value[0]
		args := value[1]

		fmt.Printf("[INFO] topic %s\n", topic)
		fmt.Printf("[INFO] args %s\n", args)

		fns, ok := b.subscribers[topic]
		if !ok {
			fmt.Printf("[ERROR] Tried to fetch subscribers for topic: `%s` but none were found.\n", topic)
		}

		wg.Add(len(fns))
		for _, fn := range fns {
			go func(f func(str string)) {
				defer wg.Done()
				f(args)
			}(fn)
		}
	}

	fmt.Printf("[INFO] Queue has been drained\n")
	wg.Wait()
}
