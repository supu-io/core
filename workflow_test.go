package main

import (
	"testing"

	"github.com/adriacidre/fsm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidTransitions(t *testing.T) {
	t.Parallel()
	Convey("Given an issue on created status", t, func() {
		issue := Issue{State: "created"}
		Convey("When I apply an todo event", func() {
			e := fsm.State("todo")
			err := issue.Apply("todo").Transition(e)
			Convey("Then issue state should be todo", func() {
				So(issue.State, ShouldEqual, "todo")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on doing status", t, func() {
		issue := Issue{State: "doing"}
		Convey("When I apply an review event", func() {
			e := fsm.State("review")
			err := issue.Apply("review").Transition(e)
			Convey("Then issue state should be review", func() {
				So(issue.State, ShouldEqual, "review")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on doing status", t, func() {
		issue := Issue{State: "doing"}
		Convey("When I apply an todo event", func() {
			e := fsm.State("todo")
			err := issue.Apply("todo").Transition(e)
			Convey("Then issue state should be todo", func() {
				So(issue.State, ShouldEqual, "todo")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on review status", t, func() {
		issue := Issue{State: "review"}
		Convey("When I apply an doing event", func() {
			e := fsm.State("doing")
			err := issue.Apply("doing").Transition(e)
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "doing")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on review status", t, func() {
		issue := Issue{State: "review"}
		Convey("When I apply an uat event", func() {
			e := fsm.State("uat")
			err := issue.Apply("uat").Transition(e)
			Convey("Then issue state should be uat", func() {
				So(issue.State, ShouldEqual, "uat")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on uat status", t, func() {
		issue := Issue{State: "uat"}
		Convey("When I apply an doing event", func() {
			e := fsm.State("doing")
			err := issue.Apply("doing").Transition(e)
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "doing")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on uat status", t, func() {
		issue := Issue{State: "uat"}
		Convey("When I apply an done event", func() {
			e := fsm.State("done")
			err := issue.Apply("done").Transition(e)
			Convey("Then issue state should be done", func() {
				So(issue.State, ShouldEqual, "done")
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestInValidTransitions(t *testing.T) {
	t.Parallel()
	Convey("Given an issue on created status", t, func() {
		issue := Issue{State: "created"}
		Convey("When I apply an doing event", func() {
			e := fsm.State("uat_ok")
			err := issue.Apply("uat_ok").Transition(e)
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "created")
				So(err.Error(), ShouldEqual, "invalid transition")
			})
		})
	})
}
