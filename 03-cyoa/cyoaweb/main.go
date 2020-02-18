package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/03-cyoa"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

	// Create a new test template
	tpl := template.Must(template.New("").Parse(testStoryTmpl))

	// Create a new path func
	// tpl := template.Must(template.New("").Parse("Hello!"))

	//get handler, add functional options
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(testPathFn),
	)

	// create new mux (mux shows a default 404 page unless it's prefixed with /story/)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	// Start server
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func testPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" || path == "/" || path == "" {
		path = "/story/intro"
	}
	// Get all the chars of the path string starting at index 1 (removes leading /)
	path = path[len("/story/"):]
	return path
}

var testStoryTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
