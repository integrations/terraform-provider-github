package github

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-github/v25/github"
	"github.com/hashicorp/terraform/helper/logging"
	"golang.org/x/oauth2"
)

type Config struct {
	Token        string
	Organization string
	BaseURL      string
	Insecure     bool
	Individual   bool
}

type Organization struct {
	name        string
	client      *github.Client
	StopContext context.Context
}

// Client configures and returns a fully initialized GithubClient
func (c *Config) Client() (interface{}, error) {
	var org Organization

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)

	ctx := context.Background()

	if c.Insecure {
		insecureClient := insecureHttpClient()
		ctx = context.WithValue(ctx, oauth2.HTTPClient, insecureClient)
	}

	if c.Individual {
		org.name = ""
	} else if c.Organization != "" {
		org.name = c.Organization
	} else {
		return nil, fmt.Errorf("If `individual` is false, `organization` is required.")
	}

	tc := oauth2.NewClient(ctx, ts)

	tc.Transport = NewEtagTransport(tc.Transport)

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
