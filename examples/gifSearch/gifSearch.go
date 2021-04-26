package main

import "github.com/ice-h2o/gogiphy"

/*
Search and SpecifiedSearch will return the top posts that fit best to the search term.
https://developers.giphy.com/docs/api/endpoint#search
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	searchForGifs()
	specifiedSearchForGifs()
}

func searchForGifs() {
	//Using basic gif search
	gifData, err := gc.Gif.Search("Memes", 2)
	if err != nil {
		println(err)
	}

	for _, gif := range gifData.Data {
		println("----------------")
		println(gif.Title)                //getting the title of the gif
		println(gif.Images.Downsized.URL) //getting the gif url of the downsized version
	}
}

func specifiedSearchForGifs() {
	//Using extended gif search
	gifData, err := gc.Gif.SpecifiedSearch("Memes", 2, 2, gogiphy.PG, gogiphy.German, "" /*Unset*/)
	if err != nil {
		println(err)
	}

	for _, gif := range gifData.Data {
		println("----------------")
		println(gif.Title)                //getting the title of the gif
		println(gif.Images.Downsized.URL) //getting the gif url of the downsized version
	}
}
