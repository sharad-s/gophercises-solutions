package cyoa

// map[string]Chapter

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
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
