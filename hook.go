package main

import (
	"github.com/nats-io/nats"
	"github.com/supu-io/messages"
)

// Hook manager
// When a issue moves from an state to the next one, a
// hook is executed, here you can find the mapped status
// to its hooks
func Hook(i *messages.Issue, c messages.Config, to string) error {
	from := i.Status
	updateIssueTrackerState(*i, c, to)
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

func updateIssueTrackerState(i messages.Issue, c messages.Config, to string) {
	i.Status = to
	msg := messages.UpdateIssue{
		Issue:  &i,
		Config: c,
		Status: to,
	}
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Publish("issues.update", ToJSON(msg))
}

// Hook executed when a issue moves to "In Progress".
func toInProgressHook(from string) {
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
func toCiHook(from string) {
	inProgressToCiHook()
}

// Hook executed when a issue moves to "In Review"
func toInReviewHook(from string) {
	ciToInReviewHook()
}

// Hook kexecuted when a issue moves to "UAT"
func toUatHook(from string) {
	inReviewToUatHook()
}

// Hook executed when a issue moves to "Done".
func toDoneHook(from string) {
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
