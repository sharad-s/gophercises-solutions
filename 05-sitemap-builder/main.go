package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	link "gophercises/04-html-link-parser"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Has to be split this way because we want <url> <loc /> </url> for each page
type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

const xmlns string = "http://www.sitemaps.org/schemas/sitemap/0.9"

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	depth := flag.Int("depth", 3, "the maximum number of links deep to traverse")

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
	pages := bfs(*urlFlag, *depth)
	// for _, href := range pages {
	// 	// fmt.Println(href)
	// }

	// XML parse
	toXML := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}
	fmt.Println()
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

func bfs(urlStr string, maxDepth int) []string {
	// C
	seen := make(map[string]struct{})

	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})

		// Range over current queue, extract key into `url`
		for url := range q {
			// Check `seen` map for url, if seen then skip
			if _, ok := seen[url]; ok {
				continue
			}
			// Not seen, mark it as seen.
			seen[url] = struct{}{}
			// Get next links for this url
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	// Create return array of url strings
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

// Returns a func(string) bool and not just a "bool" because the func needs to be run elsewhere (??)
func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
