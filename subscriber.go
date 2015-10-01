package main

import (
	"encoding/json"
	"github.com/adriacidre/fsm"
	"github.com/nats-io/nats"
	"log"
	"runtime"
)

// Subscriber is listening for events happening on
// messaging system aka. nats.io
type Subscriber struct {
	Input *InputIssue
}

// InputIssue Every event received will, at least contain these two
// fields as part of the Json body
type InputIssue struct {
	Issue  *Issue `json:"issue"`
	Status string `json:"status,omitempty"`
}

// Subscribes to all needed events
func (sub *Subscriber) subscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)

	log.Println("Listening ...")

	nc.Subscribe("workflow.move", func(m *nats.Msg) {
		i := sub.manageIssue(string(m.Data))
		nc.Publish(m.Reply, *i.toJSON())
	})

	nc.Subscribe("workflow.states.all", func(m *nats.Msg) {
		i := sub.issueDetails(string(m.Data))
		s := i.AllStates()
		json, _ := json.Marshal(s)
		nc.Publish(m.Reply, json)
	})

	nc.Subscribe("workflow.states.available", func(m *nats.Msg) {
		i := sub.issueDetails(string(m.Data))
		s := i.AvailableExitStates()
		json, _ := json.Marshal(s)
		nc.Publish(m.Reply, json)
	})

	runtime.Goexit()
}

// Gets the issue details
func (sub *Subscriber) issueDetails(body string) *Issue {
	return sub.getIssueFromRequest(body)
}

// Manages the issue and executes the transition
func (sub *Subscriber) manageIssue(body string) *Issue {
	i := sub.getIssueFromRequest(body)
	if i == nil {
		return nil
	}
	e := fsm.State(sub.Input.Status)
	err := i.Apply(sub.Input.Status).Transition(e)
	if err != nil {
		log.Println(err)
		return nil
	}
	i.State = fsm.State(sub.Input.Status)

	return i
}

// Unmarshalls the event body into an Issue struct
func (sub *Subscriber) getIssueFromRequest(body string) *Issue {
	input := InputIssue{}
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		panic(err)
	}
	sub.Input = &input
	return input.Issue
}
