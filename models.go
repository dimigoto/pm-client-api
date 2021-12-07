package user_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gopkg.in/guregu/null.v4"
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

func UsersFromResponse(r *http.Response) ([]User, error) {
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

type ResponseError struct {
	ErrorMessage string `json:"error"`
}

func (e *ResponseError) Error() string {
	return e.ErrorMessage
}

func ResponseErrorFromResponse(r *http.Response) (*ResponseError, error) {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var responseError = ResponseError{}

	err = json.Unmarshal(responseBody, &responseError)
	if err != nil {
		return nil, err
	}

	return &responseError, nil
}

type ValidationErrors struct {
	Errors map[string] interface{} `json:"errors"`
}

func (e *ValidationErrors) Error() string {
	return fmt.Sprint(e.Errors)
}

func ValidationErrorsFromResponse(r *http.Response) (*ValidationErrors, error) {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var validationErrors = ValidationErrors{}

	err = json.Unmarshal(responseBody, &validationErrors)
	if err != nil {
		return nil, err
	}

	return &validationErrors, nil
}
