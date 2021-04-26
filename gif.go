package gogiphy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GifManager struct {
	client Client
}

type gifData struct {
	Data       []Gif      `json:"data"`
	Meta       meta       `json:"meta"`
	Pagination pagination `json:"pagination"`
}

type translateData struct {
	Data Gif  `json:"data"`
	Meta meta `json:"meta"`
}

type pagination struct {
	Offset     int32 `json:"offset"`
	TotalCount int32 `json:"total_count"`
	Count      int32 `json:"count"`
}

type Gif struct {
	Type             string `json:"type"`
	Id               string `json:"id"`
	Slug             string `json:"slug"`
	URL              string `json:"url"`
	BitlyURL         string `json:"bitly_url"`
	EmbedURL         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Rating           string `json:"rating"`
	ContentURL       string `json:"content_url"`
	SourceTLD        string `json:"source_tld"`
	SourcePostURL    string `json:"source_post_url"`
	UpdateDatetime   string `json:"update_datetime"`
	CreateDatetime   string `json:"create_datetime"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Title            string `json:"title"`
	User             user   `json:"user"`
	Images           Images `json:"images"`
}

func convertResponseToGifData(response *http.Response) *gifData {
	var data gifData
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bodyBytes, &data)
	return &data
}

func convertResponseToTranslateData(response *http.Response) *translateData {
	var data translateData
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bodyBytes, &data)
	return &data
}
