package main

import (
	"log"
	"path/filepath"

	"github.com/kfcampbell/untappd/untappd/client"
	"github.com/kfcampbell/untappd/untappd/config"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

func realMain() error {
	log.Printf("loading configuration...")
	configFile, err := filepath.Abs("config/config.json")
	if err != nil {
		return err
	}

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return err
	}

	client := client.NewClient(cfg.Username, cfg.ClientID, cfg.ClientKey)
	beers, err := client.GetBeers()
	if err != nil {
		return err
	}

	log.Printf(beers.Response.Pagination.NextURL)

	return nil
}
