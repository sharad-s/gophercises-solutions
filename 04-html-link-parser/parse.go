package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...") in an HTML
// document
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

// Name the var n to represent Node
func dfs(n *html.Node, padding string) {
	msg := n.Data

	// if n.Data is an HTML Element, show it as a tag
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}

	// Print the thing
	fmt.Println(padding, msg)

	// Recursively call each this node's children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
