package main

import "github.com/ice-h2o/gogiphy"

/*
Trending and SpecifiedTrending returns a list of the most relevant and engaging content each and every day.
https://developers.giphy.com/docs/api/endpoint#trending
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	trendingGifs()
	specifiedTrendingGifs()
}

func trendingGifs() {
	//Using basic gif trending search
	gifData, err := gc.Gif.Trending(2)
	if err != nil {
		println(err)
	}

	for _, gif := range gifData.Data {
		println("----------------")
		println(gif.Title)                //getting the title of the gif
		println(gif.Images.Downsized.URL) //getting the gif url of the downsized version
	}
}

func specifiedTrendingGifs() {
	//Using extended gif trending search
	gifData, err := gc.Gif.SpecifiedTrending(2, 2, gogiphy.R, "" /*Unset*/)
	if err != nil {
		println(err)
	}

	for _, gif := range gifData.Data {
		println("----------------")
		println(gif.Title)                //getting the title of the gif
		println(gif.Images.Downsized.URL) //getting the gif url of the downsized version
	}
}
