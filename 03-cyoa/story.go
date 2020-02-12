package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Story map[string]Chapter
type Story map[string]Chapter

// Chapter   .
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option   .
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// JSONStory  Any file we open is gonna be a reader, so pass it into Decoder
func JSONStory(r io.Reader) (Story, error) {
	// Decode the file
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Json equivalent
// { arc: Chapter }

// Where..
// Chapter = {
// 	title: "",
// 	story: [{}, {}],
// 	options: {
// 		text: "",
// 		arc: ""
// 	}
//  }

// Story[arc] = Chapter

var defaultHandlerTmpl = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Choose your own adventure</title>
  </head>
  <body>
    {{.Title}} {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>
`

// NewHandler :returns http.Handler interface
// Why not return type `handler`?
//  If you were to create a godoc for type `handler` it would not return any methods since it isinternal
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get Path
	// Remove leading and trailing whitespace
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// Get all the chars of the path string starting at index 1 (removes leading /)
	path = path[1:]

	// Access the Story map to see if the path returns a chapter
	// Use comma, ok idiom on map lookup
	if chapter, ok := h.s[path]; ok {
		// Create template
		// template.Must means template must complie correctly
		err := tpl.Execute(w, chapter)
		if err != nil {
			// Log verbose error, show generic error
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)

}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template
