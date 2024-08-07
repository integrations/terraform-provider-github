package authentication

import (
	"context"

	abs "github.com/microsoft/kiota-abstractions-go"
)

// TokenProvider may use a token to authenticate each request. It also can be
// used to configure UserAgent strings, API Versions, and other request configuration.
// Note that GitHub App authentication is set at the client transport level. See the
// docs for pkg.NewApiClient for more.
type TokenProvider struct {
	options []TokenProviderOption
}

// TokenProviderOption provides a functional option
// for configuring a TokenProvider.
type TokenProviderOption func(*TokenProvider, *Request)

// WithTokenAuthentication sets the AuthorizationToken for each request to the given token.
func WithTokenAuthentication(token string) TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithTokenAuthentication(token)
	}
}

// WithDefaultUserAgent sets the User-Agent string sent for requests to the default
// for this SDK.
func WithDefaultUserAgent() TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithDefaultUserAgent()
	}
}

// WithUserAgent sets the User-Agent string sent with each request.
func WithUserAgent(userAgent string) TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithUserAgent(userAgent)
	}
}

// WithDefaultAPIVersion sets the API version header sent with each request.
func WithDefaultAPIVersion() TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithDefaultAPIVersion()
	}
}

// WithAPIVersion sets the API version header sent with each request.
func WithAPIVersion(version string) TokenProviderOption {
	return func(t *TokenProvider, r *Request) {
		r.WithAPIVersion(version)
	}
}

// TODO(kfcampbell): implement new constructor with allowedHosts

// NewTokenProvider creates an instance of TokenProvider with the specified token and options.
func NewTokenProvider(options ...TokenProviderOption) *TokenProvider {
	provider := &TokenProvider{
		options: options,
	}
	return provider
}

// defaultHandlers contains our "sensible defaults" for TokenProvider initialization
var defaultHandlers = []TokenProviderOption{WithDefaultUserAgent(), WithDefaultAPIVersion()}

// AuthenticateRequest applies the default options for each request, then the user's options
// (if present in the TokenProvider). User options are guaranteed to be run in the order they
// were input.
func (t *TokenProvider) AuthenticateRequest(context context.Context, request *abs.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	reqWrapper := &Request{RequestInformation: request}

	if reqWrapper.Headers == nil {
		reqWrapper.Headers = abs.NewRequestHeaders()
	}

	for _, option := range defaultHandlers {
		option(t, reqWrapper)
	}

	// apply user options after defaults
	for _, option := range t.options {
		option(t, reqWrapper)
	}
	return nil
}
