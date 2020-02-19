package main

import (
	"fmt"
	link "gophercises/04-html-link-parser"
	"strings"
)

var exampleHtml = `
<html>
<body>
	<h1>Yo!</h1>
	<a href="/other-page"> A link to another page </a>
	<a href="/second-page"> A link to a second page </a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
