package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func createRequest(method string, url *url.URL, payload interface{}) (*http.Request, error) {
	var requestBody bytes.Buffer

	err := json.NewEncoder(&requestBody).Encode(payload)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(method, url.String(), &requestBody)
}

func prepareURL(scheme, host, endpoint string) *url.URL {
	return &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   endpoint,
	}
}
