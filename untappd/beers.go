package untappd

// BeersBody is the overall response body for a beers request
type BeersBody struct {
	Response BeersResponse `json:"response"`
}

// CheckinsBody is the overall response body for a checkins request
type CheckinsBody struct {
	Response CheckinsResponse `json:"response"`
}

// BeersResponse is the only thing we care about inside the response body
type BeersResponse struct {
	TotalCount int        `json:"total_count"`
	BeersList  BeersList  `json:"beers"`
	Pagination Pagination `json:"pagination"`
}

// CheckinsResponse holds the response from a checkins request (inside a body)
type CheckinsResponse struct {
	Checkins Checkins `json:"checkins"`
}

// Checkins holds the list of actual checkins
type Checkins struct {
	Count int           `json:"count"`
	Items []CheckinItem `json:"items"`
}

// Pagination tells us where to look next
type Pagination struct {
	NextURL string `json:"next_url"`
	Offset  int    `json:"offset"`
}

// BeersList is the "beers" part
type BeersList struct {
	Count int        `json:"count"`
	Items []BeerItem `json:"items"`
}

// BeerItem is basically a checkin without the location. It has a time, a beer, and a brewery
type BeerItem struct {
	FirstCheckinID int     `json:"first_checkin_id"`
	Rating         float32 `json:"rating_score"`
	FirstDate      string  `json:"first_had"`
	Beer           Beer    `json:"beer"`
}

// CheckinItem is an individual checkin
type CheckinItem struct {
	ID      int     `json:"checkin_id"`
	Date    string  `json:"created_at"`
	Beer    Beer    `json:"beer"`
	Venue   Venue   `json:"venue"`
	Brewery Brewery `json:"brewery"`
}

// Venue represents the location at which a checkin occurred
type Venue struct {
	ID       int      `json:"venue_id"`
	Name     string   `json:"venue_name"`
	Location Location `json:"location"`
}

// Location represents the location a venue belongs to
type Location struct {
	Address string `json:"venue_address"`
	City    string `json:"venue_city"`
	State   string `json:"venue_state"`
	Country string `json:"venue_country"`
}

// Beer is what's inside the items list
type Beer struct {
	ID    int     `json:"bid"`
	Name  string  `json:"beer_name"`
	ABV   float32 `json:"beer_abv"`
	Style string  `json:"beer_style"`
}

// Brewery represents the brewery a beer is made by
type Brewery struct {
	ID   int    `json:"brewery_id"`
	Name string `json:"brewery_name"`
	// Location
}
