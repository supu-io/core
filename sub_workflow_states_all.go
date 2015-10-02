package main

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

// WFStatesAllInput is the representation for the input to
// workflow.states.all
type WFStatesAllInput struct {
	Issue *Issue `json:"issue"`
}

// WFStatesAll holds all related logic for event workflow.states.all
type WFStatesAll struct{}

// Subscribe to workflow.states.all and return all valid states
// for the given issue
func (w *WFStatesAll) Subscribe(nc *nats.Conn) {
	e := ErrorMessage{}
	nc.Subscribe("workflow.states.all", func(m *nats.Msg) {
		i, err := w.mapInput(m.Data)
		if err != nil {
			e.Error = err.Error()
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		s := i.Issue.AllStates()
		json, err := json.Marshal(s)
		if err != nil {
			e.Error = err.Error()
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		nc.Publish(m.Reply, json)
	})
}

// Maps the json input to an InputMove structure
func (w *WFStatesAll) mapInput(body []byte) (*WFStatesAllInput, error) {
	input := WFStatesAllInput{}
	if err := json.Unmarshal(body, &input); err != nil {
		return nil, err
	}

	return &input, nil
}
