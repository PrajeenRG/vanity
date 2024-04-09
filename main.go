package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"path"

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

	publicDir := path.Join(opts.OutputDir, "public")
	err = os.Mkdir(publicDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating public directory: %v", err)
	}

	var content bytes.Buffer
	err = template.Home(config).Render(context.Background(), &content)
	if err != nil {
		log.Printf("Error rendering home: %v", err)
	}
	err = os.WriteFile(path.Join(publicDir, "index.html"), content.Bytes(), os.ModePerm)
	if err != nil {
		log.Printf("Error writing html for home: %v", err)
	}

	info := vconf.ProcessConfig(config)
	for _, v := range info {
		err := os.Mkdir(path.Join(publicDir, v.Name), os.ModePerm)
		if err != nil {
			log.Printf("Error creating directory for %s: %v", v.Name, err)
			continue
		}

		content.Reset()
		err = template.Module(v).Render(context.Background(), &content)
		if err != nil {
			log.Printf("Error rendering module: %v", err)
		}
		err = os.WriteFile(path.Join(publicDir, v.Name, "index.html"), content.Bytes(), os.ModePerm)
		if err != nil {
			log.Printf("Error writing html for module %s: %v", v.Name, err)
		}
	}
}
