package main

import (
  "fmt"
  "time"
  "sync"
)

type Subscriber struct {
  Name string
  Fn func(wg *sync.WaitGroup)
}

type subscribers []Subscriber

func main() {
  d, _ := time.ParseDuration("2s")
  subscribers := []Subscriber{}
  sub1 := Subscriber{"sub1", func(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Inside sub 1\n")
  }}

  sub2 := Subscriber{"sub2", func(wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(d)
    fmt.Printf("Inside sub 2\n")
  }}

  subscribers = append(subscribers, sub2)
  subscribers = append(subscribers, sub1)

  publish(subscribers)
}

func publish(s []Subscriber) {
  var wg sync.WaitGroup
  wg.Add(len(s))

  for _, sub := range s {
    fmt.Printf("%s\n", sub.Name)
    go sub.Fn(&wg)
  }

  wg.Wait()
}
