deps:
	go get -u github.com/adriacidre/fsm
	go get -u github.com/nats-io/nats
	go get -u gopkg.in/redis.v3""
dev-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/smartystreets/goconvey/convey
	go get -u github.com/supu-io/messages
build:
	go build .
test:
	go test
lint:
	golint
cover:
	go test -cover
