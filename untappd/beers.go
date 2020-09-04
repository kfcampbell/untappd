package untappd

// Body is the overall response body
type Body struct {
	Response Response `json:"response"`
}

// Response is the only thing we care about inside the response body
type Response struct {
	TotalCount int        `json:"total_count"`
	BeersList  BeersList  `json:"beers"`
	Pagination Pagination `json:"pagination"`
}

// Pagination tells us where to look next
type Pagination struct {
	NextURL string `json:"next_url"`
	Offset  int    `json:"offset"`
}

// BeersList is the "beers" part
type BeersList struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}

// Item is basically a checkin without the location. It has a time, a beer, and a brewery
type Item struct {
	FirstCheckinID int     `json:"first_checkin_id"`
	Rating         float32 `json:"rating_score"`
	FirstDate      string  `json:"first_had"`
	Beer           Beer    `json:"beer"`
}

// Beer is what's inside the items list
type Beer struct {
	ID    int     `json:"bid"`
	Name  string  `json:"beer_name"`
	ABV   float32 `json:"beer_abv"`
	Style string  `json:"beer_style"`
}
