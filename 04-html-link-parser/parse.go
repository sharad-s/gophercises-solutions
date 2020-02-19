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
	//1. Find all <a> nodes in doc
	nodes := linkNodes(doc)
	//2. For each link node...
	var links []Link
	for _, node := range nodes {
		// fmt.Printf("%+v\n", node)
		//2a. Build a link
		links = append(links, buildLink(node))
	}

	// Practice DFS on document
	// dfs(doc, "")

	//3. Return Links
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	// Get Href property from Node
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = "TODO... Parse the Text"
	return ret
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
