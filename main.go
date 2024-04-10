package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path"

	vconf "go.prajeen.com/vanity/config"
	"go.prajeen.com/vanity/template"
)

func main() {
	opts := vconf.ParseCLIOpts()

	var config vconf.Config
	file, err := os.Open(opts.ConfigFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	publicDir := path.Join(opts.OutputDir, "public")
	err = os.Mkdir(publicDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating public directory: %v", err)
	}

	f, err := os.Create(path.Join(publicDir, "index.html"))
	if err != nil {
		log.Printf("Error creating index.html for home: %v", err)
	}
	err = template.Home(config).Render(context.Background(), f)
	if err != nil {
		log.Printf("Error rendering home: %v", err)
	}
	log.Printf("Created index.html for home page")

	info := vconf.ProcessConfig(config)
	for _, v := range info {
		err := os.Mkdir(path.Join(publicDir, v.Name), os.ModePerm)
		if err != nil {
			log.Printf("Error creating directory for %s: %v", v.Name, err)
			continue
		}

		f, err = os.Create(path.Join(publicDir, v.Name, "index.html"))
		if err != nil {
			log.Printf("Error creating index.html for module %s: %v", v.Name, err)
		}
		err = template.Module(v).Render(context.Background(), f)
		if err != nil {
			log.Printf("Error rendering module: %v", err)
		}
		log.Printf("Created index.html for module %s", v.Name)
	}
}
