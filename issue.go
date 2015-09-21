package main

import (
	"encoding/json"
	"github.com/adriacidre/fsm"
	"log"
)

type Issue struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Repo  string `json:"repo"`
	Url   string `json:"url"`
	State fsm.State

	// our machine cache
	machine *fsm.Machine
}

func (i *Issue) toJSON() *[]byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}
	return &json
}

// Add methods to comply with the fsm.Stater interface
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

func (i *Issue) toJson() string {
	json, _ := json.Marshal(i)
	return string(json)
}
