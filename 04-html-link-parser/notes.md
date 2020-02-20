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

## Random
 - Practice parsing HTML documents and seeing how to use net/html package
 - Get something up and running first. Use  x/net/html pkg.
   - x/net/html are by the go team but not part of the stdlib. sometimes they get merged into the stdlib 
   - done this way so it can have a bit more flexibility
 - I still don't understand WHEN to use a pointer when declaring vars/params/structs/return types... I understand WHY, but I don't understand WHEN. By this i mean dereferencing a var with `*element` ??
 - `[]...` in Go is the way to spread a slice. Similar to spread operator in JS `...[]`
 - Initialize your return var up top if you already know what you're returning and it's type
 - Why can you create a slice var in a recursive loop with `var ret []*html.Node` and append to it each recursive function call? Doesn't that statement re-intialize it every iteration ?? 

## Questions
 - Why can you create a slice var in a recursive loop with `var ret []*html.Node` and append to it each recursive function call? Doesn't that statement re-intialize it every iteration?  (see `linkNodes()` func)
 - Rephrased above: Why can you initialize an array with `var` every recursive function call but still keep any values appended to it?. 

A: I think it's because that `ret` var is actually appended to recursively from the bottom all the way back to the top of the call stack. Meaning, each iteration, the `ret` only cares about what it gets from the next level down. Traversing that call stack upwards, that `ret` will contain a spread out list of what it recieved from the level below, until at the very top you have the full list. 

 - I still don't understand WHEN to use a pointer when declaring vars/params/structs/return types... I understand WHY, but I don't understand WHEN. By this I mean dereferencing or passing in a param with `*element` 
 - In `buildLink()`, We range over the attributes of each Node using `range n.Attr {}` because n.Attr is a slice, not a key/value map. Why is n.Attr a slice? I only get one element anyway.

### Log


#### Startup
 - using html.Parse(r)
 - ignore nested links
 - make sure you just get something up and running first
 - using NodeType constants inssid the net/html package to help structure the data

#### HTML Docs are Trees
 - HTML file parsing using tree traversal..  
 - HTML Files are really just trees:
 - `<body>` is a root node and it has children nodes (like a `<h1>` node and a `<div>` node). Some children nodes may have their own children (like the `<div>` node having two separate `<a>` nodes as children)
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
     - DocumentNode: Wrapper node for everything. Wraps even `<html>` tag. It's the parent root node.
     - ElementNode: Any element like ` <html> <div> <a> <h1> <p>` etc...
     - CommentNode: HTML Comment
     - DoctypeNode: `<!DOCTYPE html>` type of shit
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

#### Finding Link Nodes
 - Want to change this `Parse()` function so that as we're going through the DFS and we're looking ata ll the diff nodes, as soon as we have a link, we want to note that we have a link, and save it somewhere. Then our `Parse()` func can use that to build the `Link` types we created earlier.


#### Building Link Structs
 - create a new func `buildLink(n *htmlNode) Link {}`which takes in the pointer to a Node, build a Link struct and returns that
 - In `buildLink()`, We range over the attributes of each Node using `range n.Attr {}` because n.Attr is a slice, not a key/value map. (Why is n.Attr a slice? I only get one element anyway. ??)
 - We can get the Href property from the Node element through its attributes property. However we cannot get the text using `n.Data` . This is because this link element is an ElementNode that wraps a TextNode. To get the text inside the we need to access `n.FirstChild.Data`. (INACCURATE - CORRECTED IN NEXT SECTION)


`buildLink()` currently looks like this:
```go
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
```

#### Extracting Node Text
 - We actually can't access all the text of a link ElementNode with `n.FirstChild.Data` because that only returns a shallow copy of text exactly in the FirstChild. It will not return text in nested elements. 

For example:
```html
<a href="/other-page"> 
    A link to another page 
	<span> some span  </span>
</a>
```

Retreiving `n.FirstChild.Data` where `<a>` is `n.FirstChild` would only return "A link to another page".

We want exactly "A link to another page some span". 
 - We will have to perform a DFS to extract all the text inside a specific node.      
 - This impelmentation is v similar to the recursive DFS `linkNodes()` func, but this isn't the most optimal way to do this. You could look into a Byte Buffer tio build the string in a more optimized way.

- If text has a bunch of newlines we should probably trim it. The easiest way to do that is using `strings.Fields(str)`. This splits a string into an array using any amount of whitespace as the separator. This is very similar to Javascript `String.split()`  https://golang.org/pkg/strings/#Fields

- After using `strings.Fields(str)` we can join them together using `strings.Join(strSlice)`. This takes a slice of strings and joins them with a single separator, in our case a single space. This is very similar to Javascript `Array.join()` except explicitly for string arrays. https://golang.org/pkg/strings/#Join

- We can use `strings.Fields()` and `strings.Join()` to remove any whitespace and normalize the string to separate each word by exactly 1 space.

- If text has a bunch of newlines we should probably trim it. The easiest way to do that is using `strings.Fields()`:  https://golang.org/pkg/strings/#Fields


The code for the recursive `extractText()` function looks like:
```go
func extractText(n *html.Node) string {

	// Base case for recursion,
	// We have a TextNode, return the text data
	if n.Type == html.TextNode {
		return n.Data
	}
	// Node type is anything other than ElementNode	{
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	// DFS and append return value to string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += extractText(c) + ""
	}

	return strings.Join(strings.Fields(ret), " ")
}
```


#### Parsing the examples