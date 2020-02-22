package main

import (
	"flag"
	"fmt"
	link "gophercises/04-html-link-parser"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()
	// fmt.Println(*urlFlag)

	/*
		1. GET the webpage
		2. Parse all the links on the page
		3. Build proper URLs with our links
		4. Filter out any links with a different domain
		5. Find all the pages (BFS)
		6. Print out XML
	*/
	pages := get(*urlFlag)
	for _, href := range pages {
		fmt.Println(href)
	}
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		// If withPrefix(base)(link) is true
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

// Returns a func(string) bool and not just a "bool" because the func needs to be run elsewhere (??)
func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func get(urlStr string) []string {
	// Make GET request to URL
	resp, err := http.Get(urlStr)
	if err != nil {
		// handle error
		panic(err)
	}
	// Close Response at the end
	defer resp.Body.Close()

	// Create Base URL from Request
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	// fmt.Println(base)

	// Create and return hrefs slice
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(body io.Reader, base string) []string {
	var ret []string

	// Use link-parser to extract links from HTML and print each out
	links, _ := link.Parse(body)

	// Loop through links on page, transform links
	// Append to hrefs slice
	for _, l := range links {
		switch {
		// link starts with /, add baseURL
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		// link starts with http:
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}
