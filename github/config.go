package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/logging"
	"golang.org/x/oauth2"
)

type Config struct {
	Token   string
	Owner   string
	BaseURL string
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
	tc := oauth2.NewClient(oauth2.NoContext, ts)

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
