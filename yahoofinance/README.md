Yahoo Finance Helper
==========

This is now DEPRECATED and no longer functional. Still remains as a learning example eh?

This was a very basic attempt at reading from the Yahoo Finance API.

Example Usage:
```go
	package main

	import (
		"fmt"
		"github.com/mickelsonm/go-helpers/yahoofinance"
	)

	func main() {
		//Read quotes for Google, IBM, Apple, Microsoft, Amazon
		quotes, err := yahoofinance.GetStockQuotes([]string{"GOOG", "IBM", "AAPL", "MSFT", "AMZN"})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", quotes)
	}
```
Note: There's more information available in the response, but not all have been implemented.
