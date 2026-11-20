package ghclient

import (
	"context"
	"net/http"
	"time"

	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
)

const installationTokenTimeout = 30 * time.Second

type installationTokenSource struct {
	id     int64
	app    oauth2.TokenSource
	client *http.Client
	opts   []githubauth.InstallationTokenSourceOpt
	gate   chan struct{}
	token  *oauth2.Token
}

func newInstallationTokenSource(id int64, app oauth2.TokenSource, opts ClientOptions, authOpts ...githubauth.InstallationTokenSourceOpt) *installationTokenSource {
	return &installationTokenSource{
		id: id, app: app, opts: authOpts,
		client: &http.Client{Transport: getHTTPTransport(http.DefaultTransport, opts), Timeout: installationTokenTimeout},
		gate:   make(chan struct{}, 1),
	}
}

func (source *installationTokenSource) Token() (*oauth2.Token, error) {
	return source.TokenContext(context.Background())
}

func (source *installationTokenSource) TokenContext(ctx context.Context) (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, source.client.Timeout)
	defer cancel()

	select {
	case source.gate <- struct{}{}:
		defer func() { <-source.gate }()
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if source.token != nil && source.token.AccessToken != "" && (source.token.Expiry.IsZero() || time.Until(source.token.Expiry) > githubauth.DefaultExpirySkew) {
		return source.token, nil
	}

	authOpts := append([]githubauth.InstallationTokenSourceOpt{}, source.opts...)
	authOpts = append(authOpts, githubauth.WithHTTPClient(source.client), githubauth.WithContext(ctx))
	token, err := githubauth.NewInstallationTokenSource(source.id, source.app, authOpts...).Token()
	if err != nil {
		return nil, err
	}
	source.token = token
	return token, nil
}

type installationTransport struct {
	base   http.RoundTripper
	source *installationTokenSource
}

func (transport *installationTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	token, err := transport.source.TokenContext(request.Context())
	if err != nil {
		if request.Body != nil {
			request.Body.Close()
		}
		return nil, err
	}
	auth := &oauth2.Transport{Base: transport.base, Source: oauth2.StaticTokenSource(token)}
	return auth.RoundTrip(request)
}
