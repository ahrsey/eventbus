package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	bus := NewBroker()
	mux := http.NewServeMux()

	bus.Subscribe("log", log)
	mux.HandleFunc("/", handlePublish(bus))

	err := http.ListenAndServe(":3333", mux)
	if err == nil {
		fmt.Printf("[INFO] the following issue occured while running the http server `%s`\n", err)
	}
}

func log(e *Event) {
	fmt.Printf("[INFO] Log handler called with `%s`\n", e)
}

func handlePublish(bus *Broker) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		topic := r.URL.Query().Get("topic")
		body := r.URL.Query().Get("body")

		e := Event{topic, body}

		if topic != "" {
			fmt.Printf("[INFO] Http publishing topic of `%s`\n", topic)
			fmt.Printf("[INFO] Http publishing body of `%s`\n", body)
			bus.QueuePublish(&e)
			bus.DrainQueue()
		} else {
			fmt.Printf("[ERROR] Http topic required, but found `%s`\n", topic)
		}

		io.WriteString(w, "Topic published.\n")
	}

	return http.HandlerFunc(fn)
}
