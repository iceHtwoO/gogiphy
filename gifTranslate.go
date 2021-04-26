package gogiphy

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (gif *GifManager) Translate(search string) (*translateData, error) {
	return gif.SpecifiedTranslate(search, -1, "")
}

func (gif *GifManager) SpecifiedTranslate(search string, weirdness int, randomId string) (*translateData, error) {
	url := buildBaseURL("gifs/translate", gif.client.apiKey)
	search = strings.TrimSpace(search)

	if search == "" {
		return &translateData{}, fmt.Errorf("Search may not be empty")
	}
	addParam(&url, "s", search)
	if weirdness >= 0 && weirdness <= 10 {
		addParam(&url, "weirdness", strconv.Itoa(weirdness))
	}
	if randomId != "" {
		addParam(&url, "random_id", randomId)
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := gif.client.httpClient.Do(req)
	return convertResponseToTranslateData(resp), err
}
