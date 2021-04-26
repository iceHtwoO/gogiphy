package main

import "github.com/ice-h2o/gogiphy"

/*
Provides users a list of the most popular trending search terms on the GIPHY network.
https://developers.giphy.com/docs/api/endpoint#trending-search-terms
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	trendingGifs()
}

func trendingGifs() {
	//Get trending search terms
	data, err := gc.Search.TrendingSearch()
	if err != nil {
		println(err)
	}

	for _, search := range data.Data {
		println(search)
	}
}
