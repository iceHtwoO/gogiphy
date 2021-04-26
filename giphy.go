package gogiphy

import (
	"fmt"
	"net/http"
)

const (
	baseURL = "https://api.giphy.com/v1"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
	lang       string

	Gif    *GifManager
	Search *Search
}

type user struct {
	AvatarURL   string `json:"avatar_url"`
	BannerURL   string `json:"banner_url"`
	ProfileURL  string `json:"profile_url"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type meta struct {
	Msg        string `json:"msg"`
	Status     int32  `json:"status"`
	ResponseId string `json:"response_id"`
}

func NewClient(client *http.Client, apiKey string, lang string) *Client {
	if client == nil {
		client = &http.Client{}
	}

	if lang == "" {
		lang = English
	}

	giphyClient := Client{
		httpClient: client,
		apiKey:     apiKey,
		baseURL:    baseURL,
		lang:       lang,
	}

	giphyClient.Gif = &GifManager{client: giphyClient}
	giphyClient.Search = &Search{client: giphyClient}

	return &giphyClient
}

func addParam(url *string, param string, value string) {
	*url = fmt.Sprintf("%s&%s=%s", *url, param, value)
}

func buildBaseURL(path string, key string) string {
	return fmt.Sprintf("%s/%s?api_key=%s", baseURL, path, key)
}
