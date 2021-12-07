package user_client

import (
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const (
	endpointGetAllUsersByIds = "/users"
)

type UserClient struct {
	serviceHost string
}

func NewClient(serviceHost string) *UserClient {
	return &UserClient{
		serviceHost: serviceHost,
	}
}

func (c *UserClient) GetAllByIDs(ids []string) ([]User, error) {
	request, err := CreateGetAllByIDsRequest(
		PrepareURL(c.serviceHost, endpointGetAllUsersByIds),
		ids,
	)
	if err != nil {
		return nil, err
	}

	response, err := new(http.Client).Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		return UsersFromResponse(response)
	}

	return nil, parseError(response)
}

func parseError(r *http.Response) error {
	if r.StatusCode == http.StatusUnprocessableEntity {
		validationErrors, err := ValidationErrorsFromResponse(r)
		if err != nil {
			return errors.Wrap(err, "User service error")
		}

		return errors.Wrap(validationErrors, "User service validation error")
	}

	responseError, err := ResponseErrorFromResponse(r)
	if err != nil {
		return errors.Wrap(err, "User service error")
	}

	return errors.Wrap(responseError, "User service response error")
}
