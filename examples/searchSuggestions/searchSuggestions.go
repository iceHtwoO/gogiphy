package main

import "github.com/ice-h2o/gogiphy"

/*
Providers users a list of tag terms related to the given tag on the GIPHY network.
https://developers.giphy.com/docs/api/endpoint#search-suggestions
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	searchSuggestions()
}

func searchSuggestions() {
	//Get search suggestions
	data, err := gc.Search.SearchSuggestions("gaming")
	if err != nil {
		println(err)
	}

	for _, term := range data.Data {
		println(term.Name)
	}
}
