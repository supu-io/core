package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidTransitions(t *testing.T) {
	t.Parallel()
	w := Workflow{}
	Convey("Given an issue on created status", t, func() {
		issue := Issue{State: "created"}
		Convey("When I apply an todo event", func() {
			err := w.transact(&issue, "todo")
			Convey("Then issue state should be todo", func() {
				So(issue.State, ShouldEqual, "todo")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on doing status", t, func() {
		issue := Issue{State: "doing"}
		Convey("When I apply an ci event", func() {
			err := w.transact(&issue, "ci")
			Convey("Then issue state should be ci", func() {
				So(issue.State, ShouldEqual, "ci")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on ci status", t, func() {
		issue := Issue{State: "ci"}
		Convey("When I apply an review event", func() {
			err := w.transact(&issue, "review")
			Convey("Then issue state should be review", func() {
				So(issue.State, ShouldEqual, "review")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on ci status", t, func() {
		issue := Issue{State: "ci"}
		Convey("When I apply an doing event", func() {
			err := w.transact(&issue, "doing")
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "doing")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on review status", t, func() {
		issue := Issue{State: "review"}
		Convey("When I apply an doing event", func() {
			err := w.transact(&issue, "doing")
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "doing")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on review status", t, func() {
		issue := Issue{State: "review"}
		Convey("When I apply an uat event", func() {
			err := w.transact(&issue, "uat")
			Convey("Then issue state should be uat", func() {
				So(issue.State, ShouldEqual, "uat")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on uat status", t, func() {
		issue := Issue{State: "uat"}
		Convey("When I apply an doing event", func() {
			err := w.transact(&issue, "doing")
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "doing")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given an issue on uat status", t, func() {
		issue := Issue{State: "uat"}
		Convey("When I apply an done event", func() {
			err := w.transact(&issue, "done")
			Convey("Then issue state should be done", func() {
				So(issue.State, ShouldEqual, "done")
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestInValidTransitions(t *testing.T) {
	t.Parallel()
	w := Workflow{}
	Convey("Given an issue on created status", t, func() {
		issue := Issue{State: "created"}
		Convey("When I apply an doing event", func() {
			err := w.transact(&issue, "uat_ok")
			Convey("Then issue state should be doing", func() {
				So(issue.State, ShouldEqual, "created")
				So(err.Error(), ShouldEqual, "invalid transition")
			})
		})
	})
}
