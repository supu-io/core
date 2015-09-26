package main

import (
	"encoding/json"
	"errors"
	"gopkg.in/redis.v3"
	"log"
	"os"
)

// Struct representing Redis client
type Storage struct {
	Addr     string
	Password string
	DB       int64
	Client   *redis.Client
}

type storageConfig struct {
	Storage Storage `json:"redis"`
}

// Setup the client from a config source string
func (s *Storage) setup(source string) {
	c := storageConfig{}
	file, err := os.Open(source)
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Println("Config file is invalid")
		log.Panic("error:", err)
	}

	s.Client = redis.NewClient(&redis.Options{
		Addr:     c.Storage.Addr,
		Password: c.Storage.Password,
		DB:       c.Storage.DB,
	})
}

// Builds a to be stored key from a given string
func (s *Storage) buildKey(str string) string {
	return str
}

// Set a value for a given key
func (s *Storage) set(key string, value string) error {
	if err := s.Client.Set(key, value, 0).Err(); err != nil {
		return errors.New("Data can't be stored")
	}
	return nil
}

// Get the value for a given key
func (s *Storage) get(key string) (string, error) {
	return s.Client.Get(key).Result()
}

// Gets an issue from database from its id
func (s *Storage) GetIssue(id string) *Issue {
	key := s.buildKey(id)
	str, err := s.get(key)
	if err != nil {
		return nil
	}

	i := Issue{}
	if err := json.Unmarshal([]byte(str), &i); err != nil {
		return nil
	}

	return &i
}

// Store an issue on the database
func (s *Storage) SetIssue(i *Issue) {
	key := s.buildKey(i.ID)
	value := i.toJSON()
	s.set(key, string(*value))
}