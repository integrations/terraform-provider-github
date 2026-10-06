package ghclient

import (
	"fmt"
	"io"
	"net/http"
)

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
