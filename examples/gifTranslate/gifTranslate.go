package main

import "github.com/ice-h2o/gogiphy"

/*
Translate and SpecifiedTranslate converts words and phrases to the perfect GIF using GIPHY's special sauce algorithm.
https://developers.giphy.com/docs/api/endpoint#translate
*/

var gc *gogiphy.Client

func main() {
	//Creating the gogiphy client
	gc = gogiphy.NewClient(nil, "API-KEY", gogiphy.English)
	translateGif()
	specifiedTranslate()
}

func translateGif() {
	//Using basic gif translate
	gifData, err := gc.Gif.Translate("minecraft")
	if err != nil {
		println(err)
	}

	println("----------------")
	println(gifData.Data.Title)                //getting the title of the gif
	println(gifData.Data.Images.Downsized.URL) //getting the gif url of the downsized version
}

func specifiedTranslate() {
	//Using extended gif translate
	gifData, err := gc.Gif.SpecifiedTranslate("golang", 5, "" /*Unset*/)
	if err != nil {
		println(err)
	}

	println("----------------")
	println(gifData.Data.Title)                //getting the title of the gif
	println(gifData.Data.Images.Downsized.URL) //getting the gif url of the downsized version
}
