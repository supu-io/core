package main

import (
	"github.com/adriacidre/fsm"
	"github.com/nats-io/nats"
)

// Hook manager
// When a issue moves from an state to the next one, a
// hook is executed, here you can find the mapped status
// to its hooks
func Hook(i *Issue, to fsm.State) error {
	from := i.State
	updateIssueTrackerState(*i, to)
	switch to {
	case "in_progress":
		toInProgressHook(from)
	case "ci":
		toCiHook(from)
	case "in_review":
		toInReviewHook(from)
	case "uat":
		toUatHook(from)
	case "done":
		toDoneHook(from)
	}
	return nil
}

func updateIssueTrackerState(i Issue, to fsm.State) {
	i.State = to
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Publish("issues.update", *i.toJSON())
}

// Hook executed when a issue moves to "In Progress".
func toInProgressHook(from fsm.State) {
	switch from {
	case "created":
		createdToInProgressHook()
	case "in_review":
		inReviewToInProgressHook()
	case "ci":
		ciToInProgressHook()
	case "uat":
		uatToInProgressHook()
	}
}

// Hook executed when a issue moves to "CI"
func toCiHook(from fsm.State) {
	inProgressToCiHook()
}

// Hook executed when a issue moves to "In Review"
func toInReviewHook(from fsm.State) {
	ciToInReviewHook()
}

// Hook kexecuted when a issue moves to "UAT"
func toUatHook(from fsm.State) {
	inReviewToUatHook()
}

// Hook executed when a issue moves to "Done".
func toDoneHook(from fsm.State) {
	uatToDoneHook()
}

// The developer starts to work on a specific issue
func createdToInProgressHook() {
	// Update issue status on the issue manager
	// Notify the team current user is starting this task
}

// Event launched when a pull request is opened and
// every time new commits happen
func inProgressToCiHook()       {}
func ciToInProgressHook()       {}
func ciToInReviewHook()         {}
func inReviewToUatHook()        {}
func inReviewToInProgressHook() {}
func uatToInProgressHook()      {}
func uatToDoneHook()            {}
