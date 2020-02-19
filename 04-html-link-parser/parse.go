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

	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Printf("%+v\n", node)
	}
	// dfs(doc, "")
	//1. Find <a> nodes in doc
	//2. for each link node...
	//2.a build a link
	//3. return Links
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {

	// Base case for recursion. We've got a link Node
	// return slice [*html.Node(n)]
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	//Declare up here because we know we're going to return an *html.Node array
	var ret []*html.Node

	// Do the depth first search
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// We want to append whatever link nodes are returned
		ret = append(ret, linkNodes(c)...)
	}

	return ret
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
