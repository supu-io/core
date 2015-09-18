package main

func main() {
	workflow := Workflow{}
	storage := Storage{}
	storage.setup("config.yml")

	s := Subscriber{Storage: &storage, Workflow: &workflow}
	s.subscribe()
}
