package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var s Subscriber

func setup() {
	s = Subscriber{}
}

func TestSubscriber(t *testing.T) {
	setup()
	Convey("Given an existing issue", t, func() {
		Convey("When I try to manage the existing issue", func() {
			stored := s.manageIssue("{\"issue\":{\"id\":\"foo\",\"status\":\"created\"}, \"status\":\"todo\"}")
			Convey("Then issue status get updated", func() {
				So(stored.State, ShouldEqual, "todo")
			})
		})
	})
}

func TestSubscribeNonExistingIssue(t *testing.T) {
	setup()
	Convey("Given an existing issue", t, func() {
		Convey("When I try to manage the existing issue", func() {
			stored := s.manageIssue("{\"issue\":{\"id\":\"bar\",\"status\":\"created\"}, \"status\":\"in_progress\"}")
			Convey("Then issue status get updated", func() {
				So(stored, ShouldEqual, nil)
			})
		})
	})
}
