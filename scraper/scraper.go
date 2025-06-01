package scraper

import (
	"fmt"
	"net/url"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

type Scraper struct {
	Url     url.URL
	Timeout int
}

func New(url url.URL, timeout int) *Scraper {
	return &Scraper{Url: url, Timeout: timeout}
}

func (scraper *Scraper) Scrape() []ShowInfo {
	shows := []ShowInfo{}
	c := colly.NewCollector()
	c.SetRequestTimeout(time.Duration(scraper.Timeout) * time.Second)

	c.OnRequest(handleRequest)
	c.OnResponse(handleResponse)
	c.OnError(handleError)
	c.OnScraped(handleScraped)

	c.OnHTML(`p[style^="font-size"]`, func(showMarkup *colly.HTMLElement) {
		rgx := regexp.MustCompile(`\(([^)]+)\)`)

		showDate := showMarkup.DOM.Find("b").First()
		showVenue := showMarkup.ChildText("span")
		showTimes := rgx.FindString(showMarkup.Text)

		show := ShowInfo{Date: showDate.Text(), Times: showTimes, Venue: showVenue}

		shows = append(shows, show)
	})

	c.Visit(scraper.Url.String())

	return shows
}

func handleRequest(r *colly.Request) {
	fmt.Println("Visiting", r.URL)
}

func handleResponse(r *colly.Response) {
	fmt.Println("Got a response from in callback", r.Request.URL)
}

func handleError(r *colly.Response, e error) {
	fmt.Println("Got this error:", e)
}

func handleScraped(r *colly.Response) {
	fmt.Println("Finished scraping", r.Request.URL)
}
