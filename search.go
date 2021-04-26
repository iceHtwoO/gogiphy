package gogiphy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Search struct {
	client Client
}

type searchSuggestions struct {
	Data []term `json:"data"`
	Meta meta   `json:"meta"`
}

type term struct {
	Name string `json:"name"`
}

type trending struct {
	Data []string `json:"data"`
	Meta meta     `json:"meta"`
}

func convertResponseToTrending(response *http.Response) *trending {
	var data trending
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bodyBytes, &data)
	return &data
}

func convertResponseToSearchSuggestions(response *http.Response) *searchSuggestions {
	var data searchSuggestions
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bodyBytes, &data)
	return &data
}
