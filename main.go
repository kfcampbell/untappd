package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kfcampbell/untappd/untappd"
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

	f, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("day_of_week,date,brewery,city,state,beer_name,rating\n")
	if err != nil {
		return err
	}

	client := client.NewClient(cfg.Username, cfg.ClientID, cfg.ClientKey)
	beers, err := client.GetBeers()
	if err != nil {
		return err
	}

	nextURL := beers.Response.Pagination.NextURL
	beersListLength := len(beers.Response.BeersList.Items)
	for strings.Contains(nextURL, client.Username) {
		log.Printf("nextURL: %v, username: %v", nextURL, client.Username)
		for i := 0; i < beersListLength; i++ {
			log.Printf("i: %v, beersListLength: %v", i, beersListLength)
			checkin, err := client.GetCheckin(beers.Response.BeersList.Items[i].FirstCheckinID)
			if err != nil {
				return nil
			}
			line := constructCSVLine(checkin.Response.Checkins.Items[0], beers.Response.BeersList.Items[i].Rating)

			_, err = f.WriteString(line)
			if err != nil {
				return err
			}
			log.Printf(line)
			//time.Sleep(time.Second * 45)
		}

		beers, err = client.GetNextBeers(nextURL)
		if err != nil {
			return err
		}
		nextURL = beers.Response.Pagination.NextURL
		beersListLength = len(beers.Response.BeersList.Items)
	}

	return nil
}

func constructCSVLine(checkin untappd.CheckinItem, rating float32) string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v\n",
		checkin.Date, checkin.Brewery.Name, checkin.Venue.Location.City, checkin.Venue.Location.State, checkin.Beer.Name, rating)
}
