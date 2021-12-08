package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	endpointIsIn = "/channel/%s/is-in"
)

type ChannelService struct {
	client *http.Client
	host string
}

func NewChannelService(client *http.Client, host string) *ChannelService {
	return &ChannelService{
		client: client,
		host: host,
	}
}

func (s *ChannelService) IsUserInChannel(userID, channelID string) (bool, error) {
	type isUserInChannelRequest struct {
		UserID string `json:"user_id"`
	}

	request, err := createRequest(
		http.MethodPost,
		prepareURL("http", s.host, fmt.Sprintf(endpointIsIn, channelID)),
		isUserInChannelRequest{UserID: userID},
	)
	if err != nil {
		return false, err
	}

	response, err := s.client.Do(request)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		return boolFromResponse(response)
	}

	return false, parseError(response)
}

func boolFromResponse(r *http.Response) (bool, error) {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return false, err
	}

	var result bool

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}
