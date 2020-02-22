package main

import (
	"flag"
	"io"
	"net/http"
	"os"
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

	resp, err := http.Get(*urlFlag)
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
