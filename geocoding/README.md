Geocoding Helper
==========

Lookups are done using an address or latitude/longitude pair.

Address: Just a string of Address, City, State/Province Postal Code

Location: A coordinate struct, which you set Latitude and Longitude.

Example:
```go
	package main

	import(
		"log"

		"github.com/mickelsonm/go-helpers/geocoding"
	)

	func main(){
		lookup := geocoding.Lookup{
			//address based lookup
			//Address: "309 South Barstow Street, Eau Claire, WI 54701",
			//latitude/longitude-based lookup
			Location: &geocoding.Point{
				Latitude: 123.45,
				Longitude: 123.45,
			}
		}
		results, err := lookup.Search()
		if err != nil{
			log.Println(err)
			return
		}
		log.Println(results)
	}
```
