package cyoa

import (
	"encoding/json"
	"io"
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
