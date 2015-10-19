package main

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAvailableStates(t *testing.T) {
	setup()
	Convey("Given a valid workflow.states.all message", t, func() {
		w := WFStatesAvailable{}
		w.Subscribe(nc)
		Convey("When send the message", func() {
			msg := WFStatesAvailableInput{
				Issue: &Issue{State: "doing"},
			}
			body, err := json.Marshal(&msg)
			response, err := nc.Request("workflow.states.available", body, 1000*time.Millisecond)
			Convey("Then 2 possible status should be received", func() {
				res := []string{}
				json.Unmarshal(response.Data, &res)

				So(len(res), ShouldEqual, 2)
				So(err, ShouldBeNil)
			})
		})
	})
}
