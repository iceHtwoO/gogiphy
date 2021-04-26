package main

import "github.com/ice-h2o/gogiphy"

/*
Random lets you add some weirdness to the conversation by returning a single random GIF related to the word or phrase entered.
https://developers.giphy.com/docs/api/endpoint#random
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	randomGifs()
	specifiedRandomGifs()
}

func randomGifs() {
	//Using basic gif random
	gif, err := gc.Gif.Random()
	if err != nil {
		println(err)
	}

	println("----------------")
	println(gif.Data.Title)                //getting the title of the gif
	println(gif.Data.Images.Downsized.URL) //getting the gif url of the downsized version
}

func specifiedRandomGifs() {
	//Using extended gif random
	gif, err := gc.Gif.SpecifiedRandom("memes", "" /*Unset*/, "1243543234")
	if err != nil {
		println(err)
	}

	println("----------------")
	println(gif.Data.Title)                //getting the title of the gif
	println(gif.Data.Images.Downsized.URL) //getting the gif url of the downsized version
}
