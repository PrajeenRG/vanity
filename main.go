package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"

	"go.prajeen.com/vanity/config"
	"go.prajeen.com/vanity/template"
)

func main() {
	var config config.Config
	file, err := os.Open("vanity.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var content bytes.Buffer
	err = template.Home(config).Render(context.Background(), &content)
	if err != nil {
		log.Printf("Error rendering home: %v", err)
	}
	log.Printf("Rendered home:\n%s\n", &content)

	for _, v := range config.Modules {
		err = template.Module(config.Domain, v).Render(context.Background(), &content)
		if err != nil {
			log.Printf("Error rendering module: %v", err)
		}
		log.Printf("Rendered for %s:\n%s\n", v.Name, &content)
	}
}
