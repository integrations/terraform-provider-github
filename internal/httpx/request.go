package httpx

import "net/http"

// isIdempotent returns true if the given request is idempotent. A request is considered idempotent if it is a GET, HEAD, OPTIONS, PUT, DELETE, or TRACE request, or if it has an "Idempotency-Key" header.
func isIdempotent(req *http.Request) bool {
	switch req.Method {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodDelete, http.MethodTrace:
		return true
	default:
		return req.Header.Get("Idempotency-Key") != ""
	}
}

// canRewind returns true if the request body can be sent again.
func canRewind(req *http.Request) bool {
	return req.Body == nil || req.Body == http.NoBody || req.GetBody != nil
}
