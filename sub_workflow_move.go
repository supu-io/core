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

// WFMove holds all related logic for event workflow.move
type WFMove struct{}

// Subscribe to workflow.move in order to move an issue to its
// new status and execute proper hooks
func (w *WFMove) Subscribe(nc *nats.Conn) {
	e := ErrorMessage{}
	nc.Subscribe("workflow.move", func(m *nats.Msg) {
		input, err := w.mapInput(string(m.Data))
		if err != nil {
			e.Error = err.Error()
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		i, err := w.executeTransition(input)
		if err != nil {
			e.Error = err.Error()
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		nc.Publish(m.Reply, *i.toJSON())
	})
}

// Executes a transition for a given input
func (w *WFMove) executeTransition(i *InputMove) (*Issue, error) {
	e := fsm.State(i.State)
	err := i.Issue.Apply(i.State).Transition(e)
	if err != nil {
		return nil, err
	}
	i.Issue.State = fsm.State(i.State)

	return i.Issue, nil
}

// Maps the json input to an InputMove structure
func (w *WFMove) mapInput(body string) (*InputMove, error) {
	input := InputMove{}
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		return nil, err
	}

	return &input, nil
}
