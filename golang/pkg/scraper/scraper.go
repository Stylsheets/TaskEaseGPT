package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

type Result struct {
	TextContent string
	Links       []Link
	Titles      map[string][]string
}

type Link struct {
	Name string
	URL  string
}

func Scrape(url string) (Result, error) {
	c := colly.NewCollector()

	result := Result{
		Titles: make(map[string][]string),
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		name := strings.TrimSpace(e.Text)
		result.Links = append(result.Links, Link{Name: name, URL: link})
	})

	c.OnHTML("h1, h2, h3", func(e *colly.HTMLElement) {
		tag := e.Name
		result.Titles[tag] = append(result.Titles[tag], e.Text)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		result.TextContent = strings.TrimSpace(e.Text)

		// text has parts like \n\n\n... or \t\n\t\n... clean duplicates and make them single \n
		// Compile the regular expression
		regex := regexp.MustCompile(`((\n+)|(\t+))`)
		result.TextContent = regex.ReplaceAllString(result.TextContent, "\n")
		regex = regexp.MustCompile(`\n+`)
		result.TextContent = regex.ReplaceAllString(result.TextContent, "\n")
		fmt.Println(result.TextContent)
	})

	err := c.Visit(url)
	if err != nil {
		return Result{}, err
	}

	return result, nil
}
