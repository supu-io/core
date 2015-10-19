package main

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAllStates(t *testing.T) {
	setup()
	Convey("Given a valid workflow.states.all message", t, func() {
		w := WFStatesAll{}
		w.Subscribe(nc)
		Convey("When send the message", func() {
			response, err := nc.Request("workflow.states.all", []byte(""), 1000*time.Millisecond)
			Convey("Then issue state should be todo", func() {
				res := []string{}
				json.Unmarshal(response.Data, &res)

				So(len(res), ShouldEqual, 6)
				So(err, ShouldBeNil)
			})
		})
	})
}
