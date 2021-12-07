package user_client

import (
	"net/url"
)

func PrepareURL(host, endpoint string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   endpoint,
	}
}
