package gogiphy

import "net/http"

func (search *Search) TrendingSearch() (*trending, error) {
	url := buildBaseURL("trending/searches", search.client.apiKey)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := search.client.httpClient.Do(req)
	return convertResponseToTrending(resp), err
}
