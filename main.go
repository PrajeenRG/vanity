package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var config Config
	file, err := os.Open("vanity.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	log.Printf("Config: %+v", config)
}
