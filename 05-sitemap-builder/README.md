# Sitemap Builder

A sitemap is basically a map of all of the pages within a specific domain. They are used by search engines and other tools to inform them of all of the pages on your domain.

One way these can be built is by first visiting the root page of the website and making a list of every link on that page that goes to a page on the same domain. For instance, on calhoun.io you might find a link to calhoun.io/hire-me/ along with several other links.

Once you have created the list of links, you could then visit each and add any new links to your list. By repeating this step over and over you would eventually visit every page that on the domain that can be reached by following links from the root page.

This will use the HTML Link parser built in example 04 to index every link on a page, then visit those links and index those links recursively.

Once you have determined all of the pages of a site, your sitemap builder should then output the data in the following XML format:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```

Where each page is listed in its own <url> tag and includes the <loc> tag inside of it.

From there you will likely need to figure out a way to determine if a link goes to the same domain or a different one. If it goes to a different domain we shouldn't include it in our sitemap builder, but if it is the same domain we should. Remember that links to the same domain can be in the format of `/just-the-path` or https://domain.com/with-domain, but both go to the same domain.

### Random
 - Cobra is a good tool for building CLI applications

### Overview
  - Useful - whenever you provide your site to a search engine, you usually provide a sitemap. Sitemap builders exist but this can help teach you something about how they work.
  - Use the link parser package we created before, use that to take the response body returned from querying a page, and parse the links out. 
  - Once you get the links for a page, you're going to figure out whic hlinks go to a page on the same domain. (no external links). These types of links are in the format `/just-the-path` or `https://domain.com//with-domain`
  - If you get just the path, you might have to do some work to get the domain included
  - Forget mailto links, fragment links, etc.
  - Be aware that links can be cyclical ( only sitemap the pages )
  ```  
   /about -> /contact
   /contact -> /pricing
   /pricing -> /about
  ```

  
#### Useful packages
  - net/http
  - linkMap package
  - encoding/xml
  - flag 
 

### Strategy
- You could draw a tree that represents your website
- RootPage '/'
   Children: /page/2, /poop, /pee, https://github.com/something 
   Su

We need to know two things: 1) What we do when we get to an individual page and 2) What to do at a higher level to know what the pages we need are

Every time we go to a page:
   - Get the HTML for the page
   - Use the link-parser package to parse out all the links for that page
   - Once we get the links, we will try to clean them up (add a domain to links that don't have a domain)
   - Go through all those links again and filter out irrelevant links (mailto links, external links.)

At this point we will have a list of links that go to other pages on the same domain.
Once we have that list, we will store it somewhere to associate the current page with the links it has. Next:

 - Traverse the newly created list of links and index the links on each of those pages. Create a new list associating these 2nd layer links with their children. Do this recursively 
 
 This is basically a Breadth-First Search (BFS)
 

 ### Main go

 ### GET-ting a webpage
  - https://golang.org/pkg/net/http/
  - use net/http to send outgoing GET/POST/PUT/etc requests 
  - the client MUST close the request once its done with it - to prevent memory leaks
  - use `defer` if you're fully handling your GET request in one function. May not be able to use defer if you have to pass the request forward (you would still have to close the req).
  ```go
  defer resp.Body.Close()
  ```

  `defer` tells go to run the following statement right when the function exits.
 Q:  Why not just write the code to clean up at the very end of your function? 
 A: `defer` helps readability and understandability of code - the cleanup is located right next to its associated code. Also if you have some case where you return before the end of the function, any cleanup at the end of the function would not be run. With `defer` it would be run.


  - use `io.Copy()` which allows you to copy from a io.Reader to an io.Writer (such as STDOUT). This works as compared to `fmt.Println()` because with fmt.Println, you are printing out the memory address and Go data structure for the response body. With STDOUT you are directly printing out the raw Resp.Body since you STD is a writer interface.  

  