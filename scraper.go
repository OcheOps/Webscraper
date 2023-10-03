package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Set the URL to scrape
	url := "http://quotes.toscrape.com"
	c.Visit(url)
//Url can change  will update as a cli in V2

	// Define a callback to be executed when a quote is found
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		// Extract the quote text
		quote := e.ChildText("span.text")
		// Extract the author
		author := e.ChildText("span small.author")
		// Extract the tags
		tags := e.ChildTexts("div.tags a.tag")

		// Print the scraped data
		fmt.Printf("Quote: %s\n", quote)
		fmt.Printf("Author: %s\n", author)
		fmt.Printf("Tags: %s\n\n", strings.Join(tags, ", "))
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
