package responses

type HostListItem struct {
	UID       string  `json:"uid"`
	Title     string  `json:"title"`
	Rate      float64 `json:"rate"`
	RateCount int     `json:"rate_count"`
	Priority  int     `json:"priority"`
}
