Yahoo Finance Helper
==========

This was a very basic attempt at reading from the Yahoo Finance API.

Example Usage:

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

Note: There's more information available in the response, but not all have been implemented.