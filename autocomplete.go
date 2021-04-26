package gogiphy

import (
	"fmt"
	"net/http"
	"strconv"
)

// Autocomplete will return 5 recommended strings for your searchTerm
// searchTerm : will be autocompleted
func (search *Search) Autocomplete(searchTerm string) (*searchSuggestions, error) {
	return search.SpecifiedAutocomplete(searchTerm, -1, -1)
}

// SpecifiedAutocomplete will return 5 recommended strings for your searchTerm
// searchTerm : will be autocompleted
// limit : will limit recommendation's Default: 5; Unset: -1
// offset : Specifies the starting point. Default: 0; Unset: -1
func (search *Search) SpecifiedAutocomplete(searchTerm string, limit int, offset int) (*searchSuggestions, error) {
	url := buildBaseURL("gifs/Search/tags", search.client.apiKey)
	if searchTerm == "" {
		return nil, fmt.Errorf("term may not be empty")
	}
	addParam(&url, "q", searchTerm)
	if limit > 0 {
		addParam(&url, "limit", strconv.Itoa(limit))
	}

	if offset > 0 {
		addParam(&url, "offset", strconv.Itoa(offset))
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, _ := search.client.httpClient.Do(req)
	return convertResponseToSearchSuggestions(resp), nil
}
