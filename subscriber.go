package main

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"log"
)

// ErrorMessage representation of a json error message
type ErrorMessage struct {
	Error string `json:"error"`
}

// Get an error message as json
func (e *ErrorMessage) toJSON() []byte {
	json, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	return json
}

// Subscriber is listening for events happening on
// messaging system aka. nats.io
type Subscriber struct {
}

// Subscribes to all needed events
func (sub *Subscriber) subscribe(nc *nats.Conn) {
	log.Println("Listening ...")
	wm := WFMove{}
	wm.Subscribe(nc)

	ws := WFStatesAll{}
	ws.Subscribe(nc)

	wsa := WFStatesAvailable{}
	wsa.Subscribe(nc)
}
