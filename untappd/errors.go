package untappd

// CheckinsBodyBadResponse is to allow us to compensate for untappd's inconsistent APIs.
type CheckinsBodyBadResponse struct {
	Response CheckinsResponseBadResponse `json:"response"`
}

// CheckinsResponseBadResponse holds the response from a checkins request (inside a body)
type CheckinsResponseBadResponse struct {
	Checkins CheckinsBadResponse `json:"checkins"`
}

// CheckinsBadResponse holds the list of actual checkins
type CheckinsBadResponse struct {
	Count int                      `json:"count"`
	Items []CheckinItemBadResponse `json:"items"`
}

// CheckinItemBadResponse is an individual checkin
type CheckinItemBadResponse struct {
	ID      int     `json:"checkin_id"`
	Date    string  `json:"created_at"`
	Beer    Beer    `json:"beer"`
	Venue   []Venue `json:"venue,omitempty"`
	Brewery Brewery `json:"brewery"`
}

// AdaptBadResponseCheckin is to convert from the bad response when venue is coming back as an array to a good response
// with placeholder venue values.
func AdaptBadResponseCheckin(badResponse CheckinsBodyBadResponse) *CheckinsBody {
	res := &CheckinsBody{}
	placeholderLocation := &Location{
		Address: "placeholder_address",
		City:    "placeholder_city",
		State:   "placeholder_state",
		Country: "placeholder_country",
	}
	placeholderVenue := &Venue{
		ID:       -1,
		Name:     "placeholder_venue",
		Location: *placeholderLocation,
	}
	checkin := &CheckinItem{
		ID:      badResponse.Response.Checkins.Items[0].ID,
		Date:    badResponse.Response.Checkins.Items[0].Date,
		Beer:    badResponse.Response.Checkins.Items[0].Beer,
		Brewery: badResponse.Response.Checkins.Items[0].Brewery,
		Venue:   *placeholderVenue,
	}
	res.Response.Checkins.Count = 1
	res.Response.Checkins.Items = append(res.Response.Checkins.Items, *checkin)

	return res
}
