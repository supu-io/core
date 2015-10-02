package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nats-io/nats"
)

var w WFMove
var nc *nats.Conn

func setup() *nats.Conn {
	nc, _ := nats.Connect(nats.DefaultURL)
	w = WFMove{}
	w.Subscribe(nc)
	return nc
}

func TestSubscriber(t *testing.T) {
	nc = setup()
	body := []byte(`{"issue":{"id":"foo","status":"created"}, "status":"todo"}`)
	res, err := nc.Request("workflow.move", body, 10000*time.Millisecond)

	i := Issue{}
	err = json.Unmarshal(res.Data, &i)
	if err != nil {
		t.Error(err.Error())
	}

	if i.State != "todo" {
		t.Error("Didn't happen transition created -> todo")
	}
}
