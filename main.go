package main

import (
	"github.com/nats-io/nats"
	"runtime"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	s := Subscriber{}
	s.subscribe(nc)
	runtime.Goexit()
}
