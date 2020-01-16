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
	Token      string
	Owner      string
	BaseURL    string
	Insecure   bool
	Individual bool
	Anonymous  bool
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
	var ts oauth2.TokenSource
	var tc *http.Client

	ctx := context.Background()

	if c.Insecure {
		insecureClient := insecureHttpClient()
		ctx = context.WithValue(ctx, oauth2.HTTPClient, insecureClient)
	}

	if !c.Individual {
		owner.Organization = true
	}

	// Either Owner needs to be set, or Individual needs to be true
	if c.Owner != "" && c.Individual {
		return nil, fmt.Errorf("If `individual` is true, `owner` cannot be set.")
	}
	if c.Owner == "" && !c.Individual {
		return nil, fmt.Errorf("If `individual` is false, `owner` is required.")
	}

	// Either run as anonymous, or run with a Token
	if c.Token != "" && c.Anonymous {
		return nil, fmt.Errorf("If `anonymous` is true, `token` cannot be set.")
	}
	if c.Token == "" && !c.Anonymous {
		return nil, fmt.Errorf("If `anonymous` is false, `token` is required.")
	}

	if !c.Anonymous {
		ts = oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: c.Token},
		)
	}

	tc = oauth2.NewClient(ctx, ts)

	if c.Anonymous {
		tc.Transport = http.DefaultTransport
	} else {
		tc.Transport = NewEtagTransport(tc.Transport)
	}

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
