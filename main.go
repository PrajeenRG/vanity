package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"

	"go.prajeen.com/vanity/config"
	vconf "go.prajeen.com/vanity/config"
	"go.prajeen.com/vanity/template"
)

func main() {
	opts := config.ParseCLIOpts()

	var config vconf.Config
	file, err := os.Open(opts.ConfigFile)
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

	info := vconf.ProcessConfig(config)
	for _, v := range info {
		err = template.Module(v).Render(context.Background(), &content)
		if err != nil {
			log.Printf("Error rendering module: %v", err)
		}
		log.Printf("Rendered for %s:\n%s\n", v.Name, &content)
	}
}
