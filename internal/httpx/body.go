package httpx

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// errRequestNil is returned when the request is nil.
	errRequestNil = errors.New("request is nil")

	// errRequestGetBodyNil is returned when the request's GetBody function is nil, which is required for rewinding the request body for retries.
	errRequestGetBodyNil = errors.New("get body is nil")
)

// rewindRequestBody rewinds the request body to the beginning, if possible. This is necessary for retrying requests with a non-nil body.
func rewindRequestBody(req *http.Request) error {
	if req == nil {
		return errRequestNil
	}

	if req.Body == nil || req.Body == http.NoBody {
		return nil
	}

	if req.GetBody == nil {
		return errRequestGetBodyNil
	}

	body, err := req.GetBody()
	if err != nil {
		return fmt.Errorf("failed to get request body: %w", err)
	}

	req.Body = body
	return nil
}
