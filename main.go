package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
    	// Visit only domains:
    	colly.AllowedDomains("xbox.com", "plus.ubisoft.com", “ea.com”),
    )

    // On every a element which has href attribute call callback
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        // Print link
        fmt.Printf("Link found: %q -> %s\n", e.Text, link)
        // Visit link found on page
        // Only those links are visited which are in AllowedDomains
        c.Visit(e.Request.AbsoluteURL(link))
    })

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    // Start scraping
    c.Visit("https://www.xbox.com/en-US/xbox-game-pass/cloud-gaming")
    // c.Visit("https://plus.ubisoft.com/")
    // c.Visit(“https://www.ea.com/ea-play/games")            
}