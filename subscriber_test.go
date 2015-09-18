package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var s Subscriber

func setup() {
	workflow := Workflow{}
	storage := Storage{}
	storage.setup("config.json")

	s = Subscriber{Workflow: &workflow, Storage: &storage}
}

func TestSubscriber(t *testing.T) {
	setup()
	Convey("Given an existing issue", t, func() {
		i := Issue{ID: "foo", Name: "bar", State: "created"}
		s.Storage.SetIssue(&i)
		Convey("When I try to manage the existing issue", func() {
			s.manageIssue("{\"issue\":\"foo\",\"status\":\"in_progress\"}")
			Convey("Then issue status get updated", func() {
				stored := s.Storage.GetIssue(i.ID)
				So(stored.State, ShouldEqual, "in_progress")
			})
		})
	})
}

func TestSubscribeNonExistingIssue(t *testing.T) {
	setup()
	Convey("Given an existing issue", t, func() {
		Convey("When I try to manage the existing issue", func() {
			s.manageIssue("{\"issue\":\"bar\",\"status\":\"in_progress\"}")
			Convey("Then issue status get updated", func() {
				stored := s.Storage.GetIssue("bar")
				So(stored, ShouldEqual, nil)
			})
		})
	})
}
