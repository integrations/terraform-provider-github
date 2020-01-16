package github

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/google/go-github/v25/github"
	"github.com/hashicorp/terraform/helper/logging"
	"golang.org/x/oauth2"
)

type Config struct {
	Token        string
	Owner string
	BaseURL      string
	Insecure     bool
}

type Owner struct {
	name         string
	client       *github.Client
	StopContext  context.Context
	Organization bool
}

func (o *Owner) IsOrganization() bool {
	return o.Organization
}

// Client configures and returns a fully initialized GithubClient
func (c *Config) Client() (interface{}, error) {
	var owner Owner
	owner.name = c.Owner
	owner.Organization = true
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)

	ctx := context.Background()

	if c.Insecure {
		insecureClient := insecureHttpClient()
		ctx = context.WithValue(ctx, oauth2.HTTPClient, insecureClient)
	}

	tc := oauth2.NewClient(ctx, ts)

	tc.Transport = NewEtagTransport(tc.Transport)

	tc.Transport = NewRateLimitTransport(tc.Transport)

	tc.Transport = logging.NewTransport("Github", tc.Transport)

	owner.client = github.NewClient(tc)
	if c.BaseURL != "" {
		u, err := url.Parse(c.BaseURL)
		if err != nil {
			return nil, err
		}
		owner.client.BaseURL = u
	}

	_, _, err := (*owner.client).Organizations.Get(context.TODO(), owner.name)
	if err != nil {
		owner.Organization = false
	}

	return &owner, nil
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
