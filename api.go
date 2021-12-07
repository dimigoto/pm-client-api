package user_client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func CreateGetAllByIDsRequest(url *url.URL, ids []string) (*http.Request, error) {
	type getUsersByIDsRequest struct {
		ID []string `json:"id"`
	}

	var requestBody bytes.Buffer

	requestGetUsersByIDs := &getUsersByIDsRequest{
		ID: ids,
	}

	err := json.NewEncoder(&requestBody).Encode(requestGetUsersByIDs)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, url.String(), &requestBody)
}
