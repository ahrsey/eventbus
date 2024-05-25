package eventbus

import (
  "fmt"
  "time"
)

func main() {
	bus := MakeBroker()
	bus.Subscribe("topic", handler)
	bus.Subscribe("topic", handler2)
	bus.Subscribe("topic1", handler2)
	bus.Subscribe("topic1", handler)

	bus.QueuePublish("topic", "Triggered handler 1")
	bus.QueuePublish("topic1", "Triggered handler 2")
	bus.DrainQueue()
}

func handler(str string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("1: Called with a 2s sleep. %s\n", str)
}

func handler2(str string) {
	fmt.Printf("2: Instantly called. %s\n", str)
}
