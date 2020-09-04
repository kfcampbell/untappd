package client

import (
	"fmt"
	"net/http"

	"github.com/golang/go/src/encoding/json"
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
func (c *Client) GetBeers() (*untappd.BeersBody, error) {
	url := apiRoot + c.getBeersURLPath() + c.getAuthString() + "&limit=50"
	return getBeers(url)
}

// GetNextBeers gets the next beers in an offset list
func (c *Client) GetNextBeers(nextURL string) (*untappd.BeersBody, error) {
	url := nextURL + c.getAuthString()
	return getBeers(url)
}

func getBeers(url string) (*untappd.BeersBody, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result untappd.BeersBody
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCheckin returns a single checkin when given an ID
func (c *Client) GetCheckin(ID int) (*untappd.CheckinsBody, error) {
	url := apiRoot + c.getCheckinURLPath() + c.getAuthString() + fmt.Sprintf("&limit=1&max_id=%v", ID)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result untappd.CheckinsBody
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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
