package httpx

import "errors"

// errRetryServerError represents a server error that can be retried.
var errRetryServerError = errors.New("server error")
