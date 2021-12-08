package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"gopkg.in/guregu/null.v4"
)

const (
	endpointGetAllUsersByIds = "/users"
)

// User Структура пользователя
type User struct {
	ID           string      `json:"id"`
	Username     string      `json:"username"`
	LastName     string      `json:"last_name"`
	FirstName    null.String `json:"first_name"`
	Mobile       string      `json:"mobile"`
	Email        null.String `json:"email"`
	Bio          null.String `json:"bio"`
	Birthday     null.Time   `json:"birthday"`
	Status       int         `json:"status"`
	LastVisitAt  time.Time   `json:"last_visit_at"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    null.Time   `json:"updated_at"`
	ArchivedAt   null.Time   `json:"archived_at"`
}

func usersFromResponse(r *http.Response) ([]User, error) {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var users []User

	err = json.Unmarshal(responseBody, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

type UserService struct {
	client *http.Client
	host string
}

func NewUserService(client *http.Client, host string) *UserService {
	return &UserService{
		client: client,
		host: host,
	}
}

func (s *UserService) GetAllByIDs(ids []string) ([]User, error) {
	type getUsersByIDsRequest struct {
		ID []string `json:"id"`
	}

	request, err := createRequest(
		http.MethodPost,
		prepareURL("http", s.host, endpointGetAllUsersByIds),
		getUsersByIDsRequest{ID: ids},
	)
	if err != nil {
		return nil, err
	}

	response, err := s.client.Do(request)
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
		return usersFromResponse(response)
	}

	return nil, parseError(response)
}
