package main

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

// InputMove is the representation for the input to
// workflow.move
type WFStatesAllInput struct {
	Issue *Issue `json:"issue"`
}

// WFMove holds all related logic for event workflow.move
type WFStatesAll struct{}

func (w *WFStatesAll) Subscribe(nc *nats.Conn) {
	e := ErrorMessage{}
	nc.Subscribe("workflow.states.all", func(m *nats.Msg) {
		err, i := w.mapInput(m.Data)
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
func (w *WFStatesAll) mapInput(body []byte) (error, *WFStatesAllInput) {
	input := WFStatesAllInput{}
	if err := json.Unmarshal(body, &input); err != nil {
		return err, nil
	}

	return nil, &input
}
