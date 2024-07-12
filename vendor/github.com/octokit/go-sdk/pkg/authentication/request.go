package authentication

import (
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg/headers"
)

// Request provides a wrapper around Kiota's abs.RequestInformation type
type Request struct {
	*abs.RequestInformation
}

// WithTokenAuthentication sets the Authorization header to the given token,
// prepended by the AuthType
func (r *Request) WithTokenAuthentication(token string) {
	if r.Headers.ContainsKey(headers.AuthorizationKey) {
		r.Headers.Remove(headers.AuthorizationKey)
	}
	r.Headers.Add(headers.AuthorizationKey, fmt.Sprintf("%v %v", headers.AuthType, token))
}

// WithUserAgent allows the caller to set the User-Agent string for each request
func (r *Request) WithUserAgent(userAgent string) {
	if r.Headers.ContainsKey(headers.UserAgentKey) {
		r.Headers.Remove(headers.UserAgentKey)
	}
	r.Headers.Add(headers.UserAgentKey, userAgent)
}

// WithDefaultUserAgent sets the default User-Agent string for each request
func (r *Request) WithDefaultUserAgent() {
	r.WithUserAgent(headers.UserAgentValue)
}

// WithAPIVersion sets the API version header for each request
func (r *Request) WithAPIVersion(version string) {
	if r.Headers.ContainsKey(headers.APIVersionKey) {
		r.Headers.Remove(headers.APIVersionKey)
	}
	r.Headers.Add(headers.APIVersionKey, version)
}

// WithDefaultAPIVersion sets the API version header to the default (the version used
// to generate the code) for each request
func (r *Request) WithDefaultAPIVersion() {
	r.WithAPIVersion(headers.APIVersionValue)
}
