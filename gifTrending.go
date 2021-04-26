package gogiphy

import (
	"net/http"
	"strconv"
	"strings"
)

// Trending will return gifData that contains an specified amount of trending gifs and its metadata.
// limit : the amount of gifs
func (gif *GifManager) Trending(limit int) (*gifData, error) {
	return gif.SpecifiedTrending(limit, -1, "", "")
}

// SpecifiedTrending will return gifData that contains an specified amount of trending gifs and its metadata.
// limit : the amount of gifs
// offset : Specifies the starting point. Default: 0; Unset: -1
// rating : Age rating
// randomId : An ID/proxy for a specific user. Default: -1
func (gif *GifManager) SpecifiedTrending(limit int, offset int, rating string, randomId string) (*gifData, error) {
	rating = strings.ToLower(rating)
	url := buildBaseURL("gifs/trending", gif.client.apiKey)
	if limit >= 0 {
		addParam(&url, "limit", strconv.Itoa(limit))
	}
	if offset >= 0 {
		addParam(&url, "offset", strconv.Itoa(offset))
	}
	if rating != "" {
		addParam(&url, "rating", rating)
	}
	if randomId != "" {
		addParam(&url, "random_id", randomId)
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := gif.client.httpClient.Do(req)
	return convertResponseToGifData(resp), err
}
