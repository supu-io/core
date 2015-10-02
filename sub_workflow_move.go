package main

import (
	"encoding/json"

	"github.com/adriacidre/fsm"
	"github.com/nats-io/nats"
)

// InputMove is the representation for the input to
// workflow.move
type InputMove struct {
	Issue *Issue `json:"issue"`
	State string `json:"status,omitempty"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

// WFMove holds all related logic for event workflow.move
type WFMove struct{}

// Given a nats connection it subscribes to workflow.move in order
// to move an issue to its new status and execute proper hooks
func (w *WFMove) Subscribe(nc *nats.Conn) {
	nc.Subscribe("workflow.move", func(m *nats.Msg) {
		err, input := w.mapInput(string(m.Data))
		if err != nil {
			nc.Publish(m.Reply, []byte(`{"error":"`+err.Error()+`"}`))
			return
		}

		err, i := w.executeTransition(input)
		if err != nil {
			nc.Publish(m.Reply, []byte(`{"error":"`+err.Error()+`"}`))
			return
		}

		nc.Publish(m.Reply, *i.toJSON())
	})
}

// Executes a transition for a given input
func (w *WFMove) executeTransition(i *InputMove) (error, *Issue) {
	e := fsm.State(i.State)
	err := i.Issue.Apply(i.State).Transition(e)
	if err != nil {
		return err, nil
	}
	i.Issue.State = fsm.State(i.State)

	return nil, i.Issue
}

// Maps the json input to an InputMove structure
func (w *WFMove) mapInput(body string) (error, *InputMove) {
	input := InputMove{}
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		return err, nil
	}

	return nil, &input
}
