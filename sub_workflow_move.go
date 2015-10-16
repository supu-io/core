package main

import (
	"encoding/json"
	"log"

	"github.com/adriacidre/fsm"
	"github.com/nats-io/nats"
	"github.com/supu-io/messages"
)

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
			log.Println(string(m.Data))
			log.Println(err.Error())
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		i, err := w.executeTransition(input)
		if err != nil {
			log.Println(err.Error())
			e.Error = err.Error()
			nc.Publish(m.Reply, e.toJSON())
			return
		}

		nc.Publish(m.Reply, ToJSON(i))
	})
}

// Executes a transition for a given input
func (w *WFMove) executeTransition(input *messages.UpdateIssue) (*messages.UpdateIssue, error) {
	e := fsm.State(input.Status)
	i := Issue{
		ID:    input.Issue.ID,
		State: fsm.State(input.Issue.Status),
	}
	err := i.Apply(input.Status).Transition(e)
	if err != nil {
		return nil, err
	}

	i.Config = i.Config
	Hook(input.Issue, input.Config, string(i.State))

	input.Issue.Status = string(i.State)

	return input, nil
}

// Maps the json input to an InputMove structure
func (w *WFMove) mapInput(body string) (*messages.UpdateIssue, error) {
	input := messages.UpdateIssue{}
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		return nil, err
	}

	return &input, nil
}
