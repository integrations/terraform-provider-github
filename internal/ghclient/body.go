package ghclient

import (
	"fmt"
	"io"
	"net/http"
)

// rewindRequestBody rewinds the request body to the beginning, if possible. This is necessary for retrying requests with a non-nil body.
func rewindRequestBody(req *http.Request) error {
	if req == nil {
		return fmt.Errorf("request is nil")
	}

	if req.Body == nil || req.Body == http.NoBody {
		return nil
	}

	if req.GetBody == nil {
		return fmt.Errorf("cannot rewind request body: get body is nil")
	}

	body, err := req.GetBody()
	if err != nil {
		return fmt.Errorf("failed to get request body: %w", err)
	}

	req.Body = body
	return nil
}

// drainResponseBody reads and closes a [http.Response] body, up to 4KB, to avoid memory leaks.
func drainResponseBody(resp *http.Response) error {
	if resp == nil {
		return fmt.Errorf("response is nil")
	}

	if resp.Body == nil {
		return nil
	}

	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
	_ = resp.Body.Close()
	return nil
}
