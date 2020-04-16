package github

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"golang.org/x/oauth2"
)

type Config struct {
	Token    string
	Owner    string
	BaseURL  string
	Insecure bool
}

type Owner struct {
	name           string
	id             int64
	client         *github.Client
	StopContext    context.Context
	IsOrganization bool
}

// Client configures and returns a fully initialized GithubClient
func (c *Config) Client() (interface{}, error) {
	var owner Owner
	owner.name = c.Owner
	owner.IsOrganization = false
	var tc *http.Client

	ctx := context.Background()

	if c.Insecure {
		insecureClient := insecureHttpClient()
		ctx = context.WithValue(ctx, oauth2.HTTPClient, insecureClient)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)
	tc = oauth2.NewClient(ctx, ts)
	tc.Transport = logging.NewTransport("Github", tc.Transport)

	owner.client = github.NewClient(tc)
	if c.BaseURL != "" {
		u, err := url.Parse(c.BaseURL)
		if err != nil {
			return nil, err
		}
		owner.client.BaseURL = u
	}

	remoteOrg, _, _ := (*owner.client).Organizations.Get(ctx, owner.name)
	if remoteOrg != nil {
		owner.IsOrganization = true
		owner.id = remoteOrg.GetID()
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
