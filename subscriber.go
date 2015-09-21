package main

import (
	"encoding/json"
	"github.com/nats-io/nats"
)

type Subscriber struct {
	Storage  *Storage
	Workflow *Workflow
	Input    *InputIssue
}

type InputIssue struct {
	Issue  string `json:"issue"`
	Status string `json:"status"`
}

func (this *Subscriber) subscribe() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	c.Subscribe("issue.update", func(m *nats.Msg) {
		i := this.manageIssue(string(m.Data))
		c.Publish(m.Reply, i.toJSON())
	})
}

func (this *Subscriber) manageIssue(body string) *Issue {
	input := InputIssue{}
	if err := json.Unmarshal([]byte(body), &input); err != nil {
		panic(err)
	}
	this.Input = &input
	i := this.Storage.GetIssue(this.Input.Issue)
	if i == nil {
		return nil
	}
	this.Workflow.transact(i, this.Input.Status)
	this.Storage.SetIssue(i)

	return i
}
