package api_client

import (
	"net/http"

	"github.com/dimigoto/pm-client-api/api"
)

type ApiClient struct {
	client *http.Client
	config *ClientConfig
	userService *api.UserService
}

func NewApiClient(config *ClientConfig) *ApiClient {
	return &ApiClient{
		client: new(http.Client),
		config: config,
	}
}

func (c* ApiClient) UserService() *api.UserService {
	if c.userService == nil {
		c.userService = api.NewUserService(c.client, c.config.UserServiceHost)
	}

	return c.userService
}
