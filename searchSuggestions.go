package gogiphy

import (
	"fmt"
	"net/http"
)

func (search *Search) SearchSuggestions(term string) (*searchSuggestions, error) {
	url := buildBaseURL("tags/related", search.client.apiKey)
	if term == "" {
		return nil, fmt.Errorf("term may not be empty")
	}
	url = fmt.Sprintf("%s/%s", url, term)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := search.client.httpClient.Do(req)
	return convertResponseToSearchSuggestions(resp), err
}
