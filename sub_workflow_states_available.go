package main

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

// WFStatesAvailableInput is the representation for the input to
// workflow.status.available
type WFStatesAvailableInput struct {
	Issue  *Issue `json:"issue"`
	Status string `json:"status"`
}

// WFStatesAvailable holds all related logic for event workflow.states.availabe
type WFStatesAvailable struct{}

func (w *WFStatesAvailable) Subscribe(nc *nats.Conn) {
	e := ErrorMessage{}
	nc.Subscribe("workflow.states.available", func(m *nats.Msg) {
		err, i := w.mapInput(m.Data)
		if err != nil {
			e.Error = err.Error()
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		s := i.Issue.AvailableExitStates()
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
func (w *WFStatesAvailable) mapInput(body []byte) (error, *WFStatesAvailableInput) {
	input := WFStatesAvailableInput{}
	if err := json.Unmarshal(body, &input); err != nil {
		return err, nil
	}

	return nil, &input
}
