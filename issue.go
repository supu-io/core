package main

import (
	"encoding/json"
	"github.com/adriacidre/fsm"
	"log"
)

// Issue representation
type Issue struct {
	ID       string    `json:"id"`
	State    fsm.State `json:"status"`
	Workflow *Workflow `json:"workflow"`

	// our machine cache
	machine *fsm.Machine
}

// Get json representation of an issue
func (t *Issue) toJSON() *[]byte {
	json, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
	}
	return &json
}

// CurrentState get the current state for the issue
func (t *Issue) CurrentState() fsm.State { return t.State }

// SetState sets the state for the given issue
func (t *Issue) SetState(s fsm.State) {
	t.State = s
}

// Apply arbitrary rules to each transition
func (t *Issue) Apply(s string) *fsm.Machine {
	w := t.Workflow
	r := w.workflowRules()
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}
	w.Hook(t.State, fsm.State(s))

	t.machine.Rules = r
	return t.machine
}

// AllStates Get all possible states for this issue given its
// particular workflow
func (t *Issue) AllStates() []string {
	w := t.Workflow
	r := w.workflowRules()
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}
	t.machine.Rules = r

	states := []string{}
	for _, s := range t.machine.Rules.AllStates() {
		states = append(states, string(s))
	}
	return states
}

// AvailableExitStates Gets all possible states for this issue
// given its particular workflow and current status
func (t *Issue) AvailableExitStates() []string {
	w := t.Workflow
	r := w.workflowRules()
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}
	t.machine.Rules = r

	states := []string{}
	for s := range t.machine.Rules.AvailableExit(t.State) {
		states = append(states, string(s))
	}
	return states
}
