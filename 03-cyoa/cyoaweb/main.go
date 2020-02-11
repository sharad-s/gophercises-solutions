package main

import (
	"encoding/json"
	"flag"
	"fmt"
	cyoa "gophercises/03-cyoa"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	// Try Open file
	f, err := os.Open(*fileName)
	if err != nil {
		// Panic not usually recommmended
		panic(err)
	}

	// Decode the file
	// Any file we open is gonna be a reader, so pass it into Decoder
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
