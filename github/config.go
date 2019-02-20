package github

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/logging"
	"golang.org/x/oauth2"
)

type Config struct {
	Token        string
	Organization string
	BaseURL      string
	Insecure     bool
}

type Organization struct {
	name        string
	client      *github.Client
	StopContext context.Context
}

// Client configures and returns a fully initialized GithubClient
func (c *Config) Client() (interface{}, error) {
	var org Organization
	org.name = c.Organization

	var ts oauth2.TokenSource
	if c.Token != "" {
		ts = oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: c.Token},
		)
	}

	ctx := context.Background()

	if c.Insecure {
		insecureClient := insecureHttpClient()
		ctx = context.WithValue(ctx, oauth2.HTTPClient, insecureClient)
	}

	tc := oauth2.NewClient(ctx, ts)

	if c.Token != "" {
		tc.Transport = NewEtagTransport(tc.Transport)
	} else {
		tc.Transport = http.DefaultTransport
	}

	tc.Transport = NewRateLimitTransport(tc.Transport)

	tc.Transport = logging.NewTransport("Github", tc.Transport)

	org.client = github.NewClient(tc)
	if c.BaseURL != "" {
		u, err := url.Parse(c.BaseURL)
		if err != nil {
			return nil, err
		}
		org.client.BaseURL = u
	}

	return &org, nil
}

func insecureHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
