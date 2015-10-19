package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/supu-io/messages"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscriber(t *testing.T) {
	setup()
	Convey("Given a valid workflow.move message", t, func() {
		w := WFMove{}
		w.Subscribe(nc)
		msg := messages.UpdateIssue{
			Issue: &messages.Issue{
				ID:     "org/repo/1",
				Number: 1,
				Status: "created",
				Org:    "org",
				Repo:   "repo",
			},
			Status: "todo",
		}
		body, _ := json.Marshal(&msg)
		Convey("When send the message", func() {
			subscribeIssuesUpdate(t)
			response, err := nc.Request("workflow.move", body, 1000*time.Millisecond)
			Convey("Then issue state should be todo", func() {
				res := messages.UpdateIssue{}
				json.Unmarshal(response.Data, &res)

				So(res.Issue.ID, ShouldEqual, "org/repo/1")
				So(res.Issue.Org, ShouldEqual, "org")
				So(res.Issue.Repo, ShouldEqual, "repo")
				So(res.Issue.Number, ShouldEqual, 1)
				So(res.Issue.Status, ShouldEqual, "todo")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an invalid workflow.move message", t, func() {
		w := WFMove{}
		w.Subscribe(nc)
		body := []byte("")
		response, err := nc.Request("workflow.move", body, 1000*time.Millisecond)
		So(err, ShouldBeNil)
		So(string(response.Data), ShouldEqual, `{"error":"unexpected end of JSON input"}`)
	})

}
