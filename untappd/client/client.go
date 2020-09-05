package client

import (
	"fmt"
	"net/http"

	"github.com/golang/go/src/encoding/json"
	"github.com/golang/go/src/log"
	"github.com/kfcampbell/untappd/untappd"
)

const apiRoot = "https://api.untappd.com/v4/"

// Client is an HTTP client for talking to the Untappd API
type Client struct {
	Username  string
	ClientID  string
	ClientKey string
}

// NewClient creates and returns a new HTTP Client
func NewClient(username string, clientID string, clientKey string) *Client {
	return &Client{
		Username:  username,
		ClientID:  clientID,
		ClientKey: clientKey,
	}
}

// GetBeers returns the beers that Xavier has tried
func (c *Client) GetBeers() (untappd.BeersBody, error) {
	url := apiRoot + c.getBeersURLPath() + c.getAuthString() + "&limit=50"
	return getBeers(url)
}

// GetNextBeers gets the next beers in an offset list
func (c *Client) GetNextBeers(nextURL string) (untappd.BeersBody, error) {
	url := nextURL + c.getAuthString()
	return getBeers(url)
}

func getBeers(url string) (untappd.BeersBody, error) {
	log.Printf("Getting beers with url: %v", url)
	beers := &untappd.BeersBody{}
	res, err := http.Get(url)
	if err != nil {
		return *beers, err
	}

	err = json.NewDecoder(res.Body).Decode(&beers)
	defer res.Body.Close()
	if err != nil {
		return *beers, err
	}

	log.Printf("Got beers! %v items in this request of %v total", beers.Response.BeersList.Count, beers.Response.TotalCount)
	return *beers, nil
}

// GetCheckin returns a single checkin when given an ID
func (c *Client) GetCheckin(ID int) (untappd.CheckinsBody, error) {
	// failing getting checkin for this ID: 928657480
	// works via API somehow
	// the bug might be related to the res.body.close() calls? those return an error but we're not "return"ing it
	log.Printf("Getting checkin for ID: %v", ID)
	url := apiRoot + c.getCheckinURLPath() + c.getAuthString() + fmt.Sprintf("&limit=1&max_id=%v", ID)

	result := &untappd.CheckinsBody{}
	res, err := http.Get(url)
	if err != nil {
		return *result, err
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	defer res.Body.Close()
	if err != nil {
		return *result, err
	}

	log.Printf("Got checkin for ID: %v, with beer: %v", ID, result.Response.Checkins.Items[0].Beer.Name)
	return *result, nil
}

func (c *Client) getAuthString() string {
	return fmt.Sprintf("?client_id=%v", c.ClientID) + fmt.Sprintf("&client_secret=%v", c.ClientKey)
}

func (c *Client) getBeersURLPath() string {
	return fmt.Sprintf("user/beers/%v", c.Username)
}

func (c *Client) getCheckinURLPath() string {
	return fmt.Sprintf("user/checkins/%v", c.Username)
}
