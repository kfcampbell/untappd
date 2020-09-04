package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kfcampbell/untappd/untappd/config"
)

const apiRoot = "https://api.untappd.com/v4/"

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

	beersURL := apiRoot + fmt.Sprintf("user/beers/%v", cfg.Username) + fmt.Sprintf("?client_id=%v", cfg.ClientID) + fmt.Sprintf("&client_secret=%v", cfg.ClientKey)

	res, err := http.Get(beersURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
