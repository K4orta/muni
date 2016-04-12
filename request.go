package muni

import "net/http"

// TransitConfig stores information used to make retquests to the NextBus API
type TransitConfig struct {
	DefaultURL string
}

var config = TransitConfig{
	DefaultURL: "http://webservices.nextbus.com/service/publicXMLFeed?a=sf-muni",
}

func transitRequest(query string) (*http.Response, error) {
	c := http.Client{}

	req, err := http.NewRequest("GET", config.DefaultURL+query, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// SetConfig allows adding a custom TransitConfig
func SetConfig(conf TransitConfig) {
	config = conf
}
