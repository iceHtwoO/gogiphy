package main

import "github.com/ice-h2o/gogiphy"

/*
Providers users a list of valid terms that completes the given tag on the GIPHY network.
https://developers.giphy.com/docs/api/endpoint#autocomplete
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	autocomplete()
}

func autocomplete() {
	//Get recommended search terms
	data, err := gc.Search.Autocomplete("hel")
	if err != nil {
		println(err)
	}

	for _, term := range data.Data {
		println(term.Name)
	}
}
