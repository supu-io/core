package main

import (
	"github.com/adriacidre/fsm"
)

type Workflow struct {
}

// Workflow definition
// This method describes the default workflow for an issue
func (w *Workflow) workflowRules() *fsm.Ruleset {
	rules := fsm.Ruleset{}

	rules.AddTransition(fsm.T{"created", "in_progress"})
	rules.AddTransition(fsm.T{"in_progress", "ci"})
	rules.AddTransition(fsm.T{"ci", "in_progress"})
	rules.AddTransition(fsm.T{"ci", "in_review"})
	rules.AddTransition(fsm.T{"in_review", "in_progress"})
	rules.AddTransition(fsm.T{"in_review", "uat"})
	rules.AddTransition(fsm.T{"uat", "in_progress"})
	rules.AddTransition(fsm.T{"uat", "done"})

	return &rules
}

// Apply a transition to the given issue
func (w *Workflow) transact(issue *Issue, event string) error {
	e := fsm.State(event)
	return issue.Apply(w, event).Transition(e)
}

// Hook manager
// When a issue moves from an state to the next one, a
// hook is executed, here you can find the mapped status
// to its hooks
func (w *Workflow) Hook(from fsm.State, to fsm.State) error {
	switch to {
	case "in_progress":
		w.toInProgressHook(from)
	case "ci":
		w.toCiHook(from)
	case "in_review":
		w.toInReviewHook(from)
	case "uat":
		w.toUatHook(from)
	case "done":
		w.toDoneHook(from)
	}
	return nil
}

// Hook executed when a issue moves to "In Progress".
func (w *Workflow) toInProgressHook(from fsm.State) {
	switch from {
	case "created":
		w.createdToInProgressHook()
	case "in_review":
		w.inReviewToInProgressHook()
	case "ci":
		w.ciToInProgressHook()
	case "uat":
		w.uatToInProgressHook()
	}
}

// Hook executed when a issue moves to "CI"
func (w *Workflow) toCiHook(from fsm.State) {
	w.inProgressToCiHook()
}

// Hook executed when a issue moves to "In Review"
func (w *Workflow) toInReviewHook(from fsm.State) {
	w.ciToInReviewHook()
}

// Hook kexecuted when a issue moves to "UAT"
func (w *Workflow) toUatHook(from fsm.State) {
	w.inReviewToUatHook()
}

// Hook executed when a issue moves to "Done".
func (w *Workflow) toDoneHook(from fsm.State) {
	w.uatToDoneHook()
}

// The developer starts to work on a specific issue
func (w *Workflow) createdToInProgressHook() {
	// Update issue status on the issue manager
	// Notify the team current user is starting this task
}

// Event launched when a pull request is opened and
// every time new commits happen
func (w *Workflow) inProgressToCiHook()       {}
func (w *Workflow) ciToInProgressHook()       {}
func (w *Workflow) ciToInReviewHook()         {}
func (w *Workflow) inReviewToUatHook()        {}
func (w *Workflow) inReviewToInProgressHook() {}
func (w *Workflow) uatToInProgressHook()      {}
func (w *Workflow) uatToDoneHook()            {}
