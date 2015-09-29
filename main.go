package main

func main() {
	workflow := Workflow{}
	storage := Storage{}
	storage.setup("config.json")

	s := Subscriber{Storage: &storage, Workflow: &workflow}
	s.subscribe()
}
