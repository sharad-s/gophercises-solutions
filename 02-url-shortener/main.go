package main

import (
	"flag"
	"fmt"
	"gophercises/02-url-shortener/urlshort"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Flags
	yamlFileName := flag.String("y", "paths.yaml", "a yaml in the format of '- path, url'")
	flag.Parse()

	// Read YAML file
	yamlFile, err := ioutil.ReadFile(*yamlFileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8081")
	http.ListenAndServe(":8081", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
