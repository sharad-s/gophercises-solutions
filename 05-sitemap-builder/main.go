package main

import (
	"flag"
	"fmt"
	link "gophercises/04-html-link-parser"
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

	// Make GET request to URL
	resp, err := http.Get(*urlFlag)
	if err != nil {
		// handle error
		panic(err)
	}
	// Close Response at the end
	defer resp.Body.Close()

	// Output resp.Body to STDOUT
	// io.Copy(os.Stdout, resp.Body)

	// Use link-parser to extract links from HTML and print each out
	links, _ := link.Parse(resp.Body)
	// for _, l := range links {
	// 	fmt.Println(l)
	// }

	// Create Base URL from Request
	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	fmt.Println(base)

	// Loop through links on page, transform links
	// Append to hrefs slice
	var hrefs []string
	for _, l := range links {
		switch {
		// link starts with /, add baseURL
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		// link starts with http:
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	// Print out hrefs slice
	for _, href := range hrefs {
		fmt.Println(href)
	}

}
