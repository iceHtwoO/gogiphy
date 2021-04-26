package gogiphy

import (
	"fmt"
	"net/http"
	"strconv"
)

// Search returns an specified amount of gifs that fit best to the search term
// search : search for specific gifs
// limit : limit the amount of gifs Default: 5; Unset: -1
func (gif *GifManager) Search(search string, limit int) (*gifData, error) {
	return gif.SpecifiedSearch(search, limit, -1, "", gif.client.lang, "")
}

// SpecifiedSearch returns an specified amount of gifs that fit best to the search term
// search : search for specific gifs
// limit : limit the amount of gifs Default: 5; Unset: -1
// offset : Specifies the starting point. Default: 0; Unset: -1
// rating : Age rating
// lang : specify the language Default: "en"; Unset: ""
// randomId : An ID/proxy for a specific user. Default: -1
func (gif *GifManager) SpecifiedSearch(
	search string, limit int, offset int, rating string, lang string, randomId string,
) (*gifData, error) {
	url := buildBaseURL("gifs/Search", gif.client.apiKey)

	if search == "" {
		return &gifData{}, fmt.Errorf("Search may not be empty")
	}
	addParam(&url, "q", search)
	if rating != "" {
		addParam(&url, "rating", rating)
	}
	if lang != "" {
		addParam(&url, "lang", lang)
	}
	if randomId != "" {
		addParam(&url, "random_id", randomId)
	}
	if limit > 0 {
		addParam(&url, "limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		addParam(&url, "offset", strconv.Itoa(offset))
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := gif.client.httpClient.Do(req)
	return convertResponseToGifData(resp), err
}
