Hacker News Helper
==========

This helper was a fun way to see how Go plays with RSS.

I also saw a code challenge recently, which I felt like trying for the fun of it. The challenge read as follows:
>> On Hacker News (https://news.ycombinator.com/) there is an RSS feed for the current articles. Using any language and framework, take this RSS feed and convert it to a RESTful service that returns the contents of this feed in a JSON format for a GET request. 


Example usage:
```go
	package main

	import (
		"log"
		"net/http"

		"github.com/gin-gonic/gin"
		"github.com/mickelsonm/go-helpers/hackernews"
	)

	func main() {
		//we're going to use a Go web framework called Gin
		g := gin.Default()

		//this is our restful api endpoint/service link, which will give the user a json format
		//for the Hacker News RSS feed
		g.GET("/hackernews", func(c *gin.Context) {
			//TODO: we could optimize this by implementing some sort of caching mechanism

			//lets get the Hacker News RSS feed
			feed, err := hackernews.GetRssFeed()
			if err != nil {
				c.Fail(http.StatusInternalServerError, err)
				return
			}

			//success! let's serve up our json!
			c.JSON(http.StatusOK, feed)
		})

		//lets be nice and handle our root path
		g.GET("/", func(c *gin.Context) {
			//TODO: this is really hacky, in a real app we would just create an html file/template
			//then just serve up the html file
			//like this: c.HTML(http.StatusOK, "index", nil)
			c.Writer.Header().Set("Content-Type", "text/html")
			c.Data(http.StatusOK, []byte("Hey, checkout: <a href=\"/hackernews\">Hacker News JSON</a>"))
		})

		log.Print("Server is listening on port 3000...")
		g.Run(":3000")
	}
```
