package main

import (
	"encoding/json"
	"log"
)

// ToJSON represents an object as a json []byte
func ToJSON(i interface{}) []byte {
	json, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}

	return json
}
