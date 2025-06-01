package scraper

import "fmt"

type ShowInfo struct {
	Date  string
	Venue string
	Times string
}

func (show *ShowInfo) String() string {
	return fmt.Sprintf("The Water Spots on %s %s at %s", show.Date, show.Times, show.Venue)
}
