package gogiphy

import (
	"net/http"
)

func (gif *GifManager) Random() (*translateData, error) {
	return gif.SpecifiedRandom("", "", "")
}

func (gif *GifManager) SpecifiedRandom(tag string, rating string, randomId string) (*translateData, error) {
	url := buildBaseURL("gifs/random", gif.client.apiKey)
	if tag != "" {
		addParam(&url, "tag", tag)
	}
	if rating != "" {
		addParam(&url, "rating", rating)
	}
	if randomId != "" {
		addParam(&url, "random_id", randomId)
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := gif.client.httpClient.Do(req)
	return convertResponseToTranslateData(resp), err
}
