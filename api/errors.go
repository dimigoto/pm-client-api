package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

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
