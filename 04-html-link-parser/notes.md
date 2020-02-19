# HTML Link Parser
https://github.com/gophercises/link

Write a package that accepts HTML in some format (io.Reader or String) and convert HTML into some codified structure. We will be using this later in an upcoming exercise to build a sitemap builder. 

Links will be nested in different HTML elements, and it is very possible that you will have to deal with HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```
In situations like these we want to get output that looks roughly like:
```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

### Random
 - Practice parsing HTML documents and seeing how to use net/html package
 - Get something up and running first. Use  x/net/html pkg.
   - x/net/html are by the go team but not part of the stdlib. sometimes they get merged into the stdlib 
   - done this way so it can have a bit more flexibility

### Log


#### Startup
 - using html.Parse(r)
 - ignore nested links
 - make sure you just get something up and running first
 - using NodeType constants inssid the net/html package to help structure the data

#### HTML Docs are Trees
 - HTML file parsing using tree traversal..  
 - HTML Files are really just trees:
 - <body> is a root node and it has children nodes (like a <h1> node and a <div> node). Some children nodes may have their own children (like the <div> node having two separate <a> nodes as children)
 - the net/html package is going to be returning things called Nodes. It might have text in it, so it will be a TextNode, 
 -    

#### Defining the API
 - We ideally want to have a reader `r io.Reader`, then say give me links, err of `link.Parse(r)`
 - The reason we want to have an io.Reader is because there are all sorts of cases in which we have a reader: IE opening a file in filesystem, making a web request to a website. It makes sense to have use reader(over a string IE) becuase most of the time that's what we will be workign with. 
 - Write a parse function `Parse()` - this will take in an HTML document and return a slice of links parsed from the HTML.
 - After defining an empty func, write some example code (test, whatever) that uses it. In this case we create `examples/ex1/main.go`  and use the function Parse()
 - using `strings.NewReader(exampleHtml)` to create the reader for the example hardcoded HTML , then call `link.Parse()`
 

#### DFSing the HTML
 - https://godoc.org/golang.org/x/net/html
 - We're using trees and recursive searching here (Depth-First Search)
How do types work with nodes?
     - TextNode: raw text
     - DocumentNode: Wrapper node for everything. Wraps even <html> tag. It's the parent root node.
     - ElementNode: Any element like <div>, <a> <h1> <p> etc
     - CommentNode: HTML Comment
     - DoctypeNode: <!DOCTYPE html> type of shit
 - When you're looking at each node you can check the node type (`if n.Type == html.ElementNode && n.Data == "a" {}`) 

First we gotta write the parsing function
```go
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}
```

Next we write the depth-first-search func for recursively parsing through the HTML tree. This just prints out what it sees right now. We will extend it later. 

```go
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
```