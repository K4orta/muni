package muni

import "net/http"

func transitRequest(reqURL string) (*http.Response, error) {
	c := http.Client{}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}
