deps:
	go get -u github.com/adriacidre/fsm
	go get -u github.com/nats-io/nats
	go get -u gopkg.in/redis.v3""
dev-deps:
	go get -u github.com/smartystreets/goconvey/convey
build:
	go build .
test:
	cp config.json.tpl config.json
	go test
