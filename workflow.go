package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/adriacidre/fsm"
)

// Transition json representation
type Transition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Workflow json representation
type Workflow struct {
	Transitions *[]Transition `json:"transitions"`
}

// Workflow definition
// This method describes the default workflow for an issue
func (w *Workflow) workflowRules() *fsm.Ruleset {
	if w == nil || w.Transitions == nil {
		w = w.load("workflows/default.json")
	}
	rules := fsm.Ruleset{}
	for _, t := range *w.Transitions {
		rules.AddTransition(fsm.T{fsm.State(t.From), fsm.State(t.To)})
	}

	return &rules
}

func (w *Workflow) load(source string) *Workflow {
	file, err := os.Open(source)
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&w)
	if err != nil {
		log.Println("Workflow " + source + " not found")
		log.Panic("error:", err)
	}
	return w
}
