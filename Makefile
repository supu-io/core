deps:
	go get -u github.com/adriacidre/fsm
	go get -u github.com/nats-io/nats
	go get -u gopkg.in/redis.v3""
build:
	go build .
