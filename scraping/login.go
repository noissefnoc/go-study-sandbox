package main

import (
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	err := c.Post("https://netflix.com/jp/login", map[string]string{"userLoginId": "takkubkiba@gmail.com", "password": "f4msm9rx782"})

	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// start scraping
	err = c.Visit("https://www.netflix.com/viewingactivity")

	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("body", func(e *colly.HTMLElement) {
		log.Print(e.DOM.Children().Text())
	})
}
