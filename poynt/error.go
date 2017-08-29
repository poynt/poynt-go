package poynt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	Code             string
	Message          string
	DeveloperMessage string
	RequestId        string
}

type Error struct {
	StatusCode       int
	Status           string
	ErrorCode        string // poynt api error code
	Message          string
	DeveloperMessage string
	RequestId        string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Status: %s, Status Code: %d", e.Status, e.StatusCode)
}

func (e *Error) IsNotFound() bool {
	return e.StatusCode == 404
}

// Handles a http error if there is one by creating the Error instance.
// Returns nil if status code is 200
func ErrorHandler(resp *http.Response) *Error {
	if resp.StatusCode >= 400 {
		err := &Error{}
		err.StatusCode = resp.StatusCode
		err.Status = resp.Status

		fmt.Println("API error", err, "\npath", resp.Request.URL)

		apiError := new(ApiError)
		decodeError := json.NewDecoder(resp.Body).Decode(apiError)

		if decodeError == nil {
			err.ErrorCode = apiError.Code
			err.Message = apiError.Message
			err.DeveloperMessage = apiError.DeveloperMessage
			err.RequestId = apiError.RequestId
		}

		PrettyPrint(err)

		return err

	} else {
		return nil

	}
}
