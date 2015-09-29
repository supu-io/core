package main

import (
	"encoding/json"
	"github.com/adriacidre/fsm"
	"log"
)

// Struct representing an Issue
type Issue struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Repo  string `json:"repo"`
	Url   string `json:"url"`
	State fsm.State

	// our machine cache
	machine *fsm.Machine
}

// Get json representation of an issue
func (i *Issue) toJSON() *[]byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return &json
}

// Add methods to accomplish the fsm.Stater interface
func (t *Issue) CurrentState() fsm.State { return t.State }

func (t *Issue) SetState(s fsm.State) {
	t.State = s
}

// A helpful function that lets us apply arbitrary rulesets to this
// instances state machine without reallocating the machine. While not
// required, it's something I like to have.
func (t *Issue) Apply(w *Workflow, s string) *fsm.Machine {
	r := w.workflowRules()
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}
	w.Hook(t.State, fsm.State(s))

	t.machine.Rules = r
	return t.machine
}

// Get all possible states for this issue given its particular
// workflow
func (t *Issue) AllStates() []string {
	states := []string{}
	for s := range t.machine.Rules.AllStates() {
		states = append(states, string(s))
	}
	return states
}

// Get all possible states for this issue given its particular
// workflow and current status
func (t *Issue) AvailableExitStates() []string {
	states := []string{}
	for s := range t.machine.Rules.AvailableExit(t.State) {
		states = append(states, string(s))
	}
	return states
}
