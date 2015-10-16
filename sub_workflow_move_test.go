package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nats-io/nats"
	"github.com/supu-io/messages"
)

var w Subscriber
var nc *nats.Conn

func setup() *nats.Conn {
	nc, _ := nats.Connect(nats.DefaultURL)
	w = Subscriber{}
	w.subscribe(nc)
	return nc
}

func TestWorkflowMove(t *testing.T) {
	nc = setup()
	msg := messages.UpdateIssue{
		Issue: &messages.Issue{
			ID:     "foo",
			Status: "created",
		},
		Status: "todo",
		Config: messages.Config{},
	}
	body, err := json.Marshal(msg)
	res, err := nc.Request("workflow.move", body, 10000*time.Millisecond)

	i := messages.UpdateIssue{}
	err = json.Unmarshal(res.Data, &i)
	if err != nil {
		t.Error(err.Error())
	}

	println(i.Status)
	if i.Status != "todo" {
		t.Error("Didn't happen transition created -> todo")
	}
}

func TestWorkflowStatesAll(t *testing.T) {
	nc = setup()
	body := []byte(`{"issue":{"id":"foo","status":"created"}}`)
	res, err := nc.Request("workflow.states.all", body, 10000*time.Millisecond)

	s := []string{}
	err = json.Unmarshal(res.Data, &s)
	if err != nil {
		t.Error(err.Error())
	}

	println(len(s))
	if len(s) != 6 {
		t.Error("Invalid number of status returned")
	}
}

func TestWorkflowStatesAvailableExit(t *testing.T) {
	nc = setup()
	body := []byte(`{"issue":{"id":"foo","status":"created"}, "status": "in_progress"}`)
	res, err := nc.Request("workflow.states.available", body, 10000*time.Millisecond)

	s := []string{}
	err = json.Unmarshal(res.Data, &s)
	if err != nil {
		t.Error(err.Error())
	}

	if len(s) != 1 {
		t.Error("Invalid number of status returned")
		println(len(s))
	}
}
