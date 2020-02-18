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
Startup
 - using html.Parse(r)
 - ignore nested links
 - make sure you just get something up and running first
 - using NodeType constants inssid the net/html package to help structure the data
  
