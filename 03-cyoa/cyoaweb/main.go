package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/03-cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3030, "the port to start the CYOA web app on")
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	// Try Open file
	f, err := os.Open(*fileName)
	if err != nil {
		// Panic not usually recommmended
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	//get handler
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)

	// Start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
